package file

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func FileLoaderInit(findFilePatterns []string) (*fileLoader, error) {
	return new(fileLoader).init(findFilePatterns)
}

func (fl *fileLoader) init(findFilePatterns []string) (*fileLoader, error) {
	fl.fileNames = make([]string, 0, len(findFilePatterns))
	if err := fl.getFileList(findFilePatterns); err != nil {
		return nil, err
	}

	return fl, nil
}

func (fl *fileLoader) Close() {
	close(fl.findWordSender)
}

func (fl *fileLoader) getFileList(findFilePatterns []string) error {
	for _, findFilePattern := range findFilePatterns {
		fileNames, err := filepath.Glob(findFilePattern)
		if len(fileNames) == 0 || err != nil {
			e := fmt.Errorf("not found. findFilePattern: %s", findFilePattern)
			if err != nil {
				e = fmt.Errorf("%s, err: %s", e.Error(), err.Error())
			}
			return e
		}
		fl.fileNames = append(fl.fileNames, fileNames...)
	}
	fl.findWordSender = make(chan []string, 30)
	return nil
}

func (fl *fileLoader) FindWord(word string) string {
	for _, fileName := range fl.fileNames {
		fs := new(fileScanner)
		if err := fs.init(fileName); err != nil {
			return err.Error()
		}
		wg.Add(1)
		go fs.wordScan(word, fl.findWordSender)
	}

	ctx, cencel := context.WithCancel(context.Background())
	go fl.findTextCollect(ctx)

	wg.Wait()
	cencel()

	return fl.getPrintMsg()
}

func (fl *fileLoader) findTextCollect(ctx context.Context) {
	fl.result = make(map[string][]string)

	for {
		select {
		case tuple := <-fl.findWordSender:
			fileName := tuple[0]
			findLine := tuple[1]
			r, ok := fl.result[fileName]
			if !ok {
				r = make([]string, 0)
			}
			r = append(r, findLine)
			fl.result[fileName] = r
		case <-ctx.Done():
			return
		}
	}
}

func (fl *fileLoader) getPrintMsg() string {
	if len(fl.result) == 0 {
		return "not found!!!"
	}

	buff := make([]string, 0)
	for fileName, findLines := range fl.result {
		buff = append(
			buff,
			fileName,
			"-----------------------------------------------------------",
		)

		buff = append(buff, findLines...)
		buff = append(buff, "-----------------------------------------------------------\n")
	}

	return strings.Join(buff, "\n")
}
