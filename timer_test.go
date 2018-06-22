package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	tt := NewTimer(time.Second, 5)
	tt.Register(Single, 3*time.Second, func(args interface{}) {
		fmt.Println(args)
	}, "fsdfsdfsa")
}
