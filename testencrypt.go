package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	const data = "hello zzp 章治鹏!"
	const key16 = "123456789000aa11"
	const key32 = key16 + key16
	c, e := aes.NewCipher([]byte(key32))
	fmt.Println(c, e)
}
