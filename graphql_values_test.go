package gir_test

import (
	"testing"

	"github.com/nfisher/gir"
)

func Test_parse_values(t *testing.T) {
	for _, td := range testValues {
		var em testEmitter
		gir.ParseValues([]byte(td.input), &em)

		if !equalTo(em.tokens, td.expected) {
			t.Errorf("%s \n| tokens = %#v\n|     want %#v", td.msg, em.tokens, td.expected)
		}
	}
}

var testValues = []struct {
	msg      string
	input    string
	expected []gir.Token
}{
	{"positive integer",
		"  123  ", []gir.Token{{Start: 2, End: 5, Type: gir.IntValue}},
	},
	{"negative integer",
		"  -123  ", []gir.Token{{Start: 2, End: 6, Type: gir.IntValue}},
	},
	{"integer with comment",
		"  -123  # 345", []gir.Token{{Start: 2, End: 6, Type: gir.IntValue}},
	},
	{"positive float",
		"  1.0  ", []gir.Token{{Start: 2, End: 5, Type: gir.FloatValue}},
	},
	{"negative float",
		"  -1.0  ", []gir.Token{{Start: 2, End: 6, Type: gir.FloatValue}},
	},
	{"float positive exponent",
		"  1.0e6  ", []gir.Token{{Start: 2, End: 7, Type: gir.FloatValue}},
	},
	{"float negative exponent",
		"  1.0e-6  ", []gir.Token{{Start: 2, End: 8, Type: gir.FloatValue}},
	},
	{"true",
		"  true  ", []gir.Token{{Start: 2, End: 6, Type: gir.BooleanValue}},
	},
	{"false",
		"  false  ", []gir.Token{{Start: 2, End: 7, Type: gir.BooleanValue}},
	},
	{"null",
		"  null  ", []gir.Token{{Start: 2, End: 6, Type: gir.NullValue}},
	},
	{"enum",
		"  _hola  ", []gir.Token{{Start: 2, End: 7, Type: gir.EnumValue}},
	},
	{"var",
		"  $hola  ", []gir.Token{{Start: 2, End: 7, Type: gir.Variable}},
	},
	{"string with escape",
		`  "\"Hello\t\r\nWorld\""  `, []gir.Token{{Start: 2, End: 24, Type: gir.StringValue}},
	},
	{"string without termination",
		`  "\"Hello\t\r\nWorld\"  `, []gir.Token{},
	},
}

func equalTo(a, b []gir.Token) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v.Start != b[i].Start {
			return false
		}

		if v.End != b[i].End {
			return false
		}

		if v.Type != b[i].Type {
			return false
		}
	}

	return true
}

func (t *testEmitter) Emit(start, end int, gtype gir.GType, d []byte) {
	t.tokens = append(t.tokens, gir.Token{Start: start, End: end, Type: gtype, Data: d})
}

type testEmitter struct {
	tokens []gir.Token
}
