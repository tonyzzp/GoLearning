package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"crypto/aes"
)

func padding(data []byte, blockSize int) []byte {
	l := blockSize - len(data)%blockSize
	p := bytes.Repeat([]byte{byte(l)}, l)
	data = append(data, p...)
	return data
}

func unpadding(data []byte) []byte {
	v := data[(len(data) - 1)]
	return data[:(len(data) - int(v))]
}

type DES struct {
	key   [8]byte
	iv    [8]byte
	block cipher.Block
}

func (d *DES) init(key, iv [8]byte) {
	d.key = key
	d.iv = iv
	block, e := des.NewCipher(d.key[:])
	if e != nil {
		panic(e)
	}
	d.block = block
}

func (d *DES) encrypt(data []byte) []byte {
	mode := cipher.NewCBCEncrypter(d.block, d.iv[:])
	data = padding(data, mode.BlockSize())
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return result
}

func (d *DES) decrypt(data []byte) []byte {
	mode := cipher.NewCBCDecrypter(d.block, d.iv[:])
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return unpadding(result)
}

type AES struct {
	iv    []byte
	block cipher.Block
}

func (a *AES) init(key, iv []byte) {
	a.iv = iv
	b, e := aes.NewCipher(key)
	if e != nil {
		panic(e)
	}
	a.block = b
}

func (a *AES) encrypt(data []byte) []byte {
	mode := cipher.NewCBCEncrypter(a.block, a.iv)
	data = padding(data, mode.BlockSize())
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return result
}

func (a *AES) decrypt(data []byte) []byte {
	mode := cipher.NewCBCDecrypter(a.block, a.iv)
	result := make([]byte, len(data))
	mode.CryptBlocks(result, data)
	return unpadding(result)
}

func main() {
	s := "我叫zzp 123"

	d := DES{}
	d.init([8]byte{1, 2, 3, 4, 5, 6, 7, 8}, [8]byte{0, 0, 1, 1, 2, 2, 3, 3})
	result := d.encrypt([]byte(s))
	fmt.Println("加密结果", result)
	result = d.decrypt(result)
	fmt.Println("解密结果", string(result))

	a := AES{}
	a.init([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6}, bytes.Repeat([]byte{byte(1)}, 16))
	result = a.encrypt([]byte(s))
	fmt.Println("加密结果", result)
	result = a.decrypt(result)
	fmt.Println("解密结果", string(result))
}
