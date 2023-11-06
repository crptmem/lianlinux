package server

import (
	"github.com/charmbracelet/log"
	"net"
	"os"
	"os/signal"
	"strings"
)

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func handleCommand(data string, c net.Conn) {
	var arg = strings.SplitAfter(data, "mode")[1]

	switch firstN(data, 4) {
	case "mode":
		log.Info(arg)
		_, err := c.Write([]byte("OK"))
		if err != nil {
			log.Fatalf("Write error: %v", err)
		}
	default:
		_, err := c.Write([]byte("UNKNOWN"))
		if err != nil {
			log.Fatalf("Write error: %v", err)
		}
	}
}

func unixServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		handleCommand(string(data), c)
	}
}

func Listen() {
	l, err := net.Listen("unix", "/tmp/lianlinux.sock")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)

	go func() {
		for {
			s := <-signalChan
			switch s {
			default:
				log.Info("Quitting...")
				err := l.Close()
				if err != nil {
					log.Fatal("Failed to close socket")
					return
				}
				os.Exit(0)
			}
		}
	}()

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatalf("Accept error: %v", err)
		}

		go unixServer(fd)
	}
}
