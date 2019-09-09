package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

// Times is default dig times
const Times = 100

var server = os.Getenv("DNS_SERVER")

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
		if newTimes < Times {
			times = newTimes
		}
	}

	h := map[string]int{}
	var c = make(chan map[int][]dns.RR, times)

	for i := 0; i < times; i++ {
		go dig(url, i, c)
		time.Sleep(10 * time.Microsecond)
	}

	for i := 0; i < times; i++ {
		channel := <-c
		for _, dnsRRs := range channel {
			for _, dnsRR := range dnsRRs {
				arecord, err := dnsRR.(*dns.A)
				if err == false {
					continue
				}
				v := arecord.A.String()
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

func dig(url string, i int, c chan map[int][]dns.RR) {
	client := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(url+".", dns.TypeA)
	r, _, err := client.Exchange(m, server+":53")
	if err != nil {
		log.Fatal(err)
	}
	mapString := make(map[int][]dns.RR)
	mapString[i] = r.Answer
	c <- mapString
}
