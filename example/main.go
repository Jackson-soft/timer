package main

import (
	"fmt"
	"time"
	"github.com/Jackson-soft/timer"
)

func main() {
	tt := timer.NewTimer(time.Second, 5)
	tt.Register(timer.Single, 3*time.Second, func(args interface{}) {
		fmt.Println(args)
	}, "fsdfsdfsa")

	for {
	}
}
