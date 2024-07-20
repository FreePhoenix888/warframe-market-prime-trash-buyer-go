package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	pkg "github.com/freephoenix888/warframe-market-prime-trash-buyer-go/pkg"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger

	logger = zap.NewNop()
	if os.Getenv("LOG_LEVEL") == "debug" {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
	}
	defer logger.Sync()

	loadingSpinner := spinner.New(spinner.CharSets[2], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	loadingSpinner.Prefix = "["
	loadingSpinner.Suffix = "] Processing...\n"
	loadingSpinner.Start()

	profitableOrders, err := pkg.GetProfitableOrders(logger)
	if err != nil {
		loadingSpinner.Stop()
		logger.Fatal("failed to get profitable orders", zap.Error(err))
	}
	if len(profitableOrders) == 0 {
		fmt.Println("There are 0 profitable orders")
		os.Exit(0)
	}

	messages, err := pkg.GeneratePurchaseMessages(profitableOrders, logger)
	if err != nil {
		loadingSpinner.Stop()
		logger.Fatal("failed to generate purchase messages", zap.Error(err))
	}

	loadingSpinner.Stop()

	fmt.Println("Found", len(profitableOrders), "orders")

	for _, message := range messages {
		fmt.Println(message)
	}
}
