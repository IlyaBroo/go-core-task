package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	printType("numDecimal", numDecimal)
	printType("numOctal", numOctal)
	printType("numHexadecimal", numHexadecimal)
	printType("pi", pi)
	printType("name", name)
	printType("isActive", isActive)
	printType("complexNum", complexNum)
	str := convertString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("stringSum:", str)
	sliceRune := []rune(str)
	fmt.Println("runeSlice:", hashWithSalt(sliceRune, "go-2024"))

}

func printType(name string, value interface{}) {
	fmt.Printf("Type of %v: %T\n", name, value)
}

func convertString(values ...interface{}) string {
	var result string
	for _, v := range values {
		result += fmt.Sprintf("%v", v)
	}
	return result
}
func hashWithSalt(v []rune, salt string) string {
	half := len(v) / 2
	runeSliceWithSalt := append(v[:half], append([]rune(salt), v[half:]...)...)
	hash := sha256.New()
	hash.Write([]byte(string(runeSliceWithSalt)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
