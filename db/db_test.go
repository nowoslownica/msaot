package db

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDB()
	fmt.Println(err)
}
