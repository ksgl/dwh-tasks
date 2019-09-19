package exporter

import (
	"bufio"
	"log"
	"os"
)

func ExportCSV() {
	threads := 5
	f, _ := os.Open("./src.txt")

	stringsToConvert := make(chan string)
	noMoreStrings := make(chan bool)

	for i := 0; i < threads; i++ {
		go processStrings(stringsToConvert, noMoreStrings)
	}

	// readString(stringsToConvert, f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		stringsToConvert <- scanner.Text()
	}
	// close(stringsToConvert)

	for i := 0; i < threads; i++ {
		noMoreStrings <- true
	}

	f.Close()
}

func readString(stringsToConvert chan string, f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		stringsToConvert <- scanner.Text()
	}
	close(stringsToConvert)
}

func processStrings(stringsToConvert chan string, noMoreStrings chan bool) {
loop:
	for true {
		select {
		case s := <-stringsToConvert:
			log.Println(s)

		case <-noMoreStrings:
			break loop
		}
	}
}
