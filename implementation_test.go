package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func ImplementationTest(t *testing.T) { TestingT(t) }

func (s *TestSuite) PostToInTest(c *C) {
// PostToIn means Postfix to Infix
	examples := map[string]string{
		"132 14.11 - 11 *":           "(132 - 14.11) * 11",
		"14112001 20 -":              "14112001 - 20",
		"4 3 2 -":                    "too many operands",
		"2 5 ^ 10 *":                 "2 ^ 5 * 10",
		"10 9 8 7 6 5 4 - / ^ - * +": "10 + 9 * (8 - 7 ^ (6 / (5 - 4)))",
		"Some text":                  "invalid input expression",
		"777 1901.2021 - 12 + - -":   "too many operators",
		"":                           "invalid input expression",
	}

	for post, expected := range examples {
		res, err := PostToIn(post)
		if err != nil {
			c.Assert(err, ErrorMatches, expected)
		} else {
			c.Assert(res, Equals, expected)
		}
	}
}

func ExamplePostToIn() {
	res, err := PostToIn("14 11 + 1 ^")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	// expected output:
	// (14 + 11) ^ 1
}
