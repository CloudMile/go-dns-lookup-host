package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
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
	var c = make(chan map[int][]net.IP, times)

	for i := 0; i < times; i++ {
		go dig(url, i, c)
		time.Sleep(10 * time.Microsecond)
	}

	for i := 0; i < times; i++ {
		item := <-c
		for _, date := range item {
			for _, ip := range date {
				v := ip.String()
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

func dig(url string, i int, c chan map[int][]net.IP) {
	ips, _ := net.LookupIP(url)
	mapString := make(map[int][]net.IP)
	mapString[i] = ips
	c <- mapString
}
