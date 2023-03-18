package jobs

import (
	"fmt"
	"time"
)

func Countdown(title string, seconds int) {
	for ; seconds > 0; seconds-- {
		fmt.Println(title, ":", seconds)
		time.Sleep(time.Second)
	}
	fmt.Println(title, ":", seconds)
}
