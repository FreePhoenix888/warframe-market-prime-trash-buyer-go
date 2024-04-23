package main

import (
	"fmt"
	"log"
	"os"

	"github.com/freephoenix888/warframe-market-prime-trash-buyer-go/lib"
)

func main() {
	env := os.Getenv("ENVIRONMENT")

	if env == "production" {
		file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Failed to open error log file")
		}
		defer file.Close()
		log.SetOutput(file)
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	} else {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	}

	goodOrders, err := lib.GetGoodOrders()
	if err != nil {
		log.Fatal("Error getting good orders:", err)
	}

	messages, err := lib.GenerateMessages(goodOrders)
	if err != nil {
		log.Fatal("Failed to generate messages:", err)

	}

	for _, message := range messages {
		fmt.Println(message)
	}

}
