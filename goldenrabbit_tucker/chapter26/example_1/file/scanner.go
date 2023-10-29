package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (f *fileScanner) init(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	f.fileName = fileName
	f.File = file
	return nil
}

func (f *fileScanner) wordScan(word string, findWordSender chan []string) {
	defer func() {
		wg.Done()
	}()

	defer f.Close()

	lineNum := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(line, word) {
			findWordSender <- []string{f.fileName, fmt.Sprintf("%d\t%s", lineNum, line)}
		}
	}
}
