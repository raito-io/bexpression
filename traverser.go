package bexpression

import (
	"time"
)

type traverseOptions struct {
	EnterBinaryExpressionFn func(node *BinaryExpression) error
	LeaveBinaryExpressionFn func(node *BinaryExpression)

	LiteralBoolFn      func(value bool) error
	LiteralIntFn       func(value int) error
	LiteralFloatFn     func(value float64) error
	LiteralStringFn    func(value string) error
	LiteralTimestampFn func(value time.Time) error

	EnterComparisonFn    func(node *BinaryComparison) error
	ComparisonOperatorFn func(node ComparisonOperator) error
	LeaveComparisonFn    func(node *BinaryComparison)

	EnterOperandFn func(node *Operand) error
	LeaveOperandFn func(node *Operand)

	ReferenceFn func(node *Reference) error

	EnterAggregatorFn     func(node *Aggregator) error
	NextAggregatorOperand func() error
	LeaveAggregatorFn     func(node *Aggregator)

	EnterUnaryExpressionFn func(node *UnaryExpression) error
	LeaveUnaryExpressionFn func(node *UnaryExpression)
}

func WithEnterBinaryExpressionFn(fn func(node *BinaryExpression) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.EnterBinaryExpressionFn = fn
	}
}

func WithLeaveBinaryExpressionFn(fn func(node *BinaryExpression)) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LeaveBinaryExpressionFn = fn
	}
}

func WithLiteralBoolFn(fn func(value bool) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LiteralBoolFn = fn
	}
}

func WithLiteralIntFn(fn func(value int) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LiteralIntFn = fn
	}
}

func WithLiteralFloatFn(fn func(value float64) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LiteralFloatFn = fn
	}
}

func WithLiteralStringFn(fn func(value string) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LiteralStringFn = fn
	}
}

func WithLiteralTimestampFn(fn func(value time.Time) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LiteralTimestampFn = fn
	}
}

func WithEnterComparisonFn(fn func(node *BinaryComparison) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.EnterComparisonFn = fn
	}
}

func WithComparisonOperatorFn(fn func(node ComparisonOperator) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.ComparisonOperatorFn = fn
	}
}

func WithLeaveComparisonFn(fn func(node *BinaryComparison)) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LeaveComparisonFn = fn
	}
}

func WithEnterOperandFn(fn func(node *Operand) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.EnterOperandFn = fn
	}
}

func WithLeaveOperandFn(fn func(node *Operand)) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LeaveOperandFn = fn
	}
}

func WithReferenceFn(fn func(node *Reference) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.ReferenceFn = fn
	}
}

func WithEnterAggregatorFn(fn func(node *Aggregator) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.EnterAggregatorFn = fn
	}
}

func WithNextAggregatorOperand(fn func() error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.NextAggregatorOperand = fn
	}
}

func WithLeaveAggregatorFn(fn func(node *Aggregator)) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LeaveAggregatorFn = fn
	}
}

func WithEnterUnaryExpressionFn(fn func(node *UnaryExpression) error) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.EnterUnaryExpressionFn = fn
	}
}

func WithLeaveUnaryExpressionFn(fn func(node *UnaryExpression)) func(o *traverseOptions) {
	return func(o *traverseOptions) {
		o.LeaveUnaryExpressionFn = fn
	}
}

type Traverser struct {
	funcs traverseOptions
}

func NewTraverser(opts ...func(*traverseOptions)) *Traverser {
	options := traverseOptions{
		EnterBinaryExpressionFn: func(node *BinaryExpression) error { return nil },
		LeaveBinaryExpressionFn: func(node *BinaryExpression) {},
		LiteralBoolFn:           func(value bool) error { return nil },
		LiteralIntFn:            func(value int) error { return nil },
		LiteralFloatFn:          func(value float64) error { return nil },
		LiteralStringFn:         func(value string) error { return nil },
		LiteralTimestampFn:      func(value time.Time) error { return nil },
		EnterComparisonFn:       func(node *BinaryComparison) error { return nil },
		ComparisonOperatorFn:    func(node ComparisonOperator) error { return nil },
		LeaveComparisonFn:       func(node *BinaryComparison) {},
		EnterOperandFn:          func(node *Operand) error { return nil },
		LeaveOperandFn:          func(node *Operand) {},
		ReferenceFn:             func(node *Reference) error { return nil },
		EnterAggregatorFn:       func(node *Aggregator) error { return nil },
		NextAggregatorOperand:   func() error { return nil },
		LeaveAggregatorFn:       func(node *Aggregator) {},
		EnterUnaryExpressionFn:  func(node *UnaryExpression) error { return nil },
		LeaveUnaryExpressionFn:  func(node *UnaryExpression) {},
	}
	for _, opt := range opts {
		opt(&options)
	}

	return &Traverser{
		funcs: options,
	}
}

