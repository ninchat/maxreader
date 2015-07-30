package maxreader_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"."
)

func Test(t *testing.T) {
	type entry struct {
		s  string
		ok bool
	}

	table := []entry{
		{"", true},
		{"h", true},
		{"hell", true},
		{"hello", true},
		{"hellow", false},
		{"helloworld", false},
	}

	for _, e := range table {
		b, err := ioutil.ReadAll(maxreader.New(strings.NewReader(e.s), 5))

		if e.ok != (err == nil) {
			t.Errorf(`input "%v" -> error %v`, e.s, err)
		}

		if err == nil {
			l := len(e.s)
			if l > 5 {
				l = 5
			}

			if len(b) != l {
				t.Errorf(`input "%v" -> length %v`, e.s, len(b))
			}
		} else if err != maxreader.ErrReadLimit {
			t.Errorf(`input "%v" -> wrong error %v`, e.s, err)
		}
	}
}
