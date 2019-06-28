package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	Quote()
}

func Quote() {
	fmt.Println("--------------------------------------")
	fmt.Println("UTF-8 MaxRune Number is ", utf8.MaxRune)
	for i:=rune(0); i< utf8.MaxRune ; i++  {
		if !strconv.CanBackquote(string(i)) {
			fmt.Printf("%q, ", i)
		}
	}

	fmt.Println("--------------------------------------")
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

	fmt.Println("--------------------------------------")
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

	fmt.Println("--------------------------------------")
	bb := []byte("quote:")
	bb = strconv.AppendQuote(bb, `"Fran & Freddie's Diner"`)
	fmt.Println(string(bb))

	fmt.Println("--------------------------------------")
	c := []byte("rune (ascii):")
	c = strconv.AppendQuoteRuneToASCII(c, '😂')
	fmt.Println(string(c))

	fmt.Println("--------------------------------------")
	d := []byte("quote (ascii):")
	d = strconv.AppendQuoteToASCII(d, `"Fran & Freddie's Diner"😂`)
	fmt.Println(string(d))

	fmt.Println("--------------------------------------")
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))

	fmt.Println("--------------------------------------")
	e := strconv.IsPrint('😂')
	fmt.Println(e)
	bel := strconv.IsPrint('\007')
	fmt.Println(bel)

	fmt.Println("--------------------------------------")
	f := strconv.QuoteRune('☺')
	fmt.Println(f)

	fmt.Println("--------------------------------------")
	j, err := strconv.Unquote(`'\u263a'`)
	if err != nil {
		panic(err)
	}
	fmt.Println(j)

}