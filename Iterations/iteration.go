package main

import "fmt"

func Repeat(char string) string {

	result := ""
	for i := 0; i < 5; i++ {
		result += char
	}
	return result
}
func main() {
	fmt.Println(Repeat("A"))
}
