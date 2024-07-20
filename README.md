# Warframe Market Prime Trash Buyer Go

A Go library for fetching profitable orders from the Warframe Market and generating purchase messages. This library is designed to help users find orders with the best platinum-to-ducats ratio for trading in Warframe.

## Features

- **GetProfitableOrders**: Fetches and filters orders to find the most profitable ones based on the platinum-to-ducats ratio.
- **GeneratePurchaseMessages**: Creates formatted purchase messages for profitable orders.

## Installation

To install the library, use the following command:

```sh
go get github.com/freephoenix888/warframe-market-prime-trash-buyer-go
```

## Usage

### Importing the Package

```go
import "github.com/freephoenix888/warframe-market-prime-trash-buyer-go"
```

### Fetching Profitable Orders

The `GetProfitableOrders` function retrieves orders that offer the best platinum-to-ducats ratio. Here’s how you can use it:

```go
package main

import (
	"fmt"
	"log"

	"github.com/freephoenix888/warframe-market-prime-trash-buyer-go"
)

func main() {
	profitableOrders, err := warframe_market_prime_trash_buyer.GetProfitableOrders()
	if err != nil {
		log.Fatalf("Error fetching profitable orders: %s", err)
	}

	fmt.Printf("Found %d profitable orders.\n", len(profitableOrders))
}
```

### Generating Purchase Messages

Once you have the profitable orders, you can generate purchase messages using the `GeneratePurchaseMessages` function:

```go
package main

import (
	"fmt"
	"log"

	"github.com/freephoenix888/warframe-market-prime-trash-buyer-go"
)

func main() {
	profitableOrders, err := warframe_market_prime_trash_buyer.GetProfitableOrders()
	if err != nil {
		log.Fatalf("Error fetching profitable orders: %s", err)
	}

	messages, err := warframe_market_prime_trash_buyer.GeneratePurchaseMessages(profitableOrders)
	if err != nil {
		log.Fatalf("Error generating purchase messages: %s", err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
```

## How It Works

### GetProfitableOrders

This function:
1. Fetches a list of items from the Warframe Market API.
2. Filters items based on predefined criteria.
3. For each item, retrieves orders and filters them based on several criteria:
   - Must be a sell order.
   - The user must be in-game.
   - Must be a PC platform.
   - Quantity should be 3 or more.
   - Price should be 4 platinum or less.
4. Returns a list of orders that meet these criteria.

### GeneratePurchaseMessages

This function:
1. Takes a list of profitable orders.
2. Creates formatted messages for each order.
3. Returns a list of these messages for use in communication with sellers.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue if you have suggestions or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

