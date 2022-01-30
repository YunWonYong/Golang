package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	leftOperand, leftBinaryCode := Int8ToBinaryCode(15)
	rightOperand, rightBinaryCode := Int8ToBinaryCode(81)
	fmt.Printf("leftOperand(%T):  %d, leftBinaryCode:  %s\n", leftOperand, leftOperand, leftBinaryCode)
	fmt.Printf("rightOperand(%T): %d, rightBinaryCode: %s\n", rightOperand, rightOperand, rightBinaryCode)
	TwoNumberAdditionCarryPrint(leftBinaryCode, rightBinaryCode)
}

func Int8ToBinaryCode(num int) (n int8, binaryCode string) {
	binaryCode = strconv.FormatInt(int64(num), 2)
	binaryCode = fmt.Sprintf("%08s", binaryCode)
	n = int8(num)
	return n, binaryCode
}

func TwoNumberAdditionCarryPrint(left, right string) {
	leftArr := strings.Split(left, "")
	rightArr := strings.Split(right, "")
	fmt.Printf(`
==============Start===============			
left:        %v
right:       %v          
==================================`, leftArr, rightArr)
	rightBit := ""
	leftBit := ""
	carryStr := strings.Split(" ________", "");
	carryFlag := false
	carryRuleNumber := 0
	index := len(leftArr) - 1
	count := 0
	for ;index > -1; index-- {
		rightBit = rightArr[index]
		leftBit = leftArr[index]
		carryRuleNumber = GetCarryRule(carryFlag, leftBit, rightBit)
		switch carryRuleNumber {
		case 4:
			carryFlag = true
			carryStr[index + 1] = "C"
			rightArr[index] = "0"
			leftArr[index] = "0"
		case 8:
			fallthrough
		case 6:
			carryFlag = true
			carryStr[index + 1] = "C"
			leftArr[index] = "0"	
		case 7:
			carryFlag = true
			carryStr[index + 1] = "C"
			rightArr[index] = "0"
		case 5:
			leftArr[index] = "1"
			fallthrough 
		case 1:
			fallthrough 
		case 2:
			fallthrough 
		case 3:
			carryFlag = false
		}
		fmt.Printf(`
       %d번 째 Bit
carryflag: %t
carry:     %v
left:        %v
right:       %v       %d번 규칙
==================================`, count, carryFlag, carryStr, leftArr, rightArr, carryRuleNumber)
		count++
		carryStr[index + 1] = "_"
	}
}

func GetCarryRule(flag bool, left, right string) int {
	ruleNum := 0
	switch {
	case flag == false && left == "0" && right == "0":
		ruleNum = 1
	case flag == false && left == "0" && right == "1":		
		ruleNum = 2
	case flag == false && left == "1" && right == "0":		
		ruleNum = 3
	case flag == false && left == "1" && right == "1":		
		ruleNum = 4
	case flag == true && left == "0" && right == "0":		
		ruleNum = 5	
	case flag == true && left == "1" && right == "0":		
		ruleNum = 6
	case flag == true && left == "0" && right == "1":		
		ruleNum = 7
	case flag == true && left == "1" && right == "1":		
		ruleNum = 8
	}
	return ruleNum
}