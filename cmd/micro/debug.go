/*
  Brian: I've rewritten most of this now.
  When running `micro -debug`, a log file will be written to.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/zyedidia/micro/v2/internal/util"
)

const LogPartialPath = ".micro/logs/debug_log.txt"

// NullWriter simply sends writes into the void
type NullWriter struct{}

// Write is empty
func (NullWriter) Write(data []byte) (n int, err error) {
	return 0, nil
}

// InitLog sets up the debug log system for micro if it has been enabled by compile-time variables
func InitLog() {
	if util.Debug == "ON" {
		// BPO: Removing O_TRUNC; I want a persistent log.
		// BPO: Adding option to APPEND
		log_path := GetLogFilePath()
		CreateLogFile(log_path)
		f, err := os.OpenFile(log_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file:\n%v", err)
		}

		log.SetOutput(f)
		// BPO:  Printing current date and time.
		log.Println("Micro started with -debug option.")

		} else {
		log.SetOutput(NullWriter{})
	}
}

func CreateLogFile(path_foo string) {
	// detect if file exists
	var _, err = os.Stat(path_foo)

	// create file if not exists
	if os.IsNotExist(err) {
		fmt.Println("Creating new log file...")

		/// First have to create the parent directories.
		parent_dir, _ := filepath.Abs(filepath.Dir(path_foo))
		os.MkdirAll(parent_dir, os.ModePerm)

		var file, err = os.Create(path_foo)
		fmt.Println("Brian Errors: ", err)
		if isError(err) { return }
		defer file.Close()
	}

	fmt.Println("==> Writing to log file: ", path_foo)
}

func GetLogFilePath() string {
	home_dir, _ := os.UserHomeDir()
	log_path := filepath.Join(home_dir, LogPartialPath)
	return log_path
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}
