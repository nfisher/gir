package gir_test

import (
	"testing"

	"github.com/nfisher/gir"
)

func Test_parse_queries(t *testing.T) {
	for _, td := range testQueries {
		var em testEmitter
		gir.ParseQuery([]byte(td.input), &em)

		if !equalTo(em.tokens, td.expected) {
			t.Errorf("%s \n| tokens = %#v\n|     want %#v", td.msg, em.tokens, td.expected)
		}
	}
}

var testQueries = []struct {
	msg      string
	input    string
	expected []gir.Token
}{
	{"simple query",
		`{
  allFilms {
    films {
      title
    }
  }
}`, []gir.Token{}},
}
