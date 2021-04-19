package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type ort struct {
	line int
}

func (o ort) Error() string {
	return fmt.Sprintf("ort in line %d", o.line)
}

func (o ort) Unwrap() error {
	return errors.New("lorem ipsum")
}

func bar() error {
	_, err := time.Parse(time.RFC3339, "maryna")
	return errors.Wrap(err, "could not bar when maryna")
}

func main() {
	if err := bar(); err != nil {
		err2 := errors.Wrapf(err, "error in bar %d", 3)
		fmt.Println(err2)
		inner := errors.Unwrap(err)
		fmt.Println(inner)
	}
	wrapped := errors.Wrap(context.Canceled, "something")
	fmt.Println(wrapped == context.Canceled)
	fmt.Println(errors.Is(wrapped, context.Canceled))
}
