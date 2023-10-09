package main

import "fmt"

const (
	PLUS string = "+"
	MINUS string = "-"
	MULTI  string = "*"
	DIV string = "/"
	REMAINDER string = "%"
	SQUARE string = "**"
)

const (
	PLUS_NO int = 1
        MINUS_NO int = 2
        MULTI_NO int = 3
        DIV_NO int = 4
        REMAINDER_NO int = 5
        SQUARE_NO int = 6
)

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

        if (operation == DIV_NO || operation == REMAINDER_NO) && rightOperand == 0 {
                result = fmt.Sprintf("%d %s %d는 연산할 수 없습니다.", leftOperand, opcode, rightOperand)
        } else if operation > 0 && operation < 7 {
                result = calc(leftOperand, rightOperand, opcode)
        }

        fmt.Println("결과: ", result)
}

func calc(left, right int64, opcode string) string {
	var result int64
	switch opcode {
	case PLUS:
		result = left + right
	case MINUS:
		result = left - right
	case MULTI:
		result = left * right
	case DIV:
		result = left / right
	case REMAINDER:
		result = left % right
	case SQUARE:
		result = pow(left, right)
	}
	return fmt.Sprintf("%d %s %d = %d", left, opcode, right, result)
}

func getOpcode(operation int) (opcode string) {
	switch operation {
	case PLUS_NO:
                opcode = PLUS
	case MINUS_NO:
                opcode = MINUS
	case MULTI_NO:
                opcode = MULTI
	case DIV_NO:
                opcode = DIV
	case REMAINDER_NO:
                opcode = REMAINDER
	case SQUARE_NO:
		opcode = SQUARE
	}
        return
}

func pow(base, cnt int64) int64 {
	var result int64 = 1
	for {
		result *= base
		cnt--
		if cnt == 0 {
			return result
		}
	}
}
