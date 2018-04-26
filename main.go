package main

import (
	"encoding/hex"
	"crypto/sha1"
)

const (
	text = "Hello World"
	KInHexEncoding = "707172737475767778797a7b7c7d7e7f80818283"

	b = 64 // block size of SHA-1 function
	L = 0
)

func H(msg []byte) []byte { return []byte{} }

func zeros(msg []byte, n int) []byte {
	for i := 0; i < n; i++ {
		msg = append(msg, 0x00)
	}
	return msg
}

func hmac() {
	K, err := hex.DecodeString(KInHexEncoding)
	if err != nil {
		panic(err)
	}

	// Шаг 1. Дополняем ключ K нулевыми байтами до размера блока. Размер блока хэш-функции SHA-1 равен 64 байтам.
	var K0 []byte
	if len(K) == b {
		K0 = K
	}

	if len(K) > b {
		x := H(K) // length( x ) == L
		K0 = zeros(x, b - L)
	}

	if len(K) < b {
		K0 = zeros(K, b - len(K))
	}

	assertStep1(K0)

	// Шаг 2. Выполняем операцию «побитовое исключающее ИЛИ» c константой 0x36.
	K0BitArray := newBitArray(K0)
	ipadBitArray := newBitArray(ipad(b))

	xoredK0 := K0BitArray.xor(ipadBitArray).data

	assertStep2(xoredK0)

	// Шаг 3. Выполняем склейку исходного сообщения со строкой, полученной на шаге 2.
	xoredK0WithText := append(xoredK0, []byte(text)...)

	assertStep3(xoredK0WithText)

	// Шаг 4. Применим хэш-функцию SHA-1 к строке, полученной на прошлом шаге.
	h := sha1.New()
	if _, err := h.Write(xoredK0WithText); err != nil {
		panic(err)
	}
	sha1Hash := h.Sum(nil)

	assertStep4(sha1Hash)

	// Шаг 5. Выполним операцию «побитовое исключающее ИЛИ» c константой 0x5c.
	K0BitArray = newBitArray(K0)
	opadBitArray := newBitArray(opad(b))

	xoredK0 = K0BitArray.xor(opadBitArray).data

	assertStep5(xoredK0)

	// Шаг 6. Склейка строки, полученной на шаге 4, со строкой, полученной на шаге 5.
	step6 := append(xoredK0, sha1Hash...)

	// fmt.Println(hex.EncodeToString(step6))
	assertStep6(step6)

	// Шаг 7. Применим хэш-функцию SHA-1 к строке, полученной на прошлом шаге.
	h = sha1.New()
	if _, err := h.Write(step6); err != nil {
		panic(err)
	}
	sha1Hash = h.Sum(nil)

	// fmt.Println(hex.EncodeToString(sha1Hash))
	assertStep7(sha1Hash)
}

func ipad(n int) []byte {
	data := make([]byte, n)
	for i := 0; i < n; i++ {
		data[i] = 0x36
	}
	return data
}

func opad(n int) []byte {
	data := make([]byte, n)
	for i := 0; i < n; i++ {
		data[i] = 0x5c
	}
	return data
}

func main() {
	hmac()
}