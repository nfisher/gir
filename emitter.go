package gir

type GType int

type Token struct {
	Start int
	End   int
	Type  GType
	Data  []byte
}

type Emitter interface {
	Emit(start int, end int, gtype GType, d []byte)
}

const (
	Undefined GType = iota
	BooleanValue
	EnumValue
	FloatValue
	IntValue
	Name
	NullValue
	StringValue
	Variable
)
