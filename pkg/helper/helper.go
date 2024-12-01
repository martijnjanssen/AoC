package helper

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func DownloadAndRead(day int, fn func(string)) {
	sCk, err := loadSessionCookie()
	if err != nil {
		logrus.Fatalf("unable to load session cookie: %s", err)
	}
	filename := fmt.Sprintf("day_%d/input.txt", day)
	err = downloadInput(sCk, 2024, day, false, filename)
	if err != nil {
		logrus.Fatalf("unable to download input: %s", err)
	}
	r := openReader(filename)
	readLines(r, fn)
}

func ReadTestInput(day int, fn func(string)) {
	r := openReader(fmt.Sprintf("day_%d/input.test", day))
	readLines(r, fn)
}
