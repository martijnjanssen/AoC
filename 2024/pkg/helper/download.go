package helper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func loadSessionCookie() (string, error) {
	f, err := os.ReadFile("../session.ck")
	if err != nil {
		return "", err
	}

	ck := strings.Split(string(f), "\n")[0]

	return ck, nil
}

func downloadInput(sessionCookie string, year int, day int, force bool, output string) error {
	flags := os.O_WRONLY | os.O_CREATE
	if force {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_EXCL
	}

	file, err := os.OpenFile(output, flags, 0666)
	if os.IsExist(err) {
		// fmt.Printf("File %q already exists, not attempting download\n", output)
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	client := new(http.Client)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", sessionCookie
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
