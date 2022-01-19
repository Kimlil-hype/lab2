package lab2

import (
	"fmt"
	"regexp"
	"strings"
)

type validator struct {
	Operator string
	Operand  string
}

func (v *validator) ValidSrting(input string) bool {
	validator := fmt.Sprintf(`^((%s|%s)\s){2,}(%s\s){0,}%s$`, v.Operand, v.Operator, v.Operator, v.Operator)
	isValid, _ := regexp.MatchString(validator, input)
	return isValid
}

func (v *validator) CheckArgsAmount(args []string) error {
	operators, operands := 0, 0
	macthOperator := fmt.Sprintf(`^%s$`, v.Operator)
	macthOperand := fmt.Sprintf(`^%s$`, v.Operand)
	for _, arg := range args {
		if isOperator, _ := regexp.MatchString(macthOperator, arg); isOperator {
			operators++
		} else if isOperand, _ := regexp.MatchString(macthOperand, arg); isOperand {
			operands++
		}
	}
	if operators+operands != len(args) {
		return fmt.Errorf("expression argument error")
	} else if operands > operators+1 {
		return fmt.Errorf("Error -- too many operands")
	} else if operators > operands-1 {
		return fmt.Errorf("Error -- too many operators")
	} else {
		return nil
	}
}

func (v *validator) IncludesOperator(str string) bool {
	includes, _ := regexp.MatchString(v.Operator, str)
	return includes
}

func (v *validator) CheckOperator(str string) bool {
	macthOperator := fmt.Sprintf(`^%s$`, v.Operator)
	includes, _ := regexp.MatchString(macthOperator, str)
	return includes
}

func PostToIn(postStr string) (inStr string, err error) {
	v := validator{Operator: `[-\+\*\^\/]`, Operand: `(\d+|(\d+[,\.]\d+))`}
	if !v.ValidSrting(postStr) {
		err = fmt.Errorf("input error")
		return
	}
	var operatorsStack []string
	var inHeap []string
	prop := map[string]uint8{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
	postArgs := strings.Split(postStr, " ")
	if agrsErr := v.CheckArgsAmount(postArgs); agrsErr != nil {
		err = agrsErr
		return
	}

	for _, operator := range postArgs {
		if !v.CheckOperator(operator) {
			inHeap = append(inHeap, operator)
			continue
		}
		operatorsStack = append(operatorsStack, operator)
		sliced := inHeap[(len(inHeap) - 2):]
		inHeap = inHeap[:(len(inHeap) - 2)]

		if len(operatorsStack) > 1 {
			prevOperator := operatorsStack[(len(operatorsStack) - 2)]

			if prop[operator] > prop[prevOperator] || prop[operator] == 3 && prop[prevOperator] == 3 {
				if v.IncludesOperator(sliced[1]) {
					sliced[1] = "(" + sliced[1] + ")"
				} else {
					sliced[0] = "(" + sliced[0] + ")"
				}
			}
		}

		operand := fmt.Sprintf("%s %s %s", sliced[0], operator, sliced[1])
		inHeap = append(inHeap, operand)
	}

	inStr = inHeap[0]
	return
}
