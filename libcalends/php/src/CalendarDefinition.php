<?php

namespace Calends;

abstract class CalendarDefinition
{
    protected static ?\FFI $ffi = null;

    protected string $name;

    protected string $defaultFormat;

    abstract public function toInternal($stamp, string $format = ''): TAITime;

    abstract public function fromInternal(TAITime $stamp, string $format = ''): string;

    abstract public function offset($stamp, string $offset): TAITime;

    protected static function ffiInit()
    {
        if (is_null(self::$ffi)) {
            self::$ffi = \FFI::scope("CALENDS");
        }
    }

    public function __construct(string $name, string $defaultFormat = '')
    {
        self::ffiInit();

        $this->name = $name;
        $this->defaultFormat = $defaultFormat;

        if (!$this->isRegistered()) {
            $toInternalString = function (\FFI\CData $name, \FFI\CData $stamp, \FFI\CData $format): \FFI\CData {
                $out = $this->toInternal(\FFI::string($stamp), \FFI::string($format))->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $toInternalInt = function (\FFI\CData $name, int $stamp, \FFI\CData $format): \FFI\CData {
                $out = $this->toInternal($stamp, \FFI::string($format))->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $toInternalFloat = function (\FFI\CData $name, float $stamp, \FFI\CData $format): \FFI\CData {
                $out = $this->toInternal($stamp, \FFI::string($format))->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $toInternalTAI = function (\FFI\CData $name, \FFI\CData $stamp): \FFI\CData {
                $out = $this->toInternal(TAITime::fromNative($stamp))->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $fromInternalFunc = function (\FFI\CData $name, \FFI\CData $stamp, \FFI\CData $format): \FFI\CData {
                $tmp = $this->fromInternal(TAITime::fromNative($stamp), \FFI::string($format));
                $len = strlen($tmp) + 1;

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                $out = \FFI::new("char[{$len}]", 0);
                \FFI::memcpy($out, $tmp, $len - 1);

                return $out;
            };

            $offsetString = function (\FFI\CData $name, \FFI\CData $stamp, \FFI\CData $offset): \FFI\CData {
                $out = $this->offset(TAITime::fromNative($stamp), \FFI::string($offset))->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $offsetInt = function (\FFI\CData $name, \FFI\CData $stamp, int $offset): \FFI\CData {
                $out = $this->offset(TAITime::fromNative($stamp), $offset)->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $offsetFloat = function (\FFI\CData $name, \FFI\CData $stamp, float $offset): \FFI\CData {
                $out = $this->offset(TAITime::fromNative($stamp), $offset)->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            $offsetTAI = function (\FFI\CData $name, \FFI\CData $stamp, \FFI\CData $offset): \FFI\CData {
                $out = $this->offset(TAITime::fromNative($stamp), TAITime::fromNative($offset))->toNative();

                if (!empty(CalendsException::current())) {
                    throw CalendsException::current();
                }

                return $out;
            };

            self::$ffi->Calends_calendar_register(
                $this->name, $this->defaultFormat,
                $toInternalString, $toInternalInt, $toInternalFloat, $toInternalTAI,
                $fromInternalFunc,
                $offsetString, $offsetInt, $offsetFloat, $offsetTAI
            );
        }
    }

    public function __destruct()
    {
        $this->unregister();
    }

    public static function isRegistered(string $name): bool
    {
        self::ffiInit();

        return (bool)self::$ffi->Calends_calendar_registered($name);
    }

    public static function listRegistered(): array
    {
        self::ffiInit();

        return explode("\n", \FFI::string(self::$ffi->Calends_calendar_list_registered()));
    }

    public function unregister()
    {
        self::$ffi->Calends_calendar_unregister($this->name);
    }
}
