<?php

namespace Calends;

final class TAITime
{
    private static ?\FFI $ffi = null;

    public int $seconds = 0;
    public int $nano = 0;
    public int $atto = 0;
    public int $ronto = 0;
    public int $udecto = 0;
    public int $xindecto = 0;

    private static function ffiInit()
    {
        if (is_null(self::$ffi)) {
            self::$ffi = \FFI::scope("CALENDS");
        }
    }

    public function __construct()
    {
        self::ffiInit();
    }

    public static function fromNative(\FFI\CData $stamp): self
    {
        $out = new self();

        $out->seconds = $stamp->Seconds;
        $out->nano = $stamp->Nano;
        $out->atto = $stamp->Atto;
        $out->ronto = $stamp->Ronto;
        $out->udecto = $stamp->Udecto;
        $out->xindecto = $stamp->Xindecto;

        return $out;
    }

    public function toNative(): \FFI\CData
    {
        $out = self::$ffi->new('TAI64Time');

        $out->Seconds = $this->seconds;
        $out->Nano = $this->nano;
        $out->Atto = $this->atto;
        $out->Ronto = $this->ronto;
        $out->Udecto = $this->udecto;
        $out->Xindecto = $this->xindecto;

        return $out;
    }

    public static function fromString(string $stamp): self
    {
        self::ffiInit();

        return self::fromNative(self::$ffi->TAI64Time_from_string($stamp));
    }

    public function toString(): string
    {
        return \FFI::string(self::$ffi->TAI64Time_string($this->toNative()));
    }

    public function __toString(): string
    {
        return $this->toString();
    }

    public static function fromHex(string $stamp): self
    {
        self::ffiInit();

        return self::fromNative(self::$ffi->TAI64Time_from_hex_string($stamp));
    }

    public function toHex(): string
    {
        return \FFI::string(self::$ffi->TAI64Time_hex_string($this->toNative()));
    }

    public static function fromNumber(float $stamp): self
    {
        self::ffiInit();

        return self::fromNative(self::$ffi->TAI64Time_from_double($stamp));
    }

    public function toNumber(): float
    {
        return (float)self::$ffi->TAI64Time_double($this->toNative());
    }

    public function add(TAITime $other): TAITime
    {
        return self::fromNative(self::$ffi->TAI64Time_add($this->toNative(), $other->toNative()));
    }

    public function sub(TAITime $other): TAITime
    {
        return self::fromNative(self::$ffi->TAI64Time_sub($this->toNative(), $other->toNative()));
    }

    public function fromUTC(): TAITime
    {
        return self::fromNative(self::$ffi->TAI64Time_utc_to_tai($this->toNative()));
    }

    public function toUTC(): TAITime
    {
        return self::fromNative(self::$ffi->TAI64Time_tai_to_utc($this->toNative()));
    }
}
