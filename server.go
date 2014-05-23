package main

import (
		"log"
		"fmt"
    "net/http"
		"runtime"
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
		fmt.Fprintf(w, "Response number %d\n", i)
		//flusher.Flush()
		//time.Sleep(50 * time.Millisecond)
	}
}

func main() {
		runtime.GOMAXPROCS(3)
	  log.Printf("Number of cores %d", runtime.GOMAXPROCS(0));
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
