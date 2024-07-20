package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	pkg "github.com/freephoenix888/warframe-market-prime-trash-buyer-go/pkg"
	"github.com/mgutz/ansi"
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

	// Used to track copied messages
	copiedMessages := make(map[string]struct{})

	// Variable to store the count of hidden messages
	var hiddenMessagesCount int

	for {
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
			loadingSpinner.Stop()
			fmt.Println("There are 0 profitable orders.")
			fmt.Print("Type 'regen' to fetch new orders or 'exit' to quit: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()

			if input == "exit" {
				fmt.Println("Exiting.")
				return
			}
			if input == "regen" {
				continue // Restart the loop to fetch new orders
			}
			fmt.Println("Invalid input.")
			continue
		}

		messages, err := pkg.GeneratePurchaseMessages(profitableOrders, logger)
		if err != nil {
			loadingSpinner.Stop()
			logger.Fatal("failed to generate purchase messages", zap.Error(err))
		}

		loadingSpinner.Stop()

		// Filter out copied messages
		filteredMessages := filterCopiedMessages(messages, copiedMessages)

		// Update hidden messages count
		hiddenMessagesCount = calculateHiddenMessagesCount(messages, copiedMessages)

		fmt.Println("Found", len(filteredMessages), "profitable orders")

		scanner := bufio.NewScanner(os.Stdin)

		var status string

		for {
			// Clear the terminal screen using ansi package
			fmt.Print(ansi.ColorCode("reset"))
			fmt.Print("\033[H\033[2J")

			printMessages(filteredMessages, hiddenMessagesCount)

			if status != "" {
				fmt.Print(status)
			}
			fmt.Print("Enter the number of the message to copy to clipboard, 'regen' to regenerate messages, or 'exit' to quit: ")
			scanner.Scan()
			input := scanner.Text()

			if input == "exit" {
				fmt.Println("Exiting.")
				return
			}

			if input == "regen" {
				break // Exit the inner loop to regenerate messages
			}

			index, err := parseInput(input, len(filteredMessages))
			if err != nil {
				status = "Invalid input."
				continue
			}

			msg := filteredMessages[index-1]
			if err := clipboard.WriteAll(msg); err != nil {
				status = "Error copying to clipboard: " + err.Error()
				continue
			}

			copiedMessages[msg] = struct{}{}
			status = fmt.Sprintf("Message %d copied to clipboard.", index)

			// Update the filtered messages after copying
			filteredMessages = filterCopiedMessages(messages, copiedMessages)
			// Update hidden messages count
			hiddenMessagesCount = calculateHiddenMessagesCount(messages, copiedMessages)
		}
	}
}

// printMessages displays the list of available messages
func printMessages(messages []string, hiddenMessagesCount int) {
	fmt.Printf("There are %d profitable orders. Generated messages", len(messages))
	if hiddenMessagesCount != 0 {
		fmt.Printf("(%d messages are not shown because they were copied before)", hiddenMessagesCount)
	}
	fmt.Print(":\n")
	for i, msg := range messages {
		fmt.Printf("%d: %s\n", i+1, msg)
	}
}

// parseInput parses the user's input and returns the index of the selected message
func parseInput(input string, messagesLength int) (int, error) {
	var index int
	_, err := fmt.Sscanf(input, "%d", &index)
	if err != nil || index < 1 || index > messagesLength {
		return 0, fmt.Errorf("invalid index")
	}
	return index, nil
}

// filterCopiedMessages filters out messages that have been copied before
func filterCopiedMessages(messages []string, copiedMessages map[string]struct{}) []string {
	var filtered []string
	for _, msg := range messages {
		if _, copied := copiedMessages[msg]; !copied {
			filtered = append(filtered, msg)
		}
	}
	return filtered
}

// calculateHiddenMessagesCount calculates the number of hidden messages
func calculateHiddenMessagesCount(messages []string, copiedMessages map[string]struct{}) int {
	count := 0
	for _, msg := range messages {
		if _, copied := copiedMessages[msg]; copied {
			count++
		}
	}
	return count
}
