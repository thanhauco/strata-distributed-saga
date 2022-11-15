package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	sagaID := flag.String("saga", "", "Saga Instance ID")
	action := flag.String("action", "status", "Action: status | replay | abort")
	flag.Parse()

	if *sagaID == "" {
		fmt.Println("Usage: saga-cli -saga=<saga_id> [-action=status|replay|abort]")
		os.Exit(1)
	}

	fmt.Printf("[saga-cli] Executing '%s' for Saga: %s\n", *action, *sagaID)
}
