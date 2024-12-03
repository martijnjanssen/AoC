package helper

import (
	"bufio"
	"fmt"

	"github.com/sirupsen/logrus"
)

func DownloadAndRead(day int) *bufio.Reader {
	sCk, err := loadSessionCookie()
	if err != nil {
		logrus.Fatalf("unable to load session cookie: %s", err)
	}
	filename := fmt.Sprintf("day_%d/input.txt", day)
	err = downloadInput(sCk, 2024, day, false, filename)
	if err != nil {
		logrus.Fatalf("unable to download input: %s", err)
	}
	return openReader(filename)
}

func ReadTestInput(day int, fn func(string)) {
	r := openReader(fmt.Sprintf("day_%d/input.test", day))
	ReadLines(r, fn)
}
