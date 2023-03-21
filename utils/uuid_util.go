package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

func CreateUUID() uint32 {
	id := uuid.New().ID()
	fmt.Println("uuid:", id)
	return id
}
func CreateUUIDToString() string {
	id := shortuuid.New()
	// id: iDeUtXY5JymyMSGXqsqLYX length: 22
	fmt.Println("id:", id, "length:", len(id))

	// V22s2vag9bQEZCWcyv5SzL 固定不变
	id = shortuuid.NewWithNamespace("http://127.0.0.1.com")
	// id: K7pnGHAp7WLKUSducPeCXq length: 22
	fmt.Println("id:", id, "length:", len(id))

	// NewWithAlphabet函数可以用于自定义的基础字符串，字符串要求不重复、定长57
	str := "12345#$%^&*67890qwerty/;'~!@uiopasdfghjklzxcvbnm,.()_+·><"
	id = shortuuid.NewWithAlphabet(str)
	// id: q7!o_+y('@;_&dyhk_in9/ length: 22
	fmt.Println("id:", id, "length:", len(id))
	return id
}
