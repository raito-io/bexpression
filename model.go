package bexpression

import (
	"errors"
	"time"
)

type OperandUnion interface {
	isOperand()
}

type BinaryExpressionUnion interface {
	OperandUnion
	isBinaryExpression()
}

type ReferenceItem interface{}

type BinaryExpression struct {
	Literal         *bool             `json:"literal,omitempty" yaml:"literal,omitempty" gqlgen:"literal"`
	Comparison      *BinaryComparison `json:"comparison,omitempty" yaml:"comparison,omitempty" gqlgen:"comparison"`
	Aggregator      *Aggregator       `json:"aggregator,omitempty" yaml:"aggregator,omitempty" gqlgen:"aggregator"`
	UnaryExpression *UnaryExpression  `json:"unaryExpression,omitempty" yaml:"unaryExpression,omitempty" gqlgen:"unaryExpression"`
}

func (b *BinaryExpression) Validate() error {
	if countNonNil(b.Literal, b.Comparison, b.Aggregator, b.UnaryExpression) != 1 {
		return errors.New("exactly one of literal, comparison, aggregator or not must be set")
	}

	if b.Literal != nil {
		return nil
	} else if b.Comparison != nil {
		return b.Comparison.Validate()
	} else if b.Aggregator != nil {
		return b.Aggregator.Validate()
	} else if b.UnaryExpression != nil {
		return b.UnaryExpression.Validate()
	}

	return errors.New("no valid expression element set")
}

func (b *BinaryExpression) ToGql() (BinaryExpressionUnion, error) {
	if b.Literal != nil {
		return &LiteralBool{Value: *b.Literal}, nil
	} else if b.Comparison != nil {
		return b.Comparison, nil
	} else if b.Aggregator != nil {
		return b.Aggregator, nil
	} else if b.UnaryExpression != nil {
		return b.UnaryExpression, nil
	}

	return nil, errors.New("no valid expression element set")
}

type LiteralUnion interface {
	OperandUnion
	isLiteral()
}

type Literal struct {
	Bool       *bool      `json:"bool,omitempty" yaml:"bool,omitempty" gqlgen:"bool"`
	Int        *int       `json:"int,omitempty" yaml:"int,omitempty" gqlgen:"int"`
	Float      *float64   `json:"float,omitempty" yaml:"float,omitempty" gqlgen:"float"`
	String     *string    `json:"string,omitempty" yaml:"string,omitempty" gqlgen:"string"`
	Timestamp  *time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty" gqlgen:"timestamp"`
	StringList []string   `json:"stringList,omitempty" yaml:"stringList,omitempty" gqlgen:"stringList"`
}

func (l *Literal) Validate() error {
	if countNonNil(l.Bool, l.Int, l.Float, l.String, l.Timestamp, l.StringList) != 1 {
		return errors.New("exactly one of literal value must be set")
	}

	return nil
}

func (l *Literal) ToGql() (LiteralUnion, error) {
	if l.Bool != nil {
		return &LiteralBool{Value: *l.Bool}, nil
	} else if l.Int != nil {
		return &LiteralInt{Value: *l.Int}, nil
	} else if l.Float != nil {
		return &LiteralFloat{Value: *l.Float}, nil
	} else if l.String != nil {
		return &LiteralString{Value: *l.String}, nil
	} else if l.StringList != nil {
		return &LiteralStringList{Value: l.StringList}, nil
	} else if l.Timestamp != nil {
		return &LiteralTimestamp{Value: *l.Timestamp}, nil
	}

	return nil, errors.New("no valid literal element set")
}

type LiteralBool struct {
	Value bool
}

func (l *LiteralBool) isBinaryExpression() {}
func (l *LiteralBool) isOperand()          {}
func (l *LiteralBool) isLiteral()          {}

type LiteralInt struct {
	Value int `json:"value" yaml:"value" gqlgen:"value"`
}

func (l *LiteralInt) isLiteral() {}
func (l *LiteralInt) isOperand() {}

type LiteralFloat struct {
	Value float64
}

func (l *LiteralFloat) isLiteral() {}
func (l *LiteralFloat) isOperand() {}

type LiteralString struct {
	Value string
}

func (l *LiteralString) isLiteral() {}
func (l *LiteralString) isOperand() {}

type LiteralStringList struct {
	Value []string
}

func (l *LiteralStringList) isLiteral() {}
func (l *LiteralStringList) isOperand() {}

