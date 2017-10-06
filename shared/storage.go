package shared

import (

	"golang.org/x/net/context"
	"log"
	"cloud.google.com/go/storage"
)

var(
	StorageClient *storage.Client
	Bkt *storage.BucketHandle
)



func  ConnectGCS(projId string){
	var err error
	ctx := context.Background()
	StorageClient, err = storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client gcs: %v", err)
	}

	Bkt = StorageClient.Bucket("wallappercollection")
	
}