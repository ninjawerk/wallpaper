package models

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/VoidArtanis/wallpaper/shared"
	"fmt"
)

type Categories struct {
	Data []Category
}

type Category struct {
	Title       string
	Description string
	Thumb       string
}

func (this *Categories) Save() (err error) {
	key := datastore.IncompleteKey(this.kind(), nil)
	ctx := context.Background()
	if _, err := shared.DatastoreClient.Put(ctx, key, this); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this *Categories) kind() (str string) {
	return "category"
}

func GetAllCategories()(objs *Categories, err error){
	q := datastore.NewQuery(objs.kind())
	ctx := context.Background()
	var data []Categories
	_, er :=  shared.DatastoreClient.GetAll(ctx , q,&data )
	if er != nil {
 		return nil, er
	}
	return &data[0], nil
}
