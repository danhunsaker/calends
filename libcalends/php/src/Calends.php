<?php

namespace Calends;

require_once('CalendsException.php');

class Calends implements \JsonSerializable
{
    private static ?\FFI $ffi = null;

    private ?int $instance = null;

    private static function ffiInit(): void
    {
        if (is_null(self::$ffi)) {
            self::$ffi = \FFI::scope("CALENDS");
        }
    }

    private function __construct(int $instance)
    {
        if (empty($instance)) {
            throw CalendsException::current();
        }

        $this->instance = $instance;
    }

    public function __destruct()
    {
        if (!empty($this->instance)) {
            self::$ffi->Calends_release($this->instance);
        } elseif (!empty(CalendsException::current())) {
            throw CalendsException::current();
        }
    }

    // CREATE

    public static function create($stamp = '', string $calendar = '', string $format = ''): self
    {
        self::ffiInit();

        if (is_string($stamp)) {
            return new self((int)self::$ffi->Calends_create_string($stamp, $calendar, $format));
        } elseif (is_int($stamp)) {
            return new self((int)self::$ffi->Calends_create_long_long($stamp, $calendar, $format));
        } elseif (is_float($stamp)) {
            return new self((int)self::$ffi->Calends_create_double($stamp, $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('start', $stamp) && array_key_exists('end', $stamp) && is_string($stamp['start'])) {
            return new self((int)self::$ffi->Calends_create_string_range($stamp['start'], $stamp['end'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('start', $stamp) && array_key_exists('duration', $stamp) && is_string($stamp['start'])) {
            return new self((int)self::$ffi->Calends_create_string_start_period($stamp['start'], $stamp['duration'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('duration', $stamp) && array_key_exists('end', $stamp) && is_string($stamp['end'])) {
            return new self((int)self::$ffi->Calends_create_string_end_period($stamp['duration'], $stamp['end'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('start', $stamp) && array_key_exists('end', $stamp) && is_int($stamp['start'])) {
            return new self((int)self::$ffi->Calends_create_long_long_range($stamp['start'], $stamp['end'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('start', $stamp) && array_key_exists('duration', $stamp) && is_int($stamp['start'])) {
            return new self((int)self::$ffi->Calends_create_long_long_start_period($stamp['start'], $stamp['duration'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('duration', $stamp) && array_key_exists('end', $stamp) && is_int($stamp['end'])) {
            return new self((int)self::$ffi->Calends_create_long_long_end_period($stamp['duration'], $stamp['end'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('start', $stamp) && array_key_exists('end', $stamp) && is_float($stamp['start'])) {
            return new self((int)self::$ffi->Calends_create_double_range($stamp['start'], $stamp['end'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('start', $stamp) && array_key_exists('duration', $stamp) && is_float($stamp['start'])) {
            return new self((int)self::$ffi->Calends_create_double_start_period($stamp['start'], $stamp['duration'], $calendar, $format));
        } elseif (is_array($stamp) && array_key_exists('duration', $stamp) && array_key_exists('end', $stamp) && is_float($stamp['end'])) {
            return new self((int)self::$ffi->Calends_create_double_end_period($stamp['duration'], $stamp['end'], $calendar, $format));
        } else {
            throw new CalendsException('Unsupported type');
        }
    }

    // READ

    public function date(string $calendar = '', string $format = ''): string
    {
        return \FFI::string(self::$ffi->Calends_date($this->instance, $calendar, $format));
    }

    public function duration(): string
    {
        return \FFI::string(self::$ffi->Calends_duration($this->instance));
    }

    public function endDate(string $calendar = '', string $format = ''): string
    {
        return \FFI::string(self::$ffi->Calends_end_date($this->instance, $calendar, $format));
    }

    // UPDATE

    public function add($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_add_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_add_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_add_double($this->instance, $offset, $calendar));
        }
    }

    public function subtract($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_subtract_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_subtract_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_subtract_double($this->instance, $offset, $calendar));
        }
    }

    public function addFromEnd($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_add_from_end_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_add_from_end_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_add_from_end_double($this->instance, $offset, $calendar));
        }
    }

    public function subtractFromEnd($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_subtract_from_end_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_subtract_from_end_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_subtract_from_end_double($this->instance, $offset, $calendar));
        }
    }

    public function next($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_next_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_next_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_next_double($this->instance, $offset, $calendar));
        }
    }

    public function previous($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_previous_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_previous_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_previous_double($this->instance, $offset, $calendar));
        }
    }

    public function withDate($stamp = '', string $calendar = '', string $format = ''): self
    {
        if (is_string($stamp)) {
            return new self(self::$ffi->Calends_with_date_string($this->instance, $stamp, $calendar, $format));
        } elseif (is_int($stamp)) {
            return new self(self::$ffi->Calends_with_date_long_long($this->instance, $stamp, $calendar, $format));
        } elseif (is_float($stamp)) {
            return new self(self::$ffi->Calends_with_date_double($this->instance, $stamp, $calendar, $format));
        }
    }

    public function withEndDate($stamp = '', string $calendar = '', string $format = ''): self
    {
        if (is_string($stamp)) {
            return new self(self::$ffi->Calends_with_end_date_string($this->instance, $stamp, $calendar, $format));
        } elseif (is_int($stamp)) {
            return new self(self::$ffi->Calends_with_end_date_long_long($this->instance, $stamp, $calendar, $format));
        } elseif (is_float($stamp)) {
            return new self(self::$ffi->Calends_with_end_date_double($this->instance, $stamp, $calendar, $format));
        }
    }

    public function withDuration($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_with_duration_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_with_duration_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_with_duration_double($this->instance, $offset, $calendar));
        }
    }

    public function withDurationFromEnd($offset = '', string $calendar = ''): self
    {
        if (is_string($offset)) {
            return new self(self::$ffi->Calends_with_duration_from_end_string($this->instance, $offset, $calendar));
        } elseif (is_int($offset)) {
            return new self(self::$ffi->Calends_with_duration_from_end_long_long($this->instance, $offset, $calendar));
        } elseif (is_float($offset)) {
            return new self(self::$ffi->Calends_with_duration_from_end_double($this->instance, $offset, $calendar));
        }
    }

    public function merge(self $other): self
    {
        return new self(self::$ffi->Calends_merge($this->instance, $other->instance));
    }

    public function intersect(self $other): self
    {
        return new self(self::$ffi->Calends_intersect($this->instance, $other->instance));
    }

    public function gap(self $other): self
    {
        return new self(self::$ffi->Calends_gap($this->instance, $other->instance));
    }

    // COMPARISON

    public function difference(self $other, string $mode = ''): string
    {
        return \FFI::string(self::$ffi->Calends_difference($this->instance, $other->instance, $mode));
    }

    public function compare(self $other, string $mode = ''): int
    {
        return (int)self::$ffi->Calends_compare($this->instance, $other->instance, $mode);
    }

    public function isSame(self $other): bool
    {
        return (int)self::$ffi->Calends_is_same($this->instance, $other->instance);
    }

    public function isSameDuration(self $other): bool
    {
        return (int)self::$ffi->Calends_is_same_duration($this->instance, $other->instance);
    }

    public function isShorter(self $other): bool
    {
        return (int)self::$ffi->Calends_is_shorter($this->instance, $other->instance);
    }

    public function isLonger(self $other): bool
    {
        return (int)self::$ffi->Calends_is_longer($this->instance, $other->instance);
    }

    public function contains(self $other): bool
    {
        return (int)self::$ffi->Calends_contains($this->instance, $other->instance);
    }

    public function overlaps(self $other): bool
    {
        return (int)self::$ffi->Calends_overlaps($this->instance, $other->instance);
    }

    public function abuts(self $other): bool
    {
        return (int)self::$ffi->Calends_abuts($this->instance, $other->instance);
    }

    public function isBefore(self $other): bool
    {
        return (int)self::$ffi->Calends_is_before($this->instance, $other->instance);
    }

    public function startsBefore(self $other): bool
    {
        return (int)self::$ffi->Calends_starts_before($this->instance, $other->instance);
    }

    public function endsBefore(self $other): bool
    {
        return (int)self::$ffi->Calends_ends_before($this->instance, $other->instance);
    }

    public function isDuring(self $other): bool
    {
        return (int)self::$ffi->Calends_is_during($this->instance, $other->instance);
    }

    public function startsDuring(self $other): bool
    {
        return (int)self::$ffi->Calends_starts_during($this->instance, $other->instance);
    }

    public function endsDuring(self $other): bool
    {
        return (int)self::$ffi->Calends_ends_during($this->instance, $other->instance);
    }

    public function isAfter(self $other): bool
    {
        return (int)self::$ffi->Calends_is_after($this->instance, $other->instance);
    }

    public function startsAfter(self $other): bool
    {
        return (int)self::$ffi->Calends_starts_after($this->instance, $other->instance);
    }

    public function endsAfter(self $other): bool
    {
        return (int)self::$ffi->Calends_ends_after($this->instance, $other->instance);
    }

    // UTILITY

    public function __toString(): string
    {
        return \FFI::string(self::$ffi->Calends_string($this->instance));
    }

    public static function jsonUnserialize(string $json): self
    {
        self::ffiInit();

        return new self((int)self::$ffi->Calends_decode_json($json));
    }

    public function jsonSerialize(): array
    {
        try {
            $out = json_decode(\FFI::string(self::$ffi->Calends_encode_json($this->instance)), true);
        } catch (\TypeError $typeError) {
            $exception = CalendsException::current();

            if (empty($exception)) {
                throw $typeError;
            }

            throw $exception;
        }

        return is_array($out) ? $out : ['start' => $out, 'end' => $out];
    }

    public function __unserialize(array $text): void
    {
        self::ffiInit();

        $this->instance = ((int)self::$ffi->Calends_decode_text($text[0]));
    }

    public function __serialize(): array
    {
        return [\FFI::string(self::$ffi->Calends_encode_text($this->instance))];
    }

    public function __debugInfo(): array
    {
        return $this->jsonSerialize();
    }
}
