<?php

namespace Zephir\Optimizers\FunctionCall;

use Zephir\Call;
use Zephir\CompilationContext;
use Zephir\CompiledExpression;
use Zephir\Optimizers\OptimizerAbstract;
use Zephir\Compiler\CompilerException;

class CalendsCalendarRegisterOptimizer extends OptimizerAbstract
{
    public function optimize(array $expression, Call $call, CompilationContext $context)
    {
        if (count($expression['parameters']) != 2) {
            throw new CompilerException("'Calends_calendar_register' only accepts 2 parameters", $expression);
        }

        $context->headersManager->add('wrap_libcalends');
        $resolvedParams = $call->getReadOnlyResolvedParams($expression['parameters'], $context, $expression);

        $context->codePrinter->output('ext_Calends_calendar_register(' . implode($resolvedParams, ', ') . ');');
        return new CompiledExpression('void', 'ext_Calends_calendar_register(' . implode($resolvedParams, ', ') . ')', $expression);
    }
}
