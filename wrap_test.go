package msvg

import (
	"bytes"
	"strings"
	"testing"

	svg "github.com/ajstarks/svgo"
)

func Test_Wrap(t *testing.T) {
	ts := []struct {
		s1 string
		bp int
		s2 string
	}{
		{"hello world", 6, "hello|world"},
		{"hello world", 4, "hello|world"},
		{"hello\nwor ld", 4, "hello|wor|ld"},
		{"hello world", 0, "hello|world"},
		{"hello-world", 3, "hello-|world"},
		{"choose a Gerrymander from the pile and Gerrymander a Hex", 20,
			"choose a Gerrymander|from the pile and|Gerrymander a Hex"},
		{"Choose any Gerrymander from the pile and Gerrymander a Hex", 20,
			"Choose any|Gerrymander from the|pile and Gerrymander|a Hex"},
	}
	for _, v := range ts {
		ss := Wrap(v.s1, v.bp)
		rs := strings.Join(ss, "|")
		if rs != v.s2 {
			t.Logf("With :%d:%s ,\nExp: %s\nGot: %s", v.bp, v.s1, v.s2, rs)
			t.Fail()
		}
	}

}

func Test_Lines(t *testing.T) {
	var b bytes.Buffer
	g := svg.New(&b)
	g.Start(500, 500)
	g.Textlines(10, 10, Wrap("Hello world poo", 4), 10, 12, "red", "middle")
	g.End()
	t.Log(string(b.Bytes()))

}
