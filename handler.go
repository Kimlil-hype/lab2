package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// ComputeHandler prefix to infix interpreter
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Prefix to infix wrapper
func (ch *ComputeHandler) Compute() error {
	if ch.Input == nil {
		return fmt.Errorf("input error")
	}
	if ch.Output == nil {
		return fmt.Errorf("output error")
	}
	buf, readErr := ioutil.ReadAll(ch.Input)
	if readErr != nil {
		return readErr
	}
	strInput := strings.Trim(string(buf), "\n")
	computed, computeErr := PostToIn(strInput)
	if computeErr != nil {
		return computeErr
	}
	res := []byte(computed + "\n")
	_, writeErr := ch.Output.Write(res)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
