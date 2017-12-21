// package command provides methods for base shell and file command processing
package command

import (
	"bufio"
	"fmt"
	"os"
)

// FileCommandProcessor to store file name for processing
type FileCommandProcessor struct {
	fileName string
}

// NewFilecommandProcessor to create command file processor
func NewFilecommandProcessor(fileName string) *FileCommandProcessor {
	return &FileCommandProcessor{
		fileName: fileName,
	}
}

// Process method to process command file
func (fcp *FileCommandProcessor) Process() error {
	commandFile, err := os.Open(fcp.fileName)
	if err != nil {
		return err
	}
	defer commandFile.Close()

	commandScanner := bufio.NewScanner(commandFile)
	commandMgr := NewManager()
	var commandString string
	for commandScanner.Scan() {
		commandString = commandScanner.Text()
		out, err := commandMgr.Run(commandString)
		processOutput(out, err)
	}

	if err := commandScanner.Err(); err != nil {
		return err
	}
	return nil
}

// processOutput will output user according to valid output, error
func processOutput(out string, err error) {
	if nil == err {
		fmt.Println(out)
	} else {
		fmt.Println(err)
	}
}
