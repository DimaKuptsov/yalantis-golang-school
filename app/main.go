package main

import (
	converter "app/convert"
	encrypting "app/encrypting"
	translation "app/translation"
	"fmt"
)

func main() {
	inputString := "ANY! raNDoM,s\"triNG\" \"foooorr\" enc'odE; decode AND TranSlate to PiG?Latin"
	fmt.Printf("Input string: \"%s\"\n", inputString)
	encrypter := encrypting.GetEncrypter()
	encryptedString := encrypter.Encrypt(inputString)
	fmt.Printf("Encrypted string: \"%s\"\n", encryptedString)
	decryptedString := encrypter.Decrypt(encryptedString)
	fmt.Printf("Decrypted string: \"%s\"\n", decryptedString)
	translater := translation.GetTranslator()
	translatedString := translater.Translate(inputString)
	fmt.Printf("Translated to pig latin: \"%s\"\n", translatedString)
	firstInputSlice := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	secondInputSlice := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	thirdInputSlice := [][]int{
		[]int{1, 2, 3, 1},
		[]int{4, 5, 6, 4},
		[]int{7, 8, 9, 7},
		[]int{7, 8, 9, 7},
	}
	fourthInputSlice := [][]int{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{20, 21, 22, 23, 24, 7},
		[]int{19, 32, 33, 34, 25, 8},
		[]int{18, 31, 36, 35, 26, 9},
		[]int{17, 30, 29, 28, 27, 10},
		[]int{16, 15, 14, 13, 12, 11},
	}
	errorInputSlice := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3},
	}
	testSlices := [][][]int{
		firstInputSlice,
		secondInputSlice,
		thirdInputSlice,
		fourthInputSlice,
		errorInputSlice,
	}
	for _, sliceForConvert := range testSlices {
		fmt.Printf("Input slice: %v\n", sliceForConvert)
		convertedSlice, err := converter.ConvertMultidimensionalSlice(sliceForConvert)
		if err != nil {
			fmt.Printf("Slice not converted. Reason: %s\n", err.Error())
		} else {
			fmt.Printf("Converted slice: %v\n", convertedSlice)
		}
	}
}
