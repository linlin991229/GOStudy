package utils

import (
	"log"
	"os"
)

var LogInfo *log.Logger

func init() {
	LogInfo = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	// LogError := log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile)
}
