package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
   for i := 0; i < 1000; i++ {
     fmt.Fprintf(w, "Response number %d\n", i)
   }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
