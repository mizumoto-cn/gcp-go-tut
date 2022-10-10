package main

import (
	"context"

	firestorex "github.com/mizumoto-cn/gcp-go-tut/firestore"
)

func main() {
	ctx := context.Background()
	projID := "my-project-apigee-364705"

	client, err := firestorex.CreateClient(ctx, projID)

}
