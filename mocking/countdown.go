package mocking

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer) {
	fmt.Fprintf(w, "3")
}
