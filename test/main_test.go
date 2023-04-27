package test

import (
	"NewApp/utils"
	"fmt"
	"testing"
)

func TestUUIDCreate(t *testing.T) {
	id := utils.CreateUUIDToString()
	fmt.Println(id)
}
