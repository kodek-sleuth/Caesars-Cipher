package main

import (
	"fmt"
	"unicode"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	var sentence []rune
	for _, letter := range s {
		if unicode.IsMark(letter) || unicode.IsPunct(letter) {
			sentence = append(sentence, letter)
		}
		if unicode.IsUpper(letter) == true && letter+k > 90 {
			sentence = append(sentence, letter+k-26)
		}
		if unicode.IsLower(letter) == true && letter+k > 122 {
			sentence = append(sentence, letter+k-26)
		}
		if unicode.IsLetter(letter) && unicode.IsLetter(letter+k) {
			sentence = append(sentence, letter+k)
		}
	}
	var str = string(sentence)
	return str

}

func main() {
	str := caesarCipher("Ciphering.", 26)
	fmt.Println(str)
}