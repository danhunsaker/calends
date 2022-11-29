<?php

namespace Zephir\Optimizers\FunctionCall;

use Zephir\Call;
use Zephir\CompilationContext;
use Zephir\CompiledExpression;
use Zephir\Optimizers\OptimizerAbstract;
use Zephir\Compiler\CompilerException;

class CalendsReleaseOptimizer extends OptimizerAbstract
{
    public function optimize(array $expression, Call $call, CompilationContext $context)
    {
        if (count($expression['parameters']) != 1) {
            throw new CompilerException("'Calends_release' only accepts 1 parameters", $expression);
        }

        $resolvedParams = $call->getReadOnlyResolvedParams($expression['parameters'], $context, $expression);

        $context->headersManager->add('wrap_libcalends');
        $resolvedParams = $call->getReadOnlyResolvedParams($expression['parameters'], $context, $expression);

        $context->codePrinter->output('ext_Calends_release(' . implode(', ', $resolvedParams) . ');');
        return new CompiledExpression('void', 'ext_Calends_release(' . implode(', ', $resolvedParams) . ')', $expression);
    }
}
