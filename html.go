package main

import (
	"fmt"
	"html"
)

func main() {
	htmlText := `"Fran & Freddie's Diner" <tasty@example.com>`
	htmlText = htmlEscapeString(htmlText)
	fmt.Printf("%s\n", htmlText)
	unes := htmlUnescapeString(htmlText)
	fmt.Printf("%s\n", unes)
}

//html 转义
func htmlEscapeString(htmlText string) string {
	return html.EscapeString(htmlText)
}

//html 取消转义
func htmlUnescapeString(htmlText string) string {
	return html.UnescapeString(htmlText)
}