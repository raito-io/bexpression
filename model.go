package bexpression

import (
	"github.com/raito-io/bexpression/base"
	"github.com/raito-io/bexpression/datacomparison"
)

type DataComparisonExpression = base.BinaryExpression[*datacomparison.DataComparison]
