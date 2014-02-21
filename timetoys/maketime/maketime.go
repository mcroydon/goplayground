package main

import (
	"fmt"
	"time"
)

// Print the current time with microsecond precision as quick as possible.
func main() {
	for {
		fmt.Println(time.Now().Format(time.StampNano))
	}
}
