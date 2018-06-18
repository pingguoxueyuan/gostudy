package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2, syscall.SIGKILL)
	for {
		sig := <-ch
		fmt.Printf("signal: %v\n", sig)

		switch sig {
		case syscall.SIGKILL:
			fmt.Printf("receive sigkill\n")
		case syscall.SIGINT:
			fmt.Printf("receive sigint\n")
		case syscall.SIGTERM:
			fmt.Printf("receive sigterm\n")
			return
		case syscall.SIGUSR2:
			// reload
			fmt.Printf("receive sigusr2\n")
			return
		}
	}
}
