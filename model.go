package bexpression

import (
	"github.com/raito-io/bexpression/base"
	"github.com/raito-io/bexpression/data_comparison"
)

type DataComparisonExpression = base.BinaryExpression[*data_comparison.DataComparison]
