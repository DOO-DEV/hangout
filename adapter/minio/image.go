package minioadapter

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"hangout/pkg/errmsg"
	"hangout/pkg/object_name"
	"hangout/pkg/richerror"
	"mime/multipart"
)

const (
	bucketName = "profile-images"
)

func (a Adapter) SaveImageIntoStorage(ctx context.Context, file *multipart.FileHeader) (string, error) {
	const op = "MinioAdapter.SaveImage"

	objectName := object_name.GenerateUniqueObjectName(file.Filename)
	src, err := file.Open()
	if err != nil {
		return "", richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgCantPreformFile)
	}
	defer src.Close()

	if _, err := a.client.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	}); err != nil {
		fmt.Println("error", err)
		return "", richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgNotFound)
	}

	objUrl := fmt.Sprintf("%s/%s", a.cfg.Endpoint, objectName)

	return objUrl, nil
}
