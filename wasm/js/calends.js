// Stolen from Emscripten boilerplate
const ENVIRONMENT_IS_WEB = typeof window === 'object';
const ENVIRONMENT_IS_WORKER = typeof importScripts === 'function';
const ENVIRONMENT_HAS_NODE = typeof process === 'object' && typeof process.versions === 'object' && typeof process.versions.node === 'string';
const ENVIRONMENT_IS_NODE = ENVIRONMENT_HAS_NODE && !ENVIRONMENT_IS_WEB && !ENVIRONMENT_IS_WORKER;
const ENVIRONMENT_IS_SHELL = !ENVIRONMENT_IS_WEB && !ENVIRONMENT_IS_NODE && !ENVIRONMENT_IS_WORKER;

var scriptDirectory = '';
function locateFile(path) {
    return scriptDirectory + path;
}

var read_, readAsync, readBinary, onDestroy;

if (ENVIRONMENT_IS_NODE) {
    const { addDestructor } = require('object-destructor');
    require('./wasm_loader');

    scriptDirectory = __dirname + '/';

    // Expose functionality in the same simple way that the shells work
    // Note that we pollute the global namespace here, otherwise we break in node
    var nodeFS;
    var nodePath;

    read_ = function shell_read(filename, binary) {
        var ret;
        if (!nodeFS) nodeFS = require('fs');
        if (!nodePath) nodePath = require('path');
        filename = nodePath['normalize'](filename);
        ret = nodeFS['readFileSync'](filename);
        return binary ? ret : ret.toString();
    };

    readBinary = function readBinary(filename) {
        var ret = read_(filename, true);
        if (!ret.buffer) {
            ret = new Uint8Array(ret);
        }
        return ret;
    };

    onDestroy = addDestructor;
} else if (ENVIRONMENT_IS_SHELL) {
    const { addDestructor } = require('object-destructor');
    require('./wasm_loader');

    if (typeof read != 'undefined') {
        read_ = function shell_read(f) {
            return read(f);
        };
    }

    readBinary = function readBinary(f) {
        var data;
        if (typeof readbuffer === 'function') {
            return new Uint8Array(readbuffer(f));
        }
        data = read(f, 'binary');
        return data;
    };

    if (typeof print !== 'undefined') {
        // Prefer to use print/printErr where they exist, as they usually work better.
        if (typeof console === 'undefined') console = {};
        console.log = print;
        console.warn = console.error = typeof printErr !== 'undefined' ? printErr : print;
    }

    onDestroy = addDestructor;
} else if (ENVIRONMENT_IS_WEB || ENVIRONMENT_IS_WORKER) {
    if (ENVIRONMENT_IS_WORKER) { // Check worker, not web, since window could be polyfilled
        scriptDirectory = self.location.href;
    } else if (document.currentScript) { // web
        scriptDirectory = document.currentScript.src;
    }
    // blob urls look like blob:http://site.com/etc/etc and we cannot infer anything from them.
    // otherwise, slice off the final part of the url to find the script directory.
    // if scriptDirectory does not contain a slash, lastIndexOf will return -1,
    // and scriptDirectory will correctly be replaced with an empty string.
    if (scriptDirectory.indexOf('blob:') !== 0) {
        scriptDirectory = scriptDirectory.substr(0, scriptDirectory.lastIndexOf('/') + 1);
    } else {
        scriptDirectory = '';
    }

    read_ = function shell_read(url) {
        var xhr = new XMLHttpRequest();
        xhr.open('GET', url, false);
        xhr.send(null);
        return xhr.responseText;
    };

    if (ENVIRONMENT_IS_WORKER) {
        readBinary = function readBinary(url) {
            var xhr = new XMLHttpRequest();
            xhr.open('GET', url, false);
            xhr.responseType = 'arraybuffer';
            xhr.send(null);
            return new Uint8Array(xhr.response);
        };
    }

    readAsync = function readAsync(url, onload, onerror) {
        var xhr = new XMLHttpRequest();
        xhr.open('GET', url, true);
        xhr.responseType = 'arraybuffer';
        xhr.onload = function xhr_onload() {
            if (xhr.status == 200 || (xhr.status == 0 && xhr.response)) { // file URLs can return 0
                onload(xhr.response);
                return;
            }
            onerror();
        };
        xhr.onerror = onerror;
        xhr.send(null);
    };

    onDestroy = addDestructor;
} else {
    throw new Error('environment detection error');
}

