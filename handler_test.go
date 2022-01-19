package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

func CompTest(t *testing.T) { TestingT(t) }

func (s *TestSuite) CompOutTest(c *C) {
	inputStr, expected := "9 5 3 ^ -", "9 - 5 ^ 3\n"
	buf := new(bytes.Buffer)
	reader := strings.NewReader(inputStr)
	handler := ComputeHandler{Input: reader, Output: buf}
	handler.Compute()
	c.Assert(buf.String(), Equals, expected)
}

func (s *TestSuite) CompSyntaxTest(c *C) {
	errorExamples := map[string]ComputeHandler{
		"input error":   ComputeHandler{},
		"output error":  ComputeHandler{Input: strings.NewReader("2 4 +")},
		"invalid input": ComputeHandler{Input: strings.NewReader(""), Output: new(bytes.Buffer)},
	}
	for expected, obtained := range errorExamples {
		errObtained := obtained.Compute()
		c.Assert(errObtained, ErrorMatches, expected)
	}
}
