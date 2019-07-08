package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/dnscache"
)

// Times is default dig times
const Times = 1000

func main() {
	http.HandleFunc("/dns-check", dnsCheck)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func dnsCheck(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	times := Times
	if newTimes, err := strconv.Atoi(r.URL.Query().Get("times")); err == nil {
		times = newTimes
	}

	h := map[string]int{}
	var c = make(chan map[int][]string, times)

	for i := 0; i < times; i++ {
		go dig(url, i, c)
	}

	for i := 0; i < times; i++ {
		item := <-c
		for _, date := range item {
			for _, v := range date {
				if h[v] == 0 {
					h[v] = 1
				} else {
					h[v]++
				}
			}
		}
	}

	fmt.Fprintf(w, "URL: %s\n", url)
	for k, v := range h {
		fmt.Fprintf(w, "IP: %s\tTimes: %d\n", k, v)
	}
}

func dig(url string, i int, c chan map[int][]string) {
	resolver := &dnscache.Resolver{}
	ctx := context.Background()
	addr, _ := resolver.LookupHost(ctx, url)
	resolver.Refresh(true)
	mapString := make(map[int][]string)
	mapString[i] = addr
	c <- mapString
}