// No longer stolen code
const go = new globalThis.Go();
let mod, inst;
let halt = true;

function initCalends(callback) {
    if (!halt) return;
    halt = false;

    if (!(mod instanceof WebAssembly.Module)) {
        (
            typeof readAsync === 'function' ?
                WebAssembly.instantiateStreaming(fetch(locateFile('calends.wasm')), go.importObject) :
                WebAssembly.instantiate(readBinary(locateFile('calends.wasm')), go.importObject)
        )
        .then((result) => {
            mod = result.module;
            inst = result.instance;
            run(); // Don't await because we need to run concurrent with the Go runtime
            waitForCalendsFuncs().then(() => {
                globalThis.CalendsFuncs.registerPanicHandler((message) => {
                    throw new CalendsError(message);
                });
                if (typeof callback == 'function') callback();
            });
        }).catch((err) => {
            throw new CalendsError(err);
        });
    } else {
        run();
        waitForCalendsFuncs().then(() => {
            globalThis.CalendsFuncs.registerPanicHandler((message) => {
                throw new CalendsError(message);
            });
            if (typeof callback == 'function') callback();
        });
    }

    async function idleFor(msec) {
        return new Promise(resolve => setTimeout(resolve, msec));
    }

    function waitForCalendsFuncs() {
        return new Promise(async function (resolve) {
            while (typeof globalThis.CalendsFuncs === 'undefined') {
                await idleFor(500);
            }

            resolve();
        });
    }

    async function run() {
        if (!halt) {
            try {
                await go.run(inst);
                inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
                if (!halt) {
                    initCalends();
                }
            } catch(err) {
                throw CalendsError(err);
            }
        }
    }
}

function stopCalends(callback) {
    if (halt) return;
    halt = true;
    globalThis.CalendsFuncs.stop();
    waitForCalendsFuncsToVanish().then(() => {
        if (typeof callback == 'function') callback();
    });

    async function idleFor(msec) {
        return new Promise(resolve => setTimeout(resolve, msec));
    }

    function waitForCalendsFuncsToVanish() {
        return new Promise(async function (resolve) {
            while (typeof globalThis.CalendsFuncs !== 'undefined') {
                await idleFor(500);
            }

            resolve();
        });
    }
}

class CalendsError extends Error {}

class Calends {
    #instance;

    // CREATE

    constructor(stamp, calendar, format) {
        let ref = 0;

        if (typeof stamp == 'bigint' && !calendar && !format) {
            ref = Number(stamp);
        } else {
            switch (typeof stamp) {
                case 'string':
                    ref = globalThis.CalendsFuncs.createString(stamp, calendar, format);
                    break;

                case 'number':
                    if (Number.isInteger(stamp))
                        ref = globalThis.CalendsFuncs.createInt64(stamp, calendar, format);
                    else
                        ref = globalThis.CalendsFuncs.createDouble(stamp, calendar, format);
                    break;
                    
                case 'object':
                    if (Object.keys(stamp).includes('start') && Object.keys(stamp).includes('end')) {
                        switch (typeof stamp.start) {
                            case 'string':
                                ref = globalThis.CalendsFuncs.createStringRange(stamp.start, stamp.end, calendar, format);
                                break;

                            case 'number':
                                if (Number.isInteger(stamp.start))
                                    ref = globalThis.CalendsFuncs.createInt64Range(stamp.start, stamp.end, calendar, format);
                                else
                                    ref = globalThis.CalendsFuncs.createDoubleRange(stamp.start, stamp.end, calendar, format);
                                break;

                            default:
                                throw new CalendsError(`Unsupported timestamp type: ${typeof stamp.start}`);
                        }
                    } else if (Object.keys(stamp).includes('start') && Object.keys(stamp).includes('duration')) {
                        switch (typeof stamp.start) {
                            case 'string':
                                ref = globalThis.CalendsFuncs.createStringStartPeriod(stamp.start, stamp.duration, calendar, format);
                                break;

                            case 'number':
                                if (Number.isInteger(stamp.start))
                                    ref = globalThis.CalendsFuncs.createInt64StartPeriod(stamp.start, stamp.duration, calendar, format);
                                else
                                    ref = globalThis.CalendsFuncs.createDoubleStartPeriod(stamp.start, stamp.duration, calendar, format);
                                break;

                            default:
                                throw new CalendsError(`Unsupported timestamp type: ${typeof stamp.start}`);
                        }
                    } else if (Object.keys(stamp).includes('duration') && Object.keys(stamp).includes('end')) {
                        switch (typeof stamp.end) {
                            case 'string':
                                ref = globalThis.CalendsFuncs.createStringEndPeriod(stamp.duration, stamp.end, calendar, format);
                                break;

                            case 'number':
                                if (Number.isInteger(stamp.end))
                                    ref = globalThis.CalendsFuncs.createInt64EndPeriod(stamp.duration, stamp.end, calendar, format);
                                else
                                    ref = globalThis.CalendsFuncs.createDoubleEndPeriod(stamp.duration, stamp.end, calendar, format);
                                break;

                            default:
                                throw new CalendsError(`Unsupported timestamp type: ${typeof stamp.end}`);
                        }
                    } else {
                        throw new CalendsError(`Unsupported timestamp type: ${JSON.stringify(stamp)}`);
                    }
                    break;

                default:
                    throw new CalendsError(`Unsupported timestamp type: ${typeof stamp}`);
            }
        }

        this.#instance = onDestroy(() => ({ ref }), () => { if (!halt) globalThis.CalendsFuncs.release(ref); });
    }

