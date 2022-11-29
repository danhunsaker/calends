<?php

namespace Zephir\Optimizers\FunctionCall;

use Zephir\Call;
use Zephir\CompilationContext;
use Zephir\CompiledExpression;
use Zephir\Optimizers\OptimizerAbstract;
use Zephir\Compiler\CompilerException;

class CalendsEndDateOptimizer extends OptimizerAbstract
{
    public function optimize(array $expression, Call $call, CompilationContext $context)
    {
        if (count($expression['parameters']) != 3) {
            throw new CompilerException("'Calends_end_date' only accepts 3 parameters", $expression);
        }

        $resolvedParams = $call->getReadOnlyResolvedParams($expression['parameters'], $context, $expression);

        /**
         * Process the expected symbol to be returned
         */
        $call->processExpectedReturn($context);
        $symbolVariable = $call->getSymbolVariable(true, $context);
        if ($symbolVariable->isNotVariableAndString()) {
            throw new CompilerException("Returned values can only be assigned to variant variables", $expression);
        }
        $context->headersManager->add('wrap_libcalends');
        $context->headersManager->add('kernel/string');
        $symbolVariable->setDynamicTypes('string');
        if ($call->mustInitSymbolVariable()) {
            $symbolVariable->initVariant($context);
        }

        $context->codePrinter->output("{$symbolVariable->getRealName()} = ext_Calends_end_date(" . implode(', ', $resolvedParams) . ');');
        return new CompiledExpression('variable', $symbolVariable->getRealName(), $expression);
    }
}
