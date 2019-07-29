package main

import (
	"context"
	"fmt"
	"net/http"
)

func foox(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "userId", 1)
		ctx = context.WithValue(ctx, "pwd", 1)
		if ctx.Value("userId") != nil {
			testTwo(w, r)
			return
		}
		next(w, r.WithContext(ctx))
	}
}

func testOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value("userId"))
}

func testTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("error")
}

func main() {
	http.HandleFunc("/test", foox(testOne))
	http.ListenAndServe(":9090", nil)
}