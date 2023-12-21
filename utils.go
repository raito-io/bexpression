package bexpression

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func countNonNil(values ...interface{}) int {
	r := 0

	for i := range values {
		value := values[i]
		if value != nil {
			v := reflect.ValueOf(value)

			if v.Kind() != reflect.Ptr || !v.IsNil() {
				r++
			}
		}
	}

	return r
}

func BinaryExpressionDebugString(b *BinaryExpression) (string, error) {
	var sb strings.Builder

	var aggregationOperationStack []AggregatorOperator

	traverser := NewTraverser(
		WithLiteralBoolFn(func(value bool) error {
			_, err := sb.WriteString(strconv.FormatBool(value))

			return err
		}),
		WithLiteralIntFn(func(value int) error {
			_, err := sb.WriteString(strconv.Itoa(value))

			return err
		}),
		WithLiteralFloatFn(func(value float64) error {
			_, err := sb.WriteString(strconv.FormatFloat(value, 'f', -1, 64))

			return err
		}),
		WithLiteralStringFn(func(value string) error {
			_, err := sb.WriteString(strconv.Quote(value))

			return err
		}),
		WithLiteralStringListFn(func(value []string) error {
			_, err := sb.WriteString(fmt.Sprintf("%+v", value))

			return err
		}),
		WithLiteralTimestampFn(func(value time.Time) error {
			_, err := sb.WriteString(value.String())

			return err
		}),
		WithComparisonOperatorFn(func(node ComparisonOperator) error {
			_, err := sb.WriteString(" " + node.String() + " ")

			return err
		}),
		WithReferenceFn(func(node *Reference) error {
			_, err := sb.WriteString(fmt.Sprintf("%s:%s", node.EntityType.String(), node.EntityId))

			return err
		}),
		WithEnterAggregatorFn(func(node *Aggregator) error {
			_, err := sb.WriteString("(")
			if err != nil {
				return err
			}

			aggregationOperationStack = append(aggregationOperationStack, node.Operator)

			return nil
		}),
		WithNextAggregatorOperand(func() error {
			_, err := sb.WriteString(") " + aggregationOperationStack[len(aggregationOperationStack)-1].String() + " (")

			return err
		}),
		WithLeaveAggregatorFn(func(node *Aggregator) {
			sb.WriteString(")")

			aggregationOperationStack = aggregationOperationStack[:len(aggregationOperationStack)-1]
		}),
		WithEnterUnaryExpressionFn(func(node *UnaryExpression) error {
			_, err := sb.WriteString(node.Operator.String() + " ")

			return err
		}),
	)

	err := traverser.TraverseBinaryExpression(b)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
