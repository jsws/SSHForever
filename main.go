package main

import (
	"flag"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	listenSocket := flag.String("l", "0.0.0.0:22", "Socket to listen for SSH connections on.")
	metricsSocket := flag.String("m", "0.0.0.0:2112", "Socket to run the Prometheus metrics server on.")

	flag.Parse()
	registerMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Started metrics server on %s", *metricsSocket)
	go http.ListenAndServe(*metricsSocket, nil)

	ln, err := net.Listen("tcp", *listenSocket)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Started listening for ssh on %s", *listenSocket)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	connectedIP := conn.RemoteAddr().(*net.TCPAddr).IP

	log.Printf("%s connected\n", connectedIP)
	currentConnections.Inc()
	defer conn.Close()

	startTime := time.Now()
	for {
		line := randomString()
		for _, char := range line {
			_, err := conn.Write([]byte(string(char)))

			if err != nil {
				// Client disonnected.
				now := time.Now()
				elapsed := now.Sub(startTime)
				currentConnections.Dec()
				log.Printf("Wasted %s of %s's time\n", elapsed, connectedIP)
				timeWasted.Observe(float64(elapsed.Milliseconds()))
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

}

func randomString() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Max length of the random string to generate is 255 but -2 for \r\n.
	const length = 253

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) + "\r\n"
}
