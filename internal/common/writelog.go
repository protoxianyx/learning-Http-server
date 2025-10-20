package common

import (
	"log"
	"os"
)

func WriteLog(s, filepath string) {
	file, err := os.OpenFile(filepath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644 )
	if err != nil {
		log.Fatal("Error: \n", err)
	}
	defer file.Close()

	add()

	file.WriteString("READLINE: " + "\n" + s + "\n")
}