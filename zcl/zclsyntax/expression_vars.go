package zclsyntax

// Generated by expression_vars_get.go. DO NOT EDIT.
// Run 'go generate' on this package to update the set of functions here.

import (
	"github.com/zclconf/go-zcl/zcl"
)

func (e *BinaryOpExpr) Variables() []zcl.Traversal {
	return Variables(e)
}

func (e *ConditionalExpr) Variables() []zcl.Traversal {
	return Variables(e)
}

func (e *FunctionCallExpr) Variables() []zcl.Traversal {
	return Variables(e)
}

func (e *LiteralValueExpr) Variables() []zcl.Traversal {
	return Variables(e)
}

func (e *ScopeTraversalExpr) Variables() []zcl.Traversal {
	return Variables(e)
}

func (e *UnaryOpExpr) Variables() []zcl.Traversal {
	return Variables(e)
}
