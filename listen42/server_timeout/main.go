package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	TraceId = "trace_id"
)

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(context.Background(), TraceId, rand.Int63())
	a(ctx)
	ranNum := rand.Intn(2)
	if ranNum == 0 {

		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "slow response, %d\n", ranNum)
		fmt.Printf("slow response, %d\n", ranNum)

		return
	}

	fmt.Fprintf(w, "quick response, %d\n", ranNum)
	fmt.Printf("quick response, %d\n", ranNum)
	return
}

func a(ctx context.Context) {

	traceId := ctx.Value(TraceId)
	fmt.Printf("trace_id:%v, process of a\n", traceId)
	b(ctx)
}

func b(ctx context.Context) {

	traceId := ctx.Value(TraceId)
	fmt.Printf("trace_id:%v, process of b\n", traceId)
	c(ctx)
}

func c(ctx context.Context) {

	traceId := ctx.Value(TraceId)
	fmt.Printf("trace_id:%v, process of c\n", traceId)
	d(ctx)
}

func d(ctx context.Context) {

	traceId := ctx.Value(TraceId)
	fmt.Printf("trace_id:%v, process of d\n", traceId)

}

func main() {
	http.HandleFunc("/", lazyHandler)
	http.ListenAndServe(":8000", nil)
}
