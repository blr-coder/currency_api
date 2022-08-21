package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT)
	defer cancel()

	for _, v := range []string{"task1", "2", "3"} {
		go worker(ctx, v)
	}

	<-ctx.Done()

	fmt.Println("завершение работы всех горутин")
	time.Sleep(time.Second * 2)
}

func worker(ctx context.Context, task string) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("завершение воркера")
			return
		case <-ticker.C:
			// do work
			fmt.Println(task)
		}
	}
}
