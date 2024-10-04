package main

import (
	"fmt"
	"time"
)

func startProcess() {
	fmt.Println("process started at:", time.Now())
}

func main() {
	targetHour, targetMinute, targetSecond := 9, 19, 0

	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			if t.Hour() == targetHour && t.Minute() == targetMinute && t.Second() == targetSecond {
				for i := 0; i < 300; i++ {
					go startProcess()
				}
				return
			}
		}
	}
}
