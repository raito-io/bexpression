package base

type Literal[T any] struct {
	Value T `json:"value" yaml:"value" gqlgen:"value"`
}

type LiteralBool Literal[bool]

func (l *LiteralBool) IsBinaryExpression() {}
