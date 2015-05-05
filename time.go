package check

import (
	"fmt"
	"time"
)

func GreaterThan(high int) func(int) error {
	return func(int value) {
		if value <= high {
			return fmt.Errorf("must be greater than %d", high)
		}
		return nil
	}
}

func Before(high time.Time) func(t *time.Time) error {
	return func(t *time.Time) error {
		if value.After(high) {
			return fmt.Errorf("must be before %v", high)
		}
		return nil
	}
}

func NotEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("Must not be Empty")
	}
}
