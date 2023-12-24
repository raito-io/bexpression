package base

import (
	"context"
	"reflect"
	"testing"
)

func TestCtxExtendPathAndSetElement(t *testing.T) {
	type args struct {
		ctx     context.Context
		path    string
		element interface{}
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "first expression ctx element",
			args: args{
				ctx:     context.Background(),
				path:    "path",
				element: "element",
			},
			want: context.WithValue(context.WithValue(context.Background(), BExpressionPath, "root.path"), BExpressionElement, "element"),
		},
		{
			name: "second expression ctx element",
			args: args{
				ctx:     context.WithValue(context.WithValue(context.Background(), BExpressionPath, "root.path"), BExpressionElement, "element"),
				path:    "secondPath",
				element: "secondElement",
			},
			want: context.WithValue(context.WithValue(context.WithValue(context.WithValue(context.Background(), BExpressionPath, "root.path"), BExpressionElement, "element"),
				BExpressionPath, "root.path.secondPath"),
				BExpressionElement, "secondElement"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtxExtendPathAndSetElement(tt.args.ctx, tt.args.path, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CtxExtendPathAndSetElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
