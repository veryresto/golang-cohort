package media

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Media interface {
	Upload(file multipart.FileHeader) (*string, *response.Error)
	Delete(file string) (*string, *response.Error)
}

type media struct {
}

// Delete implements Media
func (m *media) Delete(file string) (*string, *response.Error) {
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUDNAME"))

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	fileName := utils.GetFileName(file)

	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: fileName,
	})

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &res.Result, nil
}

// Upload implements Media
func (m *media) Upload(file multipart.FileHeader) (*string, *response.Error) {
	// Define untuk dapat mengakses ke cloudinary
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUDNAME"))

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	binary, err := file.Open()

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	defer binary.Close()

	if binary != nil {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			binary,
			uploader.UploadParams{
				PublicID: uuid.New().String(),
			},
		)

		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err,
			}
		}

		return &uploadResult.SecureURL, nil
	}

	return nil, &response.Error{
		Code: 500,
		Err:  errors.New("tidak dapat membaca file binary"),
	}

}

func NewMedia() Media {
	return &media{}
}
