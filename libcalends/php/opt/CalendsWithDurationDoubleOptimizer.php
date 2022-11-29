<?php

namespace Zephir\Optimizers\FunctionCall;

use Zephir\Call;
use Zephir\CompilationContext;
use Zephir\CompiledExpression;
use Zephir\Optimizers\OptimizerAbstract;
use Zephir\Compiler\CompilerException;


class CalendsWithDurationDoubleOptimizer extends OptimizerAbstract
{
    public function optimize(array $expression, Call $call, CompilationContext $context)
    {
        if (count($expression['parameters']) != 3) {
            throw new CompilerException("'Calends_with_duration_double' only accepts 3 parameters", $expression);
        }

        /**
         * Process the expected symbol to be returned
         */
        $call->processExpectedReturn($context);

        $symbolVariable = $call->getSymbolVariable();
        if (!$symbolVariable->isDouble()) {
            throw new CompilerException("Calends objects are identified by 'double' type values", $expression);
        }

        $context->headersManager->add('wrap_libcalends');
        $resolvedParams = $call->getReadOnlyResolvedParams($expression['parameters'], $context, $expression);

        return new CompiledExpression('double', 'ext_Calends_with_duration_double(' . implode(', ', $resolvedParams) . ')', $expression);
    }
}
