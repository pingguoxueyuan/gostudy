package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// go doTimeOutStuff(ctx)
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}
func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logg.Printf("down")
}
