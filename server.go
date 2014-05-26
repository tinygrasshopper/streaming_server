package main

import (
		"log"
		"fmt"
		"io"
    "net/http"
		"runtime"
)

//type FeedRecord struct {
   //name string
//}

func worker(processor_channel chan<- string, workers_union chan<- bool) {
	for i := 0; i < 500; i++ {
		log.Printf("Push on channel")
		processor_channel <- fmt.Sprintf("Response number %d\n", i)
	}

	workers_union <- true
}

func monitor_workers_and_close_channel(number_of_workers int, processor_channel chan<- string, workers_union <-chan bool) {
	for i:= 0; i < number_of_workers; i++ {
		<- workers_union 
	}
	close(processor_channel)
}

func handler(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type", "text/json;")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")


	//producer_channel := make(chan FeedRecord)

	processor_channel := make(chan string)
	workers_union := make(chan bool)


	number_of_workers := 5
	for i:= 0; i < number_of_workers; i++ {
		go worker(processor_channel, workers_union)
	}

	go monitor_workers_and_close_channel(number_of_workers, processor_channel, workers_union)

	log.Printf("After go")

	for s := range processor_channel {
		log.Printf("Read %s", s)
		io.WriteString(w, s)
	}

	//for i := 0; i < 50000; i++ {

		//log.Println("Prcoessing")
		//fmt.Fprintf(w, "Response number %d\n", i)
		//flusher.Flush()
		//time.Sleep(50 * time.Millisecond)
	//}
}

func main() {
		runtime.GOMAXPROCS(3)
	  log.Printf("Number of cores %d", runtime.GOMAXPROCS(0));
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
