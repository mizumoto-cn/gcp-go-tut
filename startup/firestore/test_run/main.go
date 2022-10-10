package main

import (
	"context"
	"fmt"
	"log"

	firestorex "github.com/mizumoto-cn/gcp-go-tut/firestore"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	projID := "my-project-apigee-364705"

	client, err := firestorex.CreateClient(ctx, projID)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed adding client: %v", err)
	}

	// [START firestore_setup_dataset_pt1]
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding lovelace: %v", err)
	}

	// [END firestore_setup_dataset_pt1]

	// [START firestore_setup_dataset_pt2]
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Matt",
		"last":   "Turing",
		"born":   1912,
	})
	if err != nil {
		log.Fatalf("Failed adding turing: %v", err)
	}
	// [END firestore_setup_dataset_pt2]

	// [START firestore_setup_dataset_read]
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	// [END firestore_setup_dataset_read]
}
