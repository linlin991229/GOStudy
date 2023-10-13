package study

import (
	"fmt"
	"time"
)

func TestTime() {
	p := fmt.Println

	now := time.Now()

	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// 周一至周日Weekday也可用。

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 两个时间之间的间隔
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// 提前或延后时间
	p(then.Add(diff))
	p(then.Add(-diff))
}
