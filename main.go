package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	//ctx.Deadline()
	ctx2 := context.WithValue(ctx, "key1", "value1")
	ctx3, cancel3 := context.WithTimeout(ctx2, 10*time.Second)
	defer cancel3()
	<-ctx3.Done()
	go func() {
		sl := time.Second
		log.Printf("sleeping for %+v", sl)
		time.Sleep(sl)
		cancel()
	}()
	log.Printf("context3 is done: %+v", ctx3)
}
