package main

import (
	"os"
	"time"

	"github.com/whiny-nil/learn-go-with-tests/16-maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
