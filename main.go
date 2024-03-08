package main

import (
	"context"
	"fmt"
	"log"
	"os"

	. "github.com/openfga/go-sdk/client"
)

func main() {
	url := os.Getenv("FGA_API_URL")
	if len(url) == 0 {
		log.Fatal("Please provide OpenFGA API URL")
	}
	storeID := os.Getenv("FGA_STORE_ID")
	if len(storeID) == 0 {
		log.Fatal("Please provide OpenFGA store ID")
	}
	modelID := os.Getenv("FGA_MODEL_ID")
	if len(modelID) == 0 {
		log.Fatal("Please provide OpenFGA model ID")
	}
	client := getClient(url, storeID, modelID)
	check(client, ClientCheckRequest{User: "user:vfarcic", Relation: "owner", Object: "document:Z"})
	check(client, ClientCheckRequest{User: "user:sfarcic", Relation: "owner", Object: "document:Z"})
	check(client, ClientCheckRequest{User: "user:sfarcic", Relation: "reader", Object: "document:Z"})
	check(client, ClientCheckRequest{User: "user:jdoe", Relation: "reader", Object: "document:Z"})
}

func check(client *OpenFgaClient, body ClientCheckRequest) {
	data, err := client.Check(context.Background()).Body(body).Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
	allowed := "allowed"
	if !*data.Allowed {
		allowed = "denied"
	}
	println(fmt.Sprintf("User %s with %s relation is %s to access %s", body.User, body.Relation, allowed, body.Object))
}

func getClient(url, storeID, modelID string) *OpenFgaClient {
	log.Println("Creating client...")
	client, err := NewSdkClient(&ClientConfiguration{
		ApiUrl:               url,
		StoreId:              storeID,
		AuthorizationModelId: modelID,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}
