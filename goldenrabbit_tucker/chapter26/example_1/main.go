package main

import (
	"fmt"
	"gt/chap26/ex1/ex/file"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	args := os.Args
	fmt.Println(args)
	if len(args) < 3 {
		panic(fmt.Errorf("2개 이상의 실행 인수가 필요합니다. 입력한 인수의 수: %d", len(args)))
	}
	word := args[1]
	filePaths := args[2:]
	fmt.Println("찾으려는 단어: ", word)

	loader, err := file.FileLoaderInit(filePaths)
	if err != nil {
		panic(err)
	}

	defer loader.Close()

	fmt.Println(loader.FindWord(word))
}
