package my_logger_test

import (
	"fmt"
	"my_libs/logger"
	"my_libs/logger/log_level"
	"testing"
	"time"
)

func Test_logger(t *testing.T) {
	logger.Init(10240, "./log.txt", 1024*1024*100, 3, log_level.INFO)

	logger.Fatal(5, "sdfajsd;")
	logger.Debug("-----------debugInfo-----------")
	logger.Error("NoMoney")

	time.Sleep(2 * time.Second)
	fmt.Println("finished!")
}
