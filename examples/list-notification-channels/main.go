package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdk "github.com/EmailGuard/emailguard-go"
)

func main() {
	client := sdk.New(os.Getenv("EMAILGUARD_API_KEY"))
	ctx := context.Background()
	result, err := client.Channels.ListNotificationChannels(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
