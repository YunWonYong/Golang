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

        opcode := getOpcode(operation)

        result := fmt.Sprintf("%d번은 잘 못된 입력입니다.", operation)

        if operation == 6 {
                result = fmt.Sprintf("%d(%s)번 연산자는 지원하고 있지 않습니다.", operation, opcode)
        } else if (operation == 4 || operation == 5) && rightOperand == 0 {
                result = fmt.Sprintf("%d %s %d는 연산할 수 없습니다.", leftOperand, opcode, rightOperand)
        } else if operation > 0 && operation < 7 {
                result = calc(leftOperand, rightOperand, opcode)
        }

        fmt.Println("결과: ", result)
}

func calc(left, right int64, opcode string) string {
	var result int64
	switch opcode {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	case "%":
		result = left % right
	case "**":

	}
	return fmt.Sprintf("%d %s %d = %d", left, opcode, right, result)
}

func getOpcode(operation int) (opcode string) {
	switch operation {
	case 1:
                opcode = "+"
	case 2:
                opcode = "-"
	case 3:
                opcode = "*"
	case 4:
                opcode = "/"
	case 5:
                opcode = "%"
	case 6:
		opcode = "**"
	}
        return
}
