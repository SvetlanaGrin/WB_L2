package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	strEcho := "Halo"
	servAddr := "localhost:8080"
	timeOut := 10
	var err error
	switch os.Args[1] {
	case "go-telnet":
		servAddr = os.Args[2] + ":" + os.Args[3]
		if os.Args[2] == "--timeout" {
			timeOut, err = strconv.Atoi(os.Args[3])
			if err != nil {
				println("timeOut failed:", err.Error())
				return
			}
			timeOut, _ = strconv.Atoi(os.Args[3])
			servAddr = os.Args[4] + ":" + os.Args[5]
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println(servAddr, " ", timeOut)
	conn, err := net.DialTimeout("tcp", servAddr, time.Duration(timeOut)*time.Second)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Введите данные:")
	fmt.Scan(&strEcho)
	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	conn.Close()

}
