<?php

namespace Zephir\Optimizers\FunctionCall;

use Zephir\Call;
use Zephir\CompilationContext;
use Zephir\CompiledExpression;
use Zephir\Optimizers\OptimizerAbstract;
use Zephir\Compiler\CompilerException;

class CalendsContainsOptimizer extends OptimizerAbstract
{
    public function optimize(array $expression, Call $call, CompilationContext $context)
    {
        if (count($expression['parameters']) != 2) {
            throw new CompilerException("'Calends_contains' only accepts 2 parameters", $expression);
        }

        $resolvedParams = $call->getReadOnlyResolvedParams($expression['parameters'], $context, $expression);

        /**
         * Process the expected symbol to be returned
         */
        $call->processExpectedReturn($context);
        $symbolVariable = $call->getSymbolVariable(true, $context);
        if ($symbolVariable->isNotVariable()) {
            throw new CompilerException("Return value can only be assigned to variant variables", $expression);
        }
        $context->headersManager->add('wrap_libcalends');
        $context->headersManager->add('kernel/string');
        $symbolVariable->setDynamicTypes('string');
        if ($call->mustInitSymbolVariable()) {
            $symbolVariable->initVariant($context);
        }

        $context->codePrinter->output("{$symbolVariable->getRealName()} = ext_Calends_contains(" . implode(', ', $resolvedParams) . ');');
        return new CompiledExpression('variable', $symbolVariable->getRealName(), $expression);
    }
}
