package utils

import (
	"mime/multipart"
	"image/jpeg"

	"github.com/disintegration/imaging"
	"github.com/BeastSanchez/wallpaper/models"
	"github.com/nu7hatch/gouuid"
	"github.com/BeastSanchez/wallpaper/shared"
	"golang.org/x/net/context"
	"image"
	"fmt"
)

type Resolution struct {
	Width  int
	Height int
}

var CommonResolutions = []Resolution{

	//Fullscreen
	{
		1660,
		1200,
	},
	{
		1400,
		1050,
	},
	{
		1280,
		1024,
	},
	{
		1280,
		960,
	},
	{
		1152,
		864,
	},
	{
		1024,
		768,
	},

	//Wide
	{
		3840,
		2400,
	},
	{
		3840,
		2160,
	},
	{
		2560,
		1600,
	},
	{
		2560,
		1440,
	},
	{
		2048,
		1152,
	},
	{
		2560,
		1080,
	},
	{
		2560,
		1024,
	},
	{
		1920,
		1200,
	},
	{
		1920,
		1080,
	},
	{
		1680,
		1050,
	},
	{
		1600,
		900,
	},
	{
		1440,
		900,
	},
	{
		1280,
		800,
	},
	{
		1280,
		720,
	},

	//apple
	{
		2932,
		2932,
	},
	{
		2248,
		2248,
	},
	{
		1280,
		2120,
	},
	{
		1224,
		1224,
	},

}

func CreateSizes(file multipart.File, obj *models.Wallpaper) (err error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	width := img.Bounds().Dx()
	//height := img.Bounds().Dy()

	//create thumbs
	var id *uuid.UUID
	id, _ = uuid.NewV4()
	thumbId := id.String()
	obj.Thumb = "https://storage.googleapis.com/wallpapersanchez/" + thumbId + ".jpg"
	go func() {
		thumb := imaging.Fill(img, 400, 300, imaging.Center, imaging.Lanczos)
		uploadToGCS(thumb, thumbId)
	}()


	id, _ = uuid.NewV4()
	displayImageId := id.String()
	obj.DisplayImage = "https://storage.googleapis.com/wallpapersanchez/" + displayImageId + ".jpg"
	go func() {
		displayImage := imaging.Resize(img, 1440, 0, imaging.Lanczos)
		uploadToGCS(displayImage, displayImageId)
	}()


	for res, _ := range CommonResolutions {
		if (width >= CommonResolutions[res].Width  ) {
			var id *uuid.UUID
			id, _ = uuid.NewV4()
			urlId := id.String()
			targetWidth := CommonResolutions[res].Width
			targetHeight := CommonResolutions[res].Height
			nSize := models.Size{
				Size: fmt.Sprintf("%v", targetWidth) + "x" + fmt.Sprintf("%v", targetHeight),
				Url:  "https://storage.googleapis.com/wallpapersanchez/" + urlId + ".jpg",
			}
			obj.Sizes = append(obj.Sizes, nSize)
			go func() {
				m := imaging.Fill(img, targetWidth, targetHeight, imaging.Center, imaging.Lanczos)
				uploadToGCS(m, urlId)
			}()
		}
	}

	return nil
}

func uploadToGCS(m *image.NRGBA, urlId string) {
	ctx := context.Background()
	gcsObj := shared.Bkt.Object(urlId + ".jpg")
	writer := gcsObj.NewWriter(ctx)
	defer writer.Close()
	jpeg.Encode(writer, m, nil)
}