type LiteralTimestamp struct {
	Value time.Time
}

func (l *LiteralTimestamp) isLiteral() {}
func (l *LiteralTimestamp) isOperand() {}

type BinaryComparison struct {
	Operator     ComparisonOperator `json:"operator" yaml:"operator" gqlgen:"operator"`
	LeftOperand  Operand            `json:"leftOperand,omitempty" yaml:"leftOperand,omitempty" gqlgen:"leftOperand"`
	RightOperand Operand            `json:"rightOperand,omitempty" yaml:"rightOperand,omitempty" gqlgen:"rightOperand"`
}

func (b *BinaryComparison) Validate() error {
	if err := b.LeftOperand.Validate(); err != nil {
		return err
	}

	if err := b.RightOperand.Validate(); err != nil {
		return err
	}

	return nil
}

func (b *BinaryComparison) isBinaryExpression() {}
func (b *BinaryComparison) isOperand()          {}

type Operand struct {
	Reference *Reference `json:"reference,omitempty" yaml:"reference,omitempty" gqlgen:"reference"`
	Literal   *Literal   `json:"literal,omitempty" yaml:"literal,omitempty" gqlgen:"literal"`
}

func (o *Operand) ToGql() (OperandUnion, error) {
	if o.Literal != nil {
		return o.Literal.ToGql()
	} else if o.Reference != nil {
		return o.Reference, nil
	}

	return nil, errors.New("no valid expression element set")
}

func (o *Operand) Validate() error {
	if countNonNil(o.Literal, o.Reference) != 1 {
		return errors.New("exactly one of binaryExpression or reference must be set")
	}

	if o.Literal != nil {
		return o.Literal.Validate()
	} else if o.Reference != nil {
		return o.Reference.Validate()
	}

	return errors.New("no valid operand set")
}

type Reference struct {
	EntityType EntityType `json:"entityType,omitempty" yaml:"entityType,omitempty" gqlgen:"entityType"`
	EntityId   string     `json:"entityId,omitempty" yaml:"entityId,omitempty" gqlgen:"entityId"`
}

func (r *Reference) Validate() error {
	return nil
}

func (r *Reference) isOperand() {}

type Aggregator struct {
	Operator AggregatorOperator `json:"operator" yaml:"operator" gqlgen:"operator"`
	Operands []BinaryExpression `json:"operands" yaml:"operands" gqlgen:"operands"`
}

func (a *Aggregator) Validate() error {
	if len(a.Operands) == 0 {
		return errors.New("at least one operand must be set")
	}

	for _, operand := range a.Operands {
		if err := operand.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Aggregator) isBinaryExpression() {}
func (a *Aggregator) isOperand()          {}

type UnaryExpression struct {
	Operator UnaryOperator    `json:"operator" yaml:"operator" gqlgen:"operator"`
	Operand  BinaryExpression `json:"expression" yaml:"expression" gqlgen:"expression"`
}

func (n *UnaryExpression) Validate() error {
	return n.Operand.Validate()
}

func (n *UnaryExpression) isBinaryExpression() {}
func (n *UnaryExpression) isOperand()          {}

//go:generate go run github.com/raito-io/enumer -type=EntityType -values -gqlgen -yaml -json -trimprefix=EntityType
type EntityType int

const (
	EntityTypeDataObject EntityType = iota
)

//go:generate go run github.com/raito-io/enumer -type=ComparisonOperator -values -gqlgen -yaml -json -trimprefix=ComparisonOperator
type ComparisonOperator int

const (
	ComparisonOperatorEqual ComparisonOperator = iota
	ComparisonOperatorNotEqual
	ComparisonOperatorLessThan
	ComparisonOperatorLessThanOrEqual
	ComparisonOperatorGreaterThan
	ComparisonOperatorGreaterThanOrEqual
	ComparisonOperatorHas
	ComparisonOperatorContains
)

//go:generate go run github.com/raito-io/enumer -type=AggregatorOperator -values -gqlgen -yaml -json -trimprefix=AggregatorOperator
type AggregatorOperator int

const (
	AggregatorOperatorAnd AggregatorOperator = iota
	AggregatorOperatorOr
)

//go:generate go run github.com/raito-io/enumer -type=UnaryOperator -values -gqlgen -yaml -json -trimprefix=UnaryOperator
type UnaryOperator int

const (
	UnaryOperatorNot UnaryOperator = iota
)
