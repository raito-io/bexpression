package bexpression

import (
	"testing"

	"github.com/aws/smithy-go/ptr"
)

type TestStruct struct {
}

func Test_countNonNil(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty list",
			args: args{
				values: []interface{}{},
			},
			want: 0,
		},
		{
			name: "one non nil element",
			args: args{
				values: []interface{}{nil, 1, nil, nil},
			},
			want: 1,
		},
		{
			name: "five non nil element",
			args: args{
				values: []interface{}{1, (*TestStruct)(nil), nil, 3, 4, nil, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNonNil(tt.args.values...); got != tt.want {
				t.Errorf("countNonNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryExpressionDebugString(t *testing.T) {
	type args struct {
		b BinaryExpression
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "literal binary expression",
			args: args{
				b: BinaryExpression{
					Literal: ptr.Bool(true),
				},
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "expression with comparison",
			args: args{
				b: BinaryExpression{
					Comparison: &BinaryComparison{
						LeftOperand: Operand{
							Literal: &Literal{
								Int: ptr.Int(1),
							},
						},
						Operator: ComparisonOperatorGreaterThan,
						RightOperand: Operand{
							Reference: &Reference{
								EntityId:   "EntityId",
								EntityType: EntityTypeDataObject,
							},
						},
					},
				},
			},
			want:    "1 GreaterThan DataObject:EntityId",
			wantErr: false,
		},
		{
			name: "expression with aggregator",
			args: args{
				b: BinaryExpression{
					Aggregator: &Aggregator{
						Operator: AggregatorOperatorAnd,
						Operands: []BinaryExpression{
							{
								Literal: ptr.Bool(true),
							},
							{
								UnaryExpression: &UnaryExpression{

									Operator: UnaryOperatorNot,
									Operand: BinaryExpression{
										Comparison: &BinaryComparison{LeftOperand: Operand{Reference: &Reference{EntityId: "Id1", EntityType: EntityTypeDataObject}}, Operator: ComparisonOperatorNotEqual, RightOperand: Operand{Literal: &Literal{String: ptr.String("NJ")}}},
									},
								},
							},
							{
								Comparison: &BinaryComparison{LeftOperand: Operand{Reference: &Reference{EntityId: "Id2", EntityType: EntityTypeDataObject}}, Operator: ComparisonOperatorNotEqual, RightOperand: Operand{Literal: &Literal{Int: ptr.Int(100)}}},
							},
						},
					},
				},
			},
			want:    "(true) And (Not DataObject:Id1 NotEqual \"NJ\") And (DataObject:Id2 NotEqual 100)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryExpressionDebugString(&tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryExpressionDebugString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BinaryExpressionDebugString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
