package main

import (
	"balance-sync/balance"
	"context"
	"crypto/tls"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
)

func main() {
	creds := credentials.NewTLS(&tls.Config{})
	conn, err := grpc.Dial("balance.mwc.telekom.net:443", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	client := balance.NewBalanceClient(conn)

	// Get balance
	resp, err := client.GetBalance(context.Background(), &balance.BalanceRequest{
		Did: "did:sov:2wJPyULfLLnYTEFYzByfUR",
	})
	if err != nil {
		log.Printf("failed to get balance: %s", err)
	} else {
		log.Printf("Current user balance (%v): %v", len(resp.ActionIds), resp.ActionIds)
	}

	// Subscribe for updates
	newId, err := uuid.NewV4()
	subscription, err := client.UpdateBalanceSubscribe(context.Background(), &balance.SubscribeRequest{SubscriberId: newId.String()})
	if err != nil {
		log.Fatalf("Unable to subscribe: %s", err)
	}

	log.Println("Subscribe request successful")

	waitc := make(chan struct{})
	go func() {
		for {
			log.Println("Waiting for updates")
			in, err := subscription.Recv()
			if err == io.EOF {
				log.Println("Server shutdown")
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive an update : %v", err)
			}

			log.Printf("Message from stream: %v", in.Message)
			if in.Update != nil {
				log.Printf("Got update from %s :(%v, %v)", in.Update.UpdatedBy, in.Update.Did, in.Update.ActionIds)
			}

		}
	}()

	// Send update
	_, err = client.UpdateBalance(context.Background(), &balance.BalanceUpdate{
		ActionIds: []string{"get-flyer-1", "get-flyer-77", "share-facebook-3"},
		Did:       "did:sov:2wJPyULfLLnYTEFYzByfUR",
		UpdatedBy: newId.String(),
	})
	if err != nil {
		log.Fatalf("failed to update: %s", err)
	}

	// Wait for updates
	<-waitc
}
