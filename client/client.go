package main

import (
	"balance-sync/balance"
	"context"
	"crypto/tls"
	"flag"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
)

const DID = "did:sov:2wJPyULfLLnYTEFYzByfUR"

func main() {
	// Flag to switch to local server connection
	boolPtr := flag.Bool("local", false, "Use local connection")
	flag.Parse()

	var conn *grpc.ClientConn
	if *boolPtr {
		log.Println("Using local server")
		conn, _ = grpc.Dial("localhost:50051", grpc.WithInsecure())
	} else {
		log.Println("Using telekom server")
		creds := credentials.NewTLS(&tls.Config{})
		conn, _ = grpc.Dial("balance.mwc.telekom.net:443", grpc.WithTransportCredentials(creds))
	}

	client := balance.NewBalanceClient(conn)

	// Get balance
	balanceRes, err := client.GetBalance(context.Background(), &balance.BalanceRequest{
		Did: DID,
	})
	if err != nil {
		log.Printf("failed to get balance: %s", err)
	} else {
		log.Printf("Current user balance (%v): %v", balanceRes.Balance, balanceRes.Actions)
	}

	// Get milestones
	milestones, err := client.GetMilestones(context.Background(), &balance.BalanceRequest{
		Did: DID,
	})
	if err != nil {
		log.Printf("failed to get milestones: %s", err)
	} else {
		log.Printf("Milestones: %v", milestones)
	}

	// Claim prize
	claimRes, err := client.ClaimPrize(context.Background(), &balance.BalanceRequest{
		Did: DID,
	})
	if err != nil {
		log.Printf("failed to claim prize: %s", err)
	} else {
		log.Printf("Claim success: %v", claimRes.Success)
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
				log.Printf("Got update from %s :(%v, %v)", in.Update.UpdatedBy, in.Update.Did, in.Update.Actions)
			}

		}
	}()

	// Send update
	_, err = client.UpdateBalance(context.Background(), &balance.BalanceUpdate{
		Actions: []*balance.Action{
			{Id: "registration", Timestamp: "2019-01-01"},
			{Id: "checkin-first-skt", Timestamp: "2019-01-02"},
			{Id: "checkin-first-dt", Timestamp: "2019-01-02"},
			{Id: "claim-email-dt", Timestamp: "2019-01-03"},
		},
		Did:       DID,
		UpdatedBy: newId.String(),
	})
	if err != nil {
		log.Fatalf("failed to update: %s", err)
	}

	// Wait for updates
	<-waitc
}
