package shared

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"log"
)

var(
	DatastoreClient *datastore.Client
)



func  ConnectGDS(projId string){
	var err error
	ctx := context.Background()
	DatastoreClient, err = datastore.NewClient(ctx, projId)

	if err != nil {
		log.Fatalf("Failed to create client gds: %v", err)
	}
}