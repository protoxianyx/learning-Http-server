

package main

import (
	"log"
	"os"
)

func WriteLog(s string) {
	file, err := os.OpenFile("Log.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644 )
	if err != nil {
		log.Fatal("Error: \n", err)
	}
	defer file.Close()

	file.WriteString(s)
}