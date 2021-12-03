package helper

import (
	"github.com/sirupsen/logrus"
)

func DownloadAndRead(day int, fn func(string)) {
	sCk, err := loadSessionCookie()
	if err != nil {
		logrus.Fatalf("unable to load session cookie: %s", err)
	}
	err = downloadInput(sCk, 2021, day, false, "input.txt")
	if err != nil {
		logrus.Fatalf("unable to download input: %s", err)
	}
	r := openReader("input.txt")
	readLines(r, fn)
}

func ReadTestInput(_ int, fn func(string)) {
	r := openReader("input.test")
	readLines(r, fn)
}
