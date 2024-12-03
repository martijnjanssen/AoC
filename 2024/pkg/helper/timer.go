package helper

import (
	"fmt"
	"time"
)

func Time() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Took: %s\n", time.Since(start).String())
	}
}
