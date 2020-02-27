package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		String literals are made up of unicode code points called runes.
		Runes are made of bytes. Rune --> int32.
	*/
	// var (
	// 	dir  = `/Users/local/bin/example.py`
	// 	page = `<html>
	// 			<title>Home</title>
	// 			<body>
	// 				<div>
	// 					<p> Data Dev MD blogs. </p>
	// 				</div>
	// 			</body>

	// 			</html>`
	// 	str = "This is a string."
	// )
	// fmt.Printf("string: %s\n", str)
	// fmt.Printf("File directory: %s\n", dir)
	// fmt.Printf("Sample html: %s\n", page)
	// fmt.Printf("length of html page: %d bytes.\n", len(page))

	// eq := "7 + 13 = "
	// sum := 7 + 13
	// fmt.Println(eq + strconv.Itoa(sum))

	chinese := "你好幸福"

	japanese := "こんにちわ"

	hebrew := "אוניברסיטה"

	fmt.Printf("Number of characters in 你好幸福 are: %d\n", utf8.RuneCountInString(chinese))
	fmt.Printf("Size of 你好幸福:  %d bytes.\n", len(chinese))
	fmt.Printf("Number of characters in こんにちわ are: %d\n", utf8.RuneCountInString(japanese))
	fmt.Printf("Size of こんにちわ: %d bytes.\n", len(japanese))
	fmt.Printf("Number of characters in אוניברסיטה are: %d\n", utf8.RuneCountInString(hebrew))
	fmt.Printf("Size of אוניברסיטה: %d bytes.\n", len(hebrew))
	fmt.Println(len("hétérogénéité"))
	fmt.Println(utf8.RuneCountInString("péripatéticien"))
	fmt.Println(utf8.RuneCountInString("hétérogénéité"))
}
