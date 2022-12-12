<?php

namespace Calends;

class CalendsException extends \Exception {
    protected static ?self $current = null;

    public static function current(): ?self {
        $out = self::$current;

        self::$current = null;

        return $out;
    }

    public function __construct($message = '', $code = 0)
    {
        parent::__construct($message, $code, self::$current);

        self::$current = $this;
    }
}

\FFI::scope('CALENDS')->Calends_register_panic_handler(function(\FFI\CData $message): void {
    new CalendsException(\FFI::string($message), crc32(\FFI::string($message)));
});
