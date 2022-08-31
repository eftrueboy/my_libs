package goroutine_util_test

import (
	"fmt"
	"mytest/utils/goroutine_util"
	"testing"
	"time"
)

func Test(t *testing.T) {
	for i:=0;i<5;i++ {
		go func() {
			gid := goroutine_util.GetGoroutineId()
			fmt.Println("gid:", gid)
		}()
	}
	time.Sleep(time.Second*5)
}
