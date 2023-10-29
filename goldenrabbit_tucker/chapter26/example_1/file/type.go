package file

import "os"

type (
	fileScanner struct {
		fileName string
		*os.File
	}

	fileLoader struct {
		fileNames      []string
		findWordSender chan []string
		result         map[string][]string
	}
)
