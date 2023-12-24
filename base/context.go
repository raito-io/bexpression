package base

import (
	"context"
	"fmt"
)

type BExpressionCtxKey string

const (
	BExpressionPath    = BExpressionCtxKey("BExpressionPath")
	BExpressionElement = BExpressionCtxKey("BExpressionElement")
)

func CtxExtendPathAndSetElement(ctx context.Context, path string, element interface{}) context.Context {
	currentPath, found := ctx.Value(BExpressionPath).(string)
	if !found {
		currentPath = "root."
	}

	ctx = context.WithValue(ctx, BExpressionElement, fmt.Sprintf("%s.%s", currentPath, path))
	ctx = context.WithValue(ctx, BExpressionElement, element)

	return ctx
}
