package main

import "fmt"

func main() {
	var (
		leftOperand int64
		rightOperand int64
		operation int
	)

	fmt.Println("피연산자 두 개를 입력해 주세요.")
	fmt.Printf("왼쪽 피연산자: ")
	fmt.Scanf("%d", &leftOperand)
	fmt.Printf("오른쪽 피연산자: ")
	fmt.Scanf("%d", &rightOperand)

	fmt.Println("연산자를 선택해주세요.")
	fmt.Println("1. +, 2. -, 3. *, 4. /, 5. %, 6. **")
	fmt.Scanf("%d", &operation)

	result := ""
	if (operation == 4 || operation == 5) && rightOperand == 0 {
		result = fmt.Sprintf("%d %s %d는 연산할 수 없습니다.", leftOperand, getOperationStr(operation), rightOperand)
	} else {
		result = calc(leftOperand, rightOperand, operation)
	}

	fmt.Println("결과: ", result)
}

func calc(left, right int64, operation int) string {
	var result int64
	if operation == 1 {
		result = left + right
        } else if operation == 2 {
		result = left - right
        } else if operation == 3 {
		result = left * right
        } else if operation == 4 {
		result = left / right
        } else if operation == 5 {
		result = left % right
        }

	operationStr := getOperationStr(operation)

	return fmt.Sprintf("%d %s %d = %d", left, operationStr, right, result)
}

func getOperationStr(operation int) (op string) {
        if operation == 1 {
                op = "+"
        } else if operation == 2 {
                op = "-"
        } else if operation == 3 {
                op = "*"
        } else if operation == 4 {
                op = "/"
        } else if operation == 5 {
                op = "%"
        }

        return
}
