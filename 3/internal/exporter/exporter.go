package exporter

import (
	"bufio"
	"log"
	"os"
)

func ExportCSV() {
	f, _ := os.Open("./src.txt")
	stringsToConvert := make(chan string)

	go writeCSV(stringsToConvert)
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			stringsToConvert <- scanner.Text()
		}
		close(stringsToConvert)
	}()
}

func writeCSV(stringsToConvert chan string) {
	r := <-stringsToConvert
	log.Println(r)
}
