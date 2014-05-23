package main

import (
		//"log"
		"io"
    "net/http"
		//"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain;")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	//defer func(){
	//w.Close()
	//}()

	for i := 0; i < 50000; i++ {
		//log.Println("Prcoessing")
		io.WriteString(w, "Response number")
		//time.Sleep(50 * time.Millisecond)
	}
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
