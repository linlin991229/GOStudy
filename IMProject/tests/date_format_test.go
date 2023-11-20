package tests

import (
	"fmt"
	"testing"
	"time"

	"lin.com/im/utils"
)

func TestDateFormat(t *testing.T) {
	localDateTime := utils.LocalDateTimeFormat(time.Now())
	timeNow := time.Now()
	fmt.Println("local", localDateTime)
	fmt.Println("time", time.Time(timeNow))

}
