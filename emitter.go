package gir

// GType is an enum type for classes of tokens.
type GType int

// Token contains the position and type of emitted lexical token.
type Token struct {
	Start int
	End   int
	Type  GType
	Data  []byte
}

// Emitter provides a common interface for emitting lexical tokens.
type Emitter interface {
	Emit(start int, end int, gtype GType, d []byte)
}

const (
	// Undefined is a token that is undefined based on lexical analysis.
	Undefined GType = iota
	// BooleanValue is the enum for a boolean token.
	BooleanValue
	// EnumValue is the enum for a enum token.
	EnumValue
	// FloatValue is the enum for a float token.
	FloatValue
	// IntValue is the enum for a int token.
	IntValue
	// Name is the enum for a name token.
	Name
	// NullValue is the enum for a null token.
	NullValue
	// StringValue is the enum for a string token.
	StringValue
	// Variable is the enum for a variable token.
	Variable
)
