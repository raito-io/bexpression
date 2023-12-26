package bexpression

import (
	"github.com/raito-io/bexpression/base"
	"github.com/raito-io/bexpression/datacomparison"
)

type DataComparisonExpression = base.BinaryExpression[*datacomparison.DataComparison]
type DataComparisonAggregator = base.Aggregator[*datacomparison.DataComparison]
type DataComparisonUnaryExpression = base.UnaryExpression[*datacomparison.DataComparison]
