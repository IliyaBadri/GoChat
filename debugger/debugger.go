package debugger

import (
	"fmt"
	"gochat/cryptography"
	"os"
	"path/filepath"
)

var debugging bool
var logging bool

func Initialize(enableDebugging bool, enableLogging bool, logsDirectoryPath string) {
	debugging = enableDebugging
	if enableDebugging {
		_, err := os.Stat(logsDirectoryPath)
		if os.IsNotExist(err) {
			err := os.Mkdir(logsDirectoryPath, os.ModeDir)
			if err != nil {
				logging = false
				return
			}
		}
		loggingFilePath := filepath.Join(logsDirectoryPath, fmt.Sprintf("%s.log", cryptography.GenerateFileID()))

		loggingFile, err := os.OpenFile(loggingFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			logging = false
			return
		}
	}
}

func Log(log string) {
	if !debugging {
		return
	}
	println(log)
	if !logging {
		return
	}
}
