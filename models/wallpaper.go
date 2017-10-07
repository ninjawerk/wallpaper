package models

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/BeastSanchez/wallpaper/shared"

	"fmt"
	"github.com/nu7hatch/gouuid"
	"time"
	"strings"
)

type Wallpaper struct {
	Title        string
	Id           string
	Description  string
	Tags         []string
	Sizes        []Size
	Thumb        string
	DisplayImage string
	Date         time.Time
}

type Size struct {
	Size string
	Url  string
}

func (this *Wallpaper) Save() (err error) {
	var id *uuid.UUID
	id, _ = uuid.NewV4()
	urlId := id.String()
	this.Id = urlId
	this.Date = time.Now()

	key := datastore.IncompleteKey(this.kind(), nil)
	ctx := context.Background()
	if _, err := shared.DatastoreClient.Put(ctx, key, this); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this *Wallpaper) kind() (str string) {
	return "wallpaper"
}

func GetFromTag(tag string, offset int, pageSize int) (objs *[]Wallpaper, err error) {
	ctx := context.Background()
	q := datastore.NewQuery("wallpaper")
	q = q.Filter("Tags =", strings.ToLower(tag))

	q = q.Offset(offset)
	q = q.Limit(pageSize)

	var data []Wallpaper
	_, er := shared.DatastoreClient.GetAll(ctx, q, &data)

	if er != nil {
		return nil, er
	}
	return &data, nil
}

func GetFromTagCount(tag string) (total int, err error) {
	ctx := context.Background()
	q := datastore.NewQuery("wallpaper")
	q = q.Filter("Tags =", strings.ToLower(tag))

	//get count
	count, erc := shared.DatastoreClient.Count(ctx, q)
	if erc != nil {
		return 0, erc
	}
	return count, nil
}

func GetFromId(id string) (objs *Wallpaper, err error) {
	ctx := context.Background()
	q := datastore.NewQuery("wallpaper")
	q = q.Filter("Id =", id)

	var data []Wallpaper
	_, er := shared.DatastoreClient.GetAll(ctx, q, &data)

	if er != nil {
		return nil, er
	}

	if len(data) > 0 {
		return &data[0], nil
	} else {
		return nil, nil
	}
}

func GetRecentWallpapers(offset int, pageSize int) (objs *[]Wallpaper, err error) {
	ctx := context.Background()
	q := datastore.NewQuery("wallpaper")
	q = q.Order("-Date")
	q = q.Offset(offset)
	q = q.Limit(pageSize)
	var data []Wallpaper
	_, er := shared.DatastoreClient.GetAll(ctx, q, &data)

	if er != nil {
		return nil, er
	}
	return &data, nil
}

func GetTotalRecentWallpapers() (int, error) {
	ctx := context.Background()
	q := datastore.NewQuery("wallpaper")
	//get count
	count, erc := shared.DatastoreClient.Count(ctx, q)
	if erc != nil {
		return 0, erc
	}
	return count, nil
}