func (t *Traverser) TraverseBinaryExpression(node *BinaryExpression) error {
	if err := t.funcs.EnterBinaryExpressionFn(node); err != nil {
		return err
	}

	defer t.funcs.LeaveBinaryExpressionFn(node)

	if node.Literal != nil {
		return t.TraverseLiteralBool(*node.Literal)
	} else if node.Comparison != nil {
		return t.TraverseComparison(node.Comparison)
	} else if node.Aggregator != nil {
		return t.TraverseAggregator(node.Aggregator)
	} else if node.UnaryExpression != nil {
		return t.TraverseUnaryExpression(node.UnaryExpression)
	}

	return nil
}

func (t *Traverser) TraverseLiteralBool(value bool) error {
	return t.funcs.LiteralBoolFn(value)
}

func (t *Traverser) TraverseComparison(node *BinaryComparison) error {
	if err := t.funcs.EnterComparisonFn(node); err != nil {
		return err
	}

	defer t.funcs.LeaveComparisonFn(node)

	if err := t.TraverseOperand(&node.LeftOperand); err != nil {
		return err
	}

	if err := t.funcs.ComparisonOperatorFn(node.Operator); err != nil {
		return err
	}

	if err := t.TraverseOperand(&node.RightOperand); err != nil {
		return err
	}

	return nil
}

func (t *Traverser) TraverseOperand(node *Operand) error {
	if err := t.funcs.EnterOperandFn(node); err != nil {
		return err
	}

	defer t.funcs.LeaveOperandFn(node)

	if node.Reference != nil {
		return t.TraverseReference(node.Reference)
	} else if node.Literal != nil {
		return t.TraverseLiteral(node.Literal)
	}

	return nil
}

func (t *Traverser) TraverseReference(reference *Reference) error {
	return t.funcs.ReferenceFn(reference)
}

func (t *Traverser) TraverseLiteral(literal *Literal) error {
	if literal.Bool != nil {
		return t.funcs.LiteralBoolFn(*literal.Bool)
	} else if literal.Int != nil {
		return t.funcs.LiteralIntFn(*literal.Int)
	} else if literal.Float != nil {
		return t.funcs.LiteralFloatFn(*literal.Float)
	} else if literal.Timestamp != nil {
		return t.funcs.LiteralTimestampFn(*literal.Timestamp)
	} else if literal.Str != nil {
		return t.funcs.LiteralStringFn(*literal.Str)
	}

	return nil
}

func (t *Traverser) TraverseAggregator(aggregator *Aggregator) error {
	if err := t.funcs.EnterAggregatorFn(aggregator); err != nil {
		return err
	}

	defer t.funcs.LeaveAggregatorFn(aggregator)

	for i := range aggregator.Operands {
		if i > 0 {
			if err := t.funcs.NextAggregatorOperand(); err != nil {
				return err
			}
		}

		if err := t.TraverseBinaryExpression(&aggregator.Operands[i]); err != nil {
			return err
		}
	}

	return nil
}

func (t *Traverser) TraverseUnaryExpression(unaryExpression *UnaryExpression) error {
	if err := t.funcs.EnterUnaryExpressionFn(unaryExpression); err != nil {
		return err
	}

	defer t.funcs.LeaveUnaryExpressionFn(unaryExpression)

	return t.TraverseBinaryExpression(&unaryExpression.Operand)
}
