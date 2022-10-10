package firestore

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
)

func CreateClient(ctx context.Context, projID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projID)
	if err != nil {
		return nil, fmt.Errorf("firestore.NewClient: %v", err)
	}
	return client, nil
}

// you can defer a call to client.Close() to close the client when you're done with it

// you can also set a channel to listen for a signal to close the client
func createClientCloseWhenCloseSignal(ctx context.Context, projID string, closeSignal chan os.Signal) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projID)
	if err != nil {
		return nil, fmt.Errorf("firestore.NewClient: %v", err)
	}
	go func() {
		<-closeSignal
		client.Close()
	}()
	return client, nil
}

// Notice that this might cause untraced goroutines to be left running
