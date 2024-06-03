package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	fmt.Println("Autoclicker started. Press Ctrl+C to stop.")

	for {

		x, y := robotgo.GetMousePos()

		robotgo.MouseClick("left", false)

		fmt.Printf("Clicked at (%d, %d)\n", x, y)
		time.Sleep(time.Microsecond * 80)

	}
}
