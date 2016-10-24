package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var server string
var client string
var name string
var seed int
var frequency int
var fio *FileIO

func config() {
	flag.StringVar(&server, "s", "localhost:1600", "server service")
	flag.StringVar(&client, "c", "localhost:1700", "client service")
	flag.StringVar(&name, "n", "test", "cell name")
	flag.IntVar(&seed, "x", 77, "seed value")
	flag.IntVar(&frequency, "f", 10, "send frequency (seconds)")
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func xlog(msg string) {
	fmt.Println(fmt.Sprintf("%s: %s", name, msg))
}

func currentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func main() {
	fio = NewFileIO()
	fpath := fmt.Sprintf("%s/sample.pdf", currentDir())
	fio.ReadLines(fpath)

	config()
	rand.Seed(int64(seed))
	listener, err := net.Listen("tcp", server)
	checkError(err)
	go sendRequests()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleRequest(conn)
	}
}

func sendRequests() {
	ticker := time.NewTicker(time.Second * time.Duration(frequency))
	for _ = range ticker.C {
		conn, err := net.Dial("tcp", client)
		if err == nil {
			writeRequest(conn)
		} else {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		}
	}
}

func randomMessage() string {
	return encryptMessage(fio.AssembleRandomString())
}

func writeRequest(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte(fmt.Sprintf("SENDER: %s, MESSAGE: %s", name, randomMessage())))
	xlog("msg sent")
	if err != nil {
		return
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		xlog(fmt.Sprintf("msg rcvd: %d bytes", n))
		// TODO: Randomly decode and save text to file. (new log or default log)
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
