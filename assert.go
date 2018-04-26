package main

import (
	"bytes"
	"encoding/hex"
)

const (
	expectedK0InHexEncoding = "707172737475767778797a7b7c7d7e7f" +
		"80818283000000000000000000000000" +
		"00000000000000000000000000000000" +
		"00000000000000000000000000000000"

	expectedXoredK0InHexEncoding = "46474445424340414e4f4c4d4a4b4849" +
		"b6b7b4b5363636363636363636363636" +
		"36363636363636363636363636363636" +
		"36363636363636363636363636363636"

	expectedXoredK0WithTextInHexEncoding = "46474445424340414e4f4c4d4a4b4849" +
		"b6b7b4b5363636363636363636363636" +
		"36363636363636363636363636363636" +
		"36363636363636363636363636363636" +
		"48656c6c6f20576f726c64"

	expectedSHA1HashInHexEncoding = "0d42b899d804e19ebfd86fc44f414045dfc9e39a"

	expectedK0XoredWithOpadInHexEncoding = "2c2d2e2f28292a2b2425262720212223" +
		"dcdddedf5c5c5c5c5c5c5c5c5c5c5c5c" +
		"5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c" +
		"5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c"

	expectedStep6InHexEncoding = "2c2d2e2f28292a2b2425262720212223" +
		"dcdddedf5c5c5c5c5c5c5c5c5c5c5c5c" +
		"5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c" +
		"5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c" +
		"0d42b899d804e19ebfd86fc44f414045" +
		"dfc9e39a"

	expectedStep7InHexEncoding = "2e492768aa339e32a9280569c5d026262b912431"
)

func assertStep1(actualK0 []byte) {
	expectedK0, err := hex.DecodeString(expectedK0InHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedK0) != len(actualK0) {
		panic("len(expectedK0) != len(actualK0)")
	}

	if !bytes.Equal(expectedK0, actualK0) {
		panic("!bytes.Equal(expectedK0, actualK0)")
	}
}

func assertStep2(actualXoredK0 []byte) {
	expectedXoredK0, err := hex.DecodeString(expectedXoredK0InHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedXoredK0) != len(actualXoredK0) {
		panic("len(expectedXoredK0) != len(actualXoredK0)")
	}

	if !bytes.Equal(expectedXoredK0, actualXoredK0) {
		panic("!bytes.Equal(expectedXoredK0, actualXoredK0)")
	}
}

func assertStep3(actualXoredK0WithText []byte) {
	expectedXoredK0WithText, err := hex.DecodeString(expectedXoredK0WithTextInHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedXoredK0WithText) != len(actualXoredK0WithText) {
		panic("len(expectedXoredK0WithText) != len(actualXoredK0WithText)")
	}

	if !bytes.Equal(expectedXoredK0WithText, actualXoredK0WithText) {
		panic("!bytes.Equal(expectedXoredK0WithText, actualXoredK0WithText)")
	}
}

func assertStep4(actualSHA1Hash []byte) {
	expectedSHA1Hash, err := hex.DecodeString(expectedSHA1HashInHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedSHA1Hash) != len(actualSHA1Hash) {
		panic("len(expectedSHA1Hash) != len(actualSHA1Hash)")
	}

	if !bytes.Equal(expectedSHA1Hash, actualSHA1Hash) {
		panic("!bytes.Equal(expectedSHA1Hash, actualSHA1Hash)")
	}
}

func assertStep5(actualXoredK0 []byte) {
	expectedK0XoredWithOpad, err := hex.DecodeString(expectedK0XoredWithOpadInHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedK0XoredWithOpad) != len(actualXoredK0) {
		panic("len(expectedK0XoredWithOpad) != len(actualXoredK0)")
	}

	if !bytes.Equal(expectedK0XoredWithOpad, actualXoredK0) {
		panic("!bytes.Equal(expectedK0XoredWithOpad, actualXoredK0)")
	}
}

func assertStep6(actualStep6 []byte) {
	expectedStep6, err := hex.DecodeString(expectedStep6InHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedStep6) != len(actualStep6) {
		panic("len(expectedStep6) != len(actualStep6)")
	}

	if !bytes.Equal(expectedStep6, actualStep6) {
		panic("!bytes.Equal(expectedStep6, actualStep6)")
	}
}

func assertStep7(actualStep7 []byte) {
	expectedStep7, err := hex.DecodeString(expectedStep7InHexEncoding)
	if err != nil {
		panic(err)
	}

	if len(expectedStep7) != len(actualStep7) {
		panic("len(expectedStep7) != len(actualStep7)")
	}

	if !bytes.Equal(expectedStep7, actualStep7) {
		panic("!bytes.Equal(expectedStep7, actualStep7)")
	}
}