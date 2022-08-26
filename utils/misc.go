package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func Unwrap[T any](val any) (t T, err error) {
	defer func() {
		if e := recover(); e != nil {
			t = lo.Empty[T]()
			err = fmt.Errorf("Error: %v", e)
		}
	}()

	return val.(T), nil
}

func GenUUID() string {
	return uuid.NewString()
}
