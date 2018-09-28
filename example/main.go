package main

import (
	"fmt"
	"time"

	"github.com/Jackson-soft/timer"
)

func main() {
	tt := timer.NewTimer(time.Second, 5)
	go func() {
		time.AfterFunc(time.Second, func() {
			tt.Register(timer.Single, 3*time.Second, func(args interface{}) {
				fmt.Println(args)
			}, "fsdfsdfsa1")
		})
	}()

	go func() {
		time.AfterFunc(3*time.Second, func() {
			tt.Register(timer.Single, 3*time.Second, func(args interface{}) {
				fmt.Println(args)
			}, "fsdfsdfsa2")
		})
	}()

	go func() {
		time.AfterFunc(5*time.Second, func() {
			tt.Register(timer.Repetition, 2*time.Second, func(args interface{}) {
				fmt.Println(args)
			}, "fsdfsdfsa3")
		})
	}()

	for {
	}
}
