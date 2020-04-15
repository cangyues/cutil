package util

import (
	"bytes"
	"strings"
)

type L struct {
	str bytes.Buffer
}

func NewL() *L {
	return new(L)
}

func (l *L) A(v string) *L {
	l.str.WriteString(v)
	return l
}

func (l *L) S(v string) *L {
	l.str.WriteString("'")
	l.str.WriteString(strings.Replace(v, "'", "", -1))
	l.str.WriteString("'")
	return l
}

func (l *L) S2(v string) *L {
	l.str.WriteString("'")
	l.str.WriteString(strings.Replace(v, "'", "", -1))
	l.str.WriteString("',")
	return l
}

func (l *L) ToStr() string {
	return l.str.String()
}
