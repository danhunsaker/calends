{
  const IS_WEB = typeof window === 'object';
  const IS_WORKER = typeof importScripts === 'function';
  const HAS_NODE = typeof process === 'object' && typeof process.versions === 'object' && typeof process.versions.node === 'string';
  const IS_NODE = HAS_NODE && !IS_WEB && !IS_WORKER;
  const IS_SHELL = !IS_WEB && !IS_NODE && !IS_WORKER;

  if (IS_NODE) {
    globalThis.crypto = require('node:crypto');
    var chai = require('chai');
    ({ Calends, CalendarDefinition, TAI64Time, initCalends, stopCalends } = require('./calends'));
  }
  else {
    window.onload = () => {
      mocha.run();
    };
  }

  const beforeFunc = function () {
    this.timeout(5000);
    return new Promise((resolve) => {
      initCalends(resolve);
    });
  };

  const afterFunc = function () {
    this.timeout(5000);
    return new Promise((resolve) => {
      stopCalends(resolve);
    });
  }

  describe('JS / WASM Tests', function () {
    after(function () {
      if (IS_NODE || IS_SHELL) {
        process.exit();
      }
    });

    describe('Calends class', function () {
      before(beforeFunc);
      after(afterFunc);
  
      describe('constructor', function () {
        it('creates a valid Calends object', function () {
          chai.expect(new Calends('0', 'tai64', 'decimal')).instanceof(Calends);
          chai.expect(new Calends(0, 'unix', '')).instanceof(Calends);
          chai.expect(new Calends(0.1, 'unix', '')).instanceof(Calends);
  
          chai.expect(new Calends({ start: '0', end: '10' }, 'tai64', 'decimal')).instanceof(Calends);
          chai.expect(new Calends({ start: 0, end: 10 }, 'unix', '')).instanceof(Calends);
          chai.expect(new Calends({ start: 0.1, end: 9.9 }, 'unix', '')).instanceof(Calends);
  
          chai.expect(new Calends({ start: '0', duration: '10' }, 'tai64', 'decimal')).instanceof(Calends);
          chai.expect(new Calends({ start: 0, duration: 10 }, 'unix', '')).instanceof(Calends);
          chai.expect(new Calends({ start: 0.1, duration: 9.9 }, 'unix', '')).instanceof(Calends);
  
          chai.expect(new Calends({ duration: '0', end: '10' }, 'tai64', 'decimal')).instanceof(Calends);
          chai.expect(new Calends({ duration: 0, end: 10 }, 'unix', '')).instanceof(Calends);
          chai.expect(new Calends({ duration: 0.1, end: 9.9 }, 'unix', '')).instanceof(Calends);
        });
  
        it('throws when it fails to create a valid Calends object', function () {
          chai.expect(() => new Calends([], '', '')).to.throw('Unsupported timestamp type: []');
        });
      });
  
      describe('fromText', function () {
        it('creates a valid Calends object', function () {
          chai.expect(Calends.fromText('0')).instanceof(Calends);
        });
      });
  
      describe('fromJson', function () {
        it('creates a valid Calends object', function () {
          chai.expect(Calends.fromJson('{"start":"0","end":"0"}')).instanceof(Calends);
        });
      });
  
      describe('date', function () {
        it('outputs a valid date string', function () {
          const c = new Calends('0', 'tai64', 'decimal');
  
          chai.expect(c.date('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('duration', function () {
        it('outputs a valid duration string', function () {
          const c = new Calends('0', 'tai64', 'decimal');
  
          chai.expect(c.duration()).to.equal('0');
        });
      });
  
      describe('endDate', function () {
        it('outputs a valid date string', function () {
          const c = new Calends('0', 'tai64', 'decimal');
  
          chai.expect(c.endDate('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('toText', function () {
        it('outputs a valid text value', function () {
          const c = new Calends('0', 'tai64', 'decimal');
  
          chai.expect(c.toText()).to.equal('40000000000000000000000000000000000000000000000000000000');
        });
      });
  
      describe('toJson', function () {
        it('outputs a valid JSON value', function () {
          const c = new Calends('0', 'tai64', 'decimal');
  
          chai.expect(c.toJson()).to.equal('"40000000000000000000000000000000000000000000000000000000"');
        });
      });
  
      describe('add', function () {
        it('changes the start date by an offset', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.add('10', 'unix');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('10');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('addFromEnd', function () {
        it('changes the end date by an offset', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.addFromEnd('10', 'unix');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('subtract', function () {
        it('changes the start date by an offset', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.subtract('10', 'tai64');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('-10');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('subtractFromEnd', function () {
        it('changes the end date by an offset', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.subtractFromEnd('10', 'tai64');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('-10');
        });
      });
  
      describe('next', function () {
        it('gives the next date interval by an offset', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.next('10', 'tai64');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('previous', function () {
        it('gives the previous date interval by an offset', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.previous('10', 'tai64');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('-10');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('withDate', function () {
        it('sets the start date to a specific instant', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.withDate('10', 'tai64', 'decimal');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('10');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('withEndDate', function () {
        it('sets the end date to a specific instant', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.withEndDate('10', 'tai64', 'decimal');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('withDuration', function () {
        it('sets the duration to a specific length from the start', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.withDuration('10', 'tai64');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('withDurationFromEnd', function () {
        it('sets the duration to a specific length from the end', function () {
          const c = new Calends('0', 'tai64', 'decimal');
          const z = c.withDurationFromEnd('10', 'tai64', 'decimal');
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('-10');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('0');
        });
      });
  
      describe('merge', function () {
        it('merges two moments', function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends('10', 'tai64', 'decimal');
          const z = a.merge(b);
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('intersect', function () {
        it('finds the moment shared by two other moments', function () {
          const a = new Calends({start: '0', end: '10'}, 'tai64', 'decimal');
          const b = new Calends('10', 'tai64', 'decimal');
          const z = a.intersect(b);
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('10');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('gap', function () {
        it('finds the moment between two other moments', function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends('10', 'tai64', 'decimal');
          const z = a.gap(b);
  
          chai.expect(z).instanceof(Calends);
          chai.expect(z.date('tai64', 'decimal')).to.equal('0');
          chai.expect(z.endDate('tai64', 'decimal')).to.equal('10');
        });
      });
  
      describe('difference', function () {
        it("finds the difference between two moments' instants", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends('10', 'tai64', 'decimal');
  
          chai.expect(a.difference(b, 'start')).to.equal('-10');
          chai.expect(a.difference(b, 'start-end')).to.equal('-10');
          chai.expect(a.difference(b, 'end-start')).to.equal('-10');
          chai.expect(a.difference(b, 'end')).to.equal('-10');
          chai.expect(a.difference(b, 'duration')).to.equal('0');
  
          chai.expect(b.difference(a, 'start')).to.equal('10');
          chai.expect(b.difference(a, 'start-end')).to.equal('10');
          chai.expect(b.difference(a, 'end-start')).to.equal('10');
          chai.expect(b.difference(a, 'end')).to.equal('10');
          chai.expect(b.difference(a, 'duration')).to.equal('0');
        });
      });
  
      describe('compare', function () {
        it("compares two moments' instants", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends('10', 'tai64', 'decimal');
  
          chai.expect(a.compare(b, 'start')).to.equal(-1);
          chai.expect(a.compare(b, 'start-end')).to.equal(-1);
          chai.expect(a.compare(b, 'end-start')).to.equal(-1);
          chai.expect(a.compare(b, 'end')).to.equal(-1);
          chai.expect(a.compare(b, 'duration')).to.equal(0);
  
          chai.expect(b.compare(a, 'start')).to.equal(1);
          chai.expect(b.compare(a, 'start-end')).to.equal(1);
          chai.expect(b.compare(a, 'end-start')).to.equal(1);
          chai.expect(b.compare(a, 'end')).to.equal(1);
          chai.expect(b.compare(a, 'duration')).to.equal(0);
        });
      });
  
      describe('isSame', function () {
        it("checks whether two moments are identical", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends('10', 'tai64', 'decimal');
          const c = new Calends('10', 'tai64', 'decimal');
  
          chai.expect(a.isSame(a)).true;
          chai.expect(a.isSame(b)).false;
          chai.expect(a.isSame(c)).false;
  
          chai.expect(b.isSame(a)).false;
          chai.expect(b.isSame(b)).true;
          chai.expect(b.isSame(c)).true;
  
          chai.expect(c.isSame(a)).false;
          chai.expect(c.isSame(b)).true;
          chai.expect(c.isSame(c)).true;
        });
      });
  
      describe('isSameDuration', function () {
        it("checks whether two moments have the same duration", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.isSameDuration(a)).true;
          chai.expect(a.isSameDuration(b)).false;
          chai.expect(a.isSameDuration(c)).false;
  
          chai.expect(b.isSameDuration(a)).false;
          chai.expect(b.isSameDuration(b)).true;
          chai.expect(b.isSameDuration(c)).true;
  
          chai.expect(c.isSameDuration(a)).false;
          chai.expect(c.isSameDuration(b)).true;
          chai.expect(c.isSameDuration(c)).true;
        });
      });
  
      describe('isShorter', function () {
        it("checks whether one moment is shorter than another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.isShorter(a)).false;
          chai.expect(a.isShorter(b)).true;
          chai.expect(a.isShorter(c)).true;
  
          chai.expect(b.isShorter(a)).false;
          chai.expect(b.isShorter(b)).false;
          chai.expect(b.isShorter(c)).false;
  
          chai.expect(c.isShorter(a)).false;
          chai.expect(c.isShorter(b)).false;
          chai.expect(c.isShorter(c)).false;
        });
      });
  
      describe('isLonger', function () {
        it("checks whether one moment is longer than another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.isLonger(a)).false;
          chai.expect(a.isLonger(b)).false;
          chai.expect(a.isLonger(c)).false;
  
          chai.expect(b.isLonger(a)).true;
          chai.expect(b.isLonger(b)).false;
          chai.expect(b.isLonger(c)).false;
  
          chai.expect(c.isLonger(a)).true;
          chai.expect(c.isLonger(b)).false;
          chai.expect(c.isLonger(c)).false;
        });
      });
  
      describe('isBefore', function () {
        it("checks whether one moment is before another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.isBefore(a)).false;
          chai.expect(a.isBefore(b)).true;
          chai.expect(a.isBefore(c)).false;
  
          chai.expect(b.isBefore(a)).false;
          chai.expect(b.isBefore(b)).false;
          chai.expect(b.isBefore(c)).false;
  
          chai.expect(c.isBefore(a)).false;
          chai.expect(c.isBefore(b)).true;
          chai.expect(c.isBefore(c)).false;
        });
      });
  
      describe('isDuring', function () {
        it("checks whether one moment is during another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.isDuring(a)).true;
          chai.expect(a.isDuring(b)).false;
          chai.expect(a.isDuring(c)).true;
  
          chai.expect(b.isDuring(a)).false;
          chai.expect(b.isDuring(b)).true;
          chai.expect(b.isDuring(c)).false;
  
          chai.expect(c.isDuring(a)).false;
          chai.expect(c.isDuring(b)).false;
          chai.expect(c.isDuring(c)).true;
        });
      });
  
      describe('isAfter', function () {
        it("checks whether one moment is after another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.isAfter(a)).false;
          chai.expect(a.isAfter(b)).false;
          chai.expect(a.isAfter(c)).false;
  
          chai.expect(b.isAfter(a)).true;
          chai.expect(b.isAfter(b)).false;
          chai.expect(b.isAfter(c)).true;
  
          chai.expect(c.isAfter(a)).true;
          chai.expect(c.isAfter(b)).false;
          chai.expect(c.isAfter(c)).false;
        });
      });
  
      describe('startsBefore', function () {
        it("checks whether one moment starts before another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.startsBefore(a)).false;
          chai.expect(a.startsBefore(b)).true;
          chai.expect(a.startsBefore(c)).false;
  
          chai.expect(b.startsBefore(a)).false;
          chai.expect(b.startsBefore(b)).false;
          chai.expect(b.startsBefore(c)).false;
  
          chai.expect(c.startsBefore(a)).false;
          chai.expect(c.startsBefore(b)).true;
          chai.expect(c.startsBefore(c)).false;
        });
      });
  
      describe('startsDuring', function () {
        it("checks whether one moment starts during another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.startsDuring(a)).true;
          chai.expect(a.startsDuring(b)).false;
          chai.expect(a.startsDuring(c)).true;
  
          chai.expect(b.startsDuring(a)).false;
          chai.expect(b.startsDuring(b)).true;
          chai.expect(b.startsDuring(c)).false;
  
          chai.expect(c.startsDuring(a)).false;
          chai.expect(c.startsDuring(b)).false;
          chai.expect(c.startsDuring(c)).true;
        });
      });
  
      describe('startsAfter', function () {
        it("checks whether one moment starts after another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.startsAfter(a)).false;
          chai.expect(a.startsAfter(b)).false;
          chai.expect(a.startsAfter(c)).false;
  
          chai.expect(b.startsAfter(a)).true;
          chai.expect(b.startsAfter(b)).false;
          chai.expect(b.startsAfter(c)).true;
  
          chai.expect(c.startsAfter(a)).false;
          chai.expect(c.startsAfter(b)).false;
          chai.expect(c.startsAfter(c)).false;
        });
      });
  
      describe('endsBefore', function () {
        it("checks whether one moment ends before another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.endsBefore(a)).false;
          chai.expect(a.endsBefore(b)).true;
          chai.expect(a.endsBefore(c)).true;
  
          chai.expect(b.endsBefore(a)).false;
          chai.expect(b.endsBefore(b)).false;
          chai.expect(b.endsBefore(c)).false;
  
          chai.expect(c.endsBefore(a)).false;
          chai.expect(c.endsBefore(b)).true;
          chai.expect(c.endsBefore(c)).false;
        });
      });
  
      describe('endsDuring', function () {
        it("checks whether one moment ends during another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.endsDuring(a)).true;
          chai.expect(a.endsDuring(b)).false;
          chai.expect(a.endsDuring(c)).true;
  
          chai.expect(b.endsDuring(a)).false;
          chai.expect(b.endsDuring(b)).true;
          chai.expect(b.endsDuring(c)).false;
  
          chai.expect(c.endsDuring(a)).false;
          chai.expect(c.endsDuring(b)).false;
          chai.expect(c.endsDuring(c)).true;
        });
      });
  
      describe('endsAfter', function () {
        it("checks whether one moment ends after another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.endsAfter(a)).false;
          chai.expect(a.endsAfter(b)).false;
          chai.expect(a.endsAfter(c)).false;
  
          chai.expect(b.endsAfter(a)).true;
          chai.expect(b.endsAfter(b)).false;
          chai.expect(b.endsAfter(c)).true;
  
          chai.expect(c.endsAfter(a)).true;
          chai.expect(c.endsAfter(b)).false;
          chai.expect(c.endsAfter(c)).false;
        });
      });
  
      describe('contains', function () {
        it("checks whether one moment contains another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.contains(a)).true;
          chai.expect(a.contains(b)).false;
          chai.expect(a.contains(c)).false;
  
          chai.expect(b.contains(a)).false;
          chai.expect(b.contains(b)).true;
          chai.expect(b.contains(c)).false;
  
          chai.expect(c.contains(a)).true;
          chai.expect(c.contains(b)).false;
          chai.expect(c.contains(c)).true;
        });
      });
  
      describe('overlaps', function () {
        it("checks whether one moment overlaps with another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.overlaps(a)).true;
          chai.expect(a.overlaps(b)).false;
          chai.expect(a.overlaps(c)).true;
  
          chai.expect(b.overlaps(a)).false;
          chai.expect(b.overlaps(b)).true;
          chai.expect(b.overlaps(c)).false;
  
          chai.expect(c.overlaps(a)).true;
          chai.expect(c.overlaps(b)).false;
          chai.expect(c.overlaps(c)).true;
        });
      });
  
      describe('abuts', function () {
        it("checks whether one moment is next to another", function () {
          const a = new Calends('0', 'tai64', 'decimal');
          const b = new Calends({ start: '10', end: '20' }, 'tai64', 'decimal');
          const c = new Calends({ start: '0', end: '10' }, 'tai64', 'decimal');
  
          chai.expect(a.abuts(a)).false;
          chai.expect(a.abuts(b)).false;
          chai.expect(a.abuts(c)).false;
  
          chai.expect(b.abuts(a)).false;
          chai.expect(b.abuts(b)).false;
          chai.expect(b.abuts(c)).true;
  
          chai.expect(c.abuts(a)).false;
          chai.expect(c.abuts(b)).true;
          chai.expect(c.abuts(c)).false;
        });
      });
    });
  
    describe('TAI64TIme class', function () {
      before(beforeFunc);
      after(afterFunc);
  
      describe('constructor', function () {
        it('creates a valid TAI64Time object', function () {
          const t = new TAI64Time();
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.seconds).to.equal(0);
        });
      });
  
      describe('fromString', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromString('0');
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.seconds).to.equal(0);
        });
      });
  
      describe('fromHex', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromHex('0');
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.seconds).to.equal(-4611686018427388000);
        });
      });
  
      describe('fromNumber', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromNumber(0);
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.seconds).to.equal(0);
        });
      });
  
      describe('fromText', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromText('0');
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.seconds).to.equal(-4611686018427388000);
        });
      });
  
      describe('toString', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromString('0');
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.toString()).to.equal('0');
        });
      });
  
      describe('toHex', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromHex('0');
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.toHex()).to.equal('00000000000000000000000000000000000000000000000000000000');
        });
      });
  
      describe('toNumber', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromNumber(0);
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.toNumber()).to.equal(0);
        });
      });
  
      describe('toText', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromText('0');
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.toText()).to.equal('00000000000000000000000000000000000000000000000000000000');
        });
      });
  
      describe('add', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromNumber(0);
          const z = TAI64Time.fromNumber(16);
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(z).instanceof(TAI64Time);
          chai.expect(t.add(z).toHex()).to.equal('40000000000000100000000000000000000000000000000000000000');
        });
      });
  
      describe('sub', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromNumber(0);
          const z = TAI64Time.fromNumber(16);
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(z).instanceof(TAI64Time);
          chai.expect(t.sub(z).toHex()).to.equal('3FFFFFFFFFFFFFF00000000000000000000000000000000000000000');
        });
      });
  
      describe('fromUTC', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromNumber(0);
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.fromUTC().toHex()).to.equal('3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046');
        });
      });
  
      describe('toUTC', function () {
        it('creates a valid TAI64Time object', function () {
          const t = TAI64Time.fromNumber(0);
  
          chai.expect(t).instanceof(TAI64Time);
          chai.expect(t.toUTC().toHex()).to.equal('40000000000000080001404F3B9AC633119BB28532DB8E0D3399E120');
        });
      });
    });
  
    describe('CalendarDefinition interface (abstract class, really)', function () {
      before(beforeFunc);
      after(afterFunc);
  
      it('constructor (should fail)', function () {
        chai.expect(() => new CalendarDefinition()).throws("Don't directly create a CalendarDefinition object - use a subclass instead")
      });
  
      it('checkRegistered', function () {
        chai.expect(CalendarDefinition.checkRegistered('unix')).true;
        chai.expect(CalendarDefinition.checkRegistered('invalid')).false;
      });
  
      it('registered', function () {
        chai.expect(CalendarDefinition.registered()).to.deep.equal(['Gregorian', 'Jdc', 'Stardate', 'Tai64', 'Unix']);
      });
    });
  
    describe('FakeCalendar class', function () {
      class FakeCalendar extends CalendarDefinition {
        name = 'fake';
        defaultFormat = '';
  
        toInternal(stamp, format) {
          return new TAI64Time({ seconds: 0 });
        }
  
        fromInternal(instant, format) {
          return `${this.name}:0:${format}`;
        }
  
        offset(instant, offset) {
          return instant;
        }
      }
  
      before(beforeFunc);
      after(afterFunc);
  
      it('constructor (should succeed)', function () {
        const c = new FakeCalendar();
  
        chai.expect(c).instanceof(CalendarDefinition);
        chai.expect(c).instanceof(FakeCalendar);
      });
  
      it('register/unregister/isRegistered', function () {
        const c = new FakeCalendar();
  
        chai.expect(c).instanceof(CalendarDefinition);
        chai.expect(c).instanceof(FakeCalendar);
        chai.expect(FakeCalendar.checkRegistered(c.name)).false;
        chai.expect(c.isRegistered()).false;
  
        c.register();
  
        chai.expect(FakeCalendar.checkRegistered(c.name)).true;
        chai.expect(c.isRegistered()).true;
  
        c.unregister();
  
        chai.expect(FakeCalendar.checkRegistered(c.name)).false;
        chai.expect(c.isRegistered()).false;
      });
  
      it('toInternal', function () {
        const c = new FakeCalendar();
  
        chai.expect(c).instanceof(CalendarDefinition);
        chai.expect(c).instanceof(FakeCalendar);
  
        chai.expect(c.toInternal('', '')).instanceof(TAI64Time);
      });
  
      it('fromInternal', function () {
        const c = new FakeCalendar();
  
        chai.expect(c).instanceof(CalendarDefinition);
        chai.expect(c).instanceof(FakeCalendar);
  
        chai.expect(c.fromInternal(new TAI64Time(), '')).to.equal('fake:0:');
      });
  
      it('offset', function () {
        const c = new FakeCalendar();
  
        chai.expect(c).instanceof(CalendarDefinition);
        chai.expect(c).instanceof(FakeCalendar);
  
        chai.expect(c.offset(new TAI64Time(), '')).instanceof(TAI64Time);
      });
    });
  });
}
