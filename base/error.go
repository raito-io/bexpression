package base

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyOperands = errors.New("at least one operand must be set")

type ErrUnionExpression struct {
	ElementType       string
	ElementChildTypes []string
}

func NewUnionExpressionError(elementType string, elementChildTypes []string) *ErrUnionExpression {
	return &ErrUnionExpression{
		ElementType:       elementType,
		ElementChildTypes: elementChildTypes,
	}
}

func (e *ErrUnionExpression) Error() string {
	return fmt.Sprintf("%s expect only one of %s", e.ElementType, strings.Join(e.ElementChildTypes, ", "))
}
