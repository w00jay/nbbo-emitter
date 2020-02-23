package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fptr := flag.String("fpath", "23870.csv", "./")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		// inBuf := strings.Split(s.Text(), ",")

		fmt.Println(s.Text())

		// Create a new POST request
		server := os.Getenv("SERVER")
		port := os.Getenv("SERVER_PORT")
		req, err := http.NewRequest("POST", "http://"+server+":"+port, strings.NewReader(s.Text()))
		if err != nil {
			log.Fatal("Error reading request. ", err)
		}

		// Set headers
		req.Header.Set("Content-Type", "text/plain")

		// Set client timeout
		client := &http.Client{Timeout: time.Second * 10}

		// Manual jitter
		r := rand.Intn(1000)
		time.Sleep(time.Duration(r) * time.Millisecond)

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)

		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal("Error reading body. ", err)
		// }
		// fmt.Printf("%s\n", body)

		i++
		if i >= 200 {
			log.Fatal("done with 10")
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
