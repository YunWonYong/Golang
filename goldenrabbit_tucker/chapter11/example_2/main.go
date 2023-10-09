package main

import "fmt"

func main() {
	fmt.Println(join('_', "won", "yong", "yun", "golang", "node", "java", "ts", "php"))
}

func join(separator rune, inputs ...string) string {
	result := inputs[0]
	for _, input := range inputs[1:] {
		result = fmt.Sprintf("%s%c%s", result, separator, input)
	}
	
	return result
}