    static fromText(stamp) {
        return new Calends(BigInt(globalThis.CalendsFuncs.decodeText(stamp)));
    }

    static fromJson(stamp) {
        return new Calends(BigInt(globalThis.CalendsFuncs.decodeJson(stamp)));
    }

    // For Use With JSON.decode()

    static reviver(key, value) {
        if ((typeof value == 'string') || (typeof value == 'object' && typeof value.start == 'string' && typeof value.end == 'string')) {
            return new Calends(value, 'tai64', 'tai64narux');
        }

        return value;
    }

    // READ

    date(calendar, format) {
        return globalThis.CalendsFuncs.date(this.#instance.ref, calendar, format);
    }

    duration() {
        return globalThis.CalendsFuncs.duration(this.#instance.ref);
    }

    endDate(calendar, format) {
        return globalThis.CalendsFuncs.endDate(this.#instance.ref, calendar, format);
    }

    toText() {
        return globalThis.CalendsFuncs.encodeText(this.#instance.ref);
    }

    toJson() {
        return globalThis.CalendsFuncs.encodeJson(this.#instance.ref);
    }

    // Treated Specially By JS

    toString() {
        return globalThis.CalendsFuncs.string(this.#instance.ref);
    }

    toJSON() {
        return JSON.parse(this.toJson());
    }

    // "UPDATE" (actually returns new object)

    add(offset, calendar) {
        let ret;

        switch (typeof offset) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.addString(this.#instance.ref, offset, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.addInt64(this.#instance.ref, offset, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.addDouble(this.#instance.ref, offset, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported offset type');
        }

        return ret;
    }

    addFromEnd(offset, calendar) {
        let ret;

        switch (typeof offset) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.addFromEndString(this.#instance.ref, offset, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.addFromEndInt64(this.#instance.ref, offset, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.addFromEndDouble(this.#instance.ref, offset, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported offset type');
        }

        return ret;
    }

    subtract(offset, calendar) {
        let ret;

        switch (typeof offset) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.subtractString(this.#instance.ref, offset, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.subtractInt64(this.#instance.ref, offset, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.subtractDouble(this.#instance.ref, offset, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported offset type');
        }

        return ret;
    }

    subtractFromEnd(offset, calendar) {
        let ret;

        switch (typeof offset) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.subtractFromEndString(this.#instance.ref, offset, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.subtractFromEndInt64(this.#instance.ref, offset, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.subtractFromEndDouble(this.#instance.ref, offset, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported offset type');
        }

        return ret;
    }

    next(offset, calendar) {
        let ret;

        switch (typeof offset) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.nextString(this.#instance.ref, offset, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.nextInt64(this.#instance.ref, offset, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.nextDouble(this.#instance.ref, offset, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported offset type');
        }

        return ret;
    }

    previous(offset, calendar) {
        let ret;

        switch (typeof offset) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.previousString(this.#instance.ref, offset, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.previousInt64(this.#instance.ref, offset, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.previousDouble(this.#instance.ref, offset, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported offset type');
        }

        return ret;
    }

    withDate(stamp, calendar, format) {
        let ret;

        switch (typeof stamp) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDateString(this.#instance.ref, stamp, calendar, format)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDateInt64(this.#instance.ref, stamp, calendar, format)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDateDouble(this.#instance.ref, stamp, calendar, format)));
                break;

            default:
                throw new CalendsError('Unsupported stamp type');
        }

        return ret;
    }

    withEndDate(stamp, calendar, format) {
        let ret;

        switch (typeof stamp) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withEndDateString(this.#instance.ref, stamp, calendar, format)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withEndDateInt64(this.#instance.ref, stamp, calendar, format)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withEndDateDouble(this.#instance.ref, stamp, calendar, format)));
                break;

            default:
                throw new CalendsError('Unsupported stamp type');
        }

        return ret;
    }

    withDuration(duration, calendar) {
        let ret;

        switch (typeof duration) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDurationString(this.#instance.ref, duration, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDurationInt64(this.#instance.ref, duration, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDurationDouble(this.#instance.ref, duration, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported duration type');
        }

        return ret;
    }

    withDurationFromEnd(duration, calendar) {
        let ret;

        switch (typeof duration) {
            case 'string':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDurationFromEndString(this.#instance.ref, duration, calendar)));
                break;

            case 'bigint':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDurationFromEndInt64(this.#instance.ref, duration, calendar)));
                break;

            case 'number':
                ret = new Calends(BigInt(globalThis.CalendsFuncs.withDurationFromEndDouble(this.#instance.ref, duration, calendar)));
                break;

            default:
                throw new CalendsError('Unsupported duration type');
        }

        return ret;
    }

    merge(other) {
        if (other instanceof Calends)
            return new Calends(BigInt(globalThis.CalendsFuncs.merge(this.#instance.ref, other.#instance.ref)));

        throw new CalendsError('Merge requires a Calends object as an argument');
    }

    intersect(other) {
        if (other instanceof Calends)
            return new Calends(BigInt(globalThis.CalendsFuncs.intersect(this.#instance.ref, other.#instance.ref)));

        throw new CalendsError('Intersect requires a Calends object as an argument');
    }

    gap(other) {
        if (other instanceof Calends)
            return new Calends(BigInt(globalThis.CalendsFuncs.gap(this.#instance.ref, other.#instance.ref)));

        throw new CalendsError('Gap requires a Calends object as an argument');
    }

    // COMPARE

    difference(other, mode) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.difference(this.#instance.ref, other.#instance.ref, mode);

        throw new CalendsError('Difference requires a Calends object as a first argument');
    }

    compare(other, mode) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.compare(this.#instance.ref, other.#instance.ref, mode);

        throw new CalendsError('Compare requires a Calends object as a first argument');
    }

    isSame(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isSame(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsSame requires a Calends object as an argument');
    }

    isSameDuration(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isSameDuration(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsSameDuration requires a Calends object as an argument');
    }

    isShorter(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isShorter(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsShorter requires a Calends object as an argument');
    }

    isLonger(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isLonger(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsLonger requires a Calends object as an argument');
    }

    isBefore(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isBefore(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsBefore requires a Calends object as an argument');
    }

    isDuring(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isDuring(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsDuring requires a Calends object as an argument');
    }

    isAfter(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.isAfter(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('IsAfter requires a Calends object as an argument');
    }

    startsBefore(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.startsBefore(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('StartsBefore requires a Calends object as an argument');
    }

    startsDuring(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.startsDuring(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('StartsDuring requires a Calends object as an argument');
    }

    startsAfter(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.startsAfter(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('StartsAfter requires a Calends object as an argument');
    }

    endsBefore(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.endsBefore(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('EndsBefore requires a Calends object as an argument');
    }

    endsDuring(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.endsDuring(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('EndsDuring requires a Calends object as an argument');
    }

    endsAfter(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.endsAfter(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('EndsAfter requires a Calends object as an argument');
    }

    contains(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.contains(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('Contains requires a Calends object as an argument');
    }

    overlaps(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.overlaps(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('Overlaps requires a Calends object as an argument');
    }

    abuts(other) {
        if (other instanceof Calends)
            return globalThis.CalendsFuncs.abuts(this.#instance.ref, other.#instance.ref);

        throw new CalendsError('Abuts requires a Calends object as an argument');
    }
}

class TAI64Time {
    seconds;
    nano;
    atto;
    ronto;
    udecto;
    xindecto;

    // CREATE

    constructor({ seconds, nano, atto, ronto, udecto, xindecto } = {}) {
        this.seconds = seconds || 0;
        this.nano = nano || 0;
        this.atto = atto || 0;
        this.ronto = ronto || 0;
        this.udecto = udecto || 0;
        this.xindecto = xindecto || 0;
    }

    static fromString(raw) {
        return new TAI64Time(globalThis.CalendsFuncs.taiFromString(raw));
    }

    static fromHex(raw) {
        return new TAI64Time(globalThis.CalendsFuncs.taiFromHexString(raw));
    }

    static fromNumber(raw) {
        return new TAI64Time(globalThis.CalendsFuncs.taiFromDouble(raw));
    }

    static fromText(raw) {
        return new TAI64Time(globalThis.CalendsFuncs.taiDecodeText(raw));
    }

    // READ

    toString() {
        return globalThis.CalendsFuncs.taiString(this);
    }

    toHex() {
        return globalThis.CalendsFuncs.taiHexString(this);
    }

    toNumber() {
        return globalThis.CalendsFuncs.taiDouble(this);
    }

    toText() {
        return globalThis.CalendsFuncs.taiEncodeText(this);
    }

    // "UPDATE" (returns new object)

    add(instant) {
        return new TAI64Time(globalThis.CalendsFuncs.taiAdd(this, instant));
    }

    sub(instant) {
        return new TAI64Time(globalThis.CalendsFuncs.taiSub(this, instant));
    }

    toUTC() {
        return new TAI64Time(globalThis.CalendsFuncs.taiToUtc(this));
    }

    fromUTC() {
        return new TAI64Time(globalThis.CalendsFuncs.taiFromUtc(this));
    }
}

class CalendarDefinition {
    // SET ME
    name;
    // SET ME
    defaultFormat;

    #toInternalStringWrap(name, stamp, format) {
        return this.toInternal(stamp, format);
    }

    #toInternalBigIntWrap(name, stamp, format) {
        return this.toInternal(stamp, format);
    }

    #toInternalDoubleWrap(name, stamp, format) {
        return this.toInternal(stamp, format);
    }

    #toInternalTaiWrap(name, stamp) {
        return this.toInternal(new TAI64Time(stamp));
    }

    #fromInternalWrap(name, instant, format) {
        return this.fromInternal(new TAI64Time(instant), format);
    }

    #offsetStringWrap(name, instant, offset) {
        return this.offset(new TAI64Time(instant), offset);
    }

    #offsetBigIntWrap(name, instant, offset) {
        return this.offset(new TAI64Time(instant), offset);
    }

    #offsetDoubleWrap(name, instant, offset) {
        return this.offset(new TAI64Time(instant), offset);
    }

    #offsetTaiWrap(name, instant, offset) {
        return this.offset(new TAI64Time(instant), new TAI64Time(offset));
    }

    // IMPLEMENT ME
    constructor() {
        if (this.constructor === CalendarDefinition)
            throw new CalendsError("Don't directly create a CalendarDefinition object - use a subclass instead");
    }

    // IMPLEMENT ME
    toInternal(stamp, format) {
        throw new CalendsError('toInternal NOT IMPLEMENTED - Something is very wrong');
    }

    // IMPLEMENT ME
    fromInternal(instant, format) {
        throw new CalendsError('fromInternal NOT IMPLEMENTED - Something is very wrong');
    }

    // IMPLEMENT ME
    offset(instant, offset) {
        throw new CalendsError('offset NOT IMPLEMENTED - Something is very wrong');
    }

    register() {
        globalThis.CalendsFuncs.calendarRegister(
            this.name, this.defaultFormat,
            this.#toInternalStringWrap, this.#toInternalBigIntWrap, this.#toInternalDoubleWrap, this.#toInternalTaiWrap,
            this.#fromInternalWrap,
            this.#offsetStringWrap, this.#offsetBigIntWrap, this.#offsetDoubleWrap, this.#offsetTaiWrap,
        );
    }

    static registered() {
        return String(globalThis.CalendsFuncs.calendarListRegistered()).split('\n');
    }

    static checkRegistered(name) {
        return globalThis.CalendsFuncs.calendarRegistered(name);
    }

    isRegistered() {
        return globalThis.CalendsFuncs.calendarRegistered(this.name);
    }

    unregister() {
        globalThis.CalendsFuncs.calendarUnregister(this.name);
    }
}

if (ENVIRONMENT_IS_NODE) {
    module.exports = {
        Calends,
        CalendsError,
        CalendarDefinition,
        TAI64Time,
        initCalends,
        stopCalends,
    }
}
