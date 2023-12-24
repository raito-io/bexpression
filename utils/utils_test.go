package utils

import (
	"testing"
)

type testInterface interface {
	isTestInterface()
}

type testStruct struct{}

func (t *testStruct) isTestInterface() {}

func TestCountNonNil(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				values: []interface{}{},
			},
			want: 0,
		},
		{
			name: "all non-nil",
			args: args{
				values: []interface{}{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "all nil",
			args: args{
				values: []interface{}{nil, nil, nil},
			},
			want: 0,
		},
		{
			name: "some nil",
			args: args{
				values: []interface{}{nil, 1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNonNil(tt.args.values...); got != tt.want {
				t.Errorf("CountNonNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNull(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				value: nil,
			},
			want: true,
		},
		{
			name: "interface typed nil",
			args: args{
				value: (testInterface)(nil),
			},
			want: true,
		},
		{
			name: "typed nil",
			args: args{
				value: (*int)(nil),
			},
			want: true,
		},
		{
			name: "non-nil struct",
			args: args{
				value: testStruct{},
			},
			want: false,
		},
		{
			name: "non-nil interface",
			args: args{
				value: (testInterface)(&testStruct{}),
			},
			want: false,
		},
		{
			name: "non-nil pointer",
			args: args{
				value: Ptr(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNull(tt.args.value); got != tt.want {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}
