package day_4

import (
	"testing"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
)

func Benchmark(b *testing.B) {
	c := &run{}
	for i := 0; i < b.N; i++ {
		c.Run(helper.DownloadAndRead(4))
	}
}
