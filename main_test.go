package main_test

import (
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {
	t.Run("run main", func(t *testing.T) {
		fmt.Println("Hello")
	})
}
