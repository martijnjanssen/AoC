package helper

import (
	"bufio"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func openReader(file string) *bufio.Reader {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Fatalf("File \"%s\" does not exist", file)
	} else if err != nil {
		log.Fatalf("Unable to check for file existence: %s", err)
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	return bufio.NewReader(f)
}

func readLines(r *bufio.Reader, fn func(string)) {
	for {
		line, err := r.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Unable to read line: %s", err)
		}
		line = strings.TrimSuffix(line, "\n")

		fn(line)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error while reading line: %s", err)
		}
	}
}
