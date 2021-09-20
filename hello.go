package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf("Error from server: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(response.Time.Add(7 * time.Hour))
}
