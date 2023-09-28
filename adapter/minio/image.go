package minioadapter

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"hangout/pkg/errmsg"
	"hangout/pkg/object_name"
	"hangout/pkg/richerror"
	"mime/multipart"
	"time"
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

	return objectName, nil
}

func (a Adapter) GetTemporaryProfileImageUrl(ctx context.Context, fileName string) (string, error) {
	const op = "MinioAdapter.GetTemporaryProfileImageUrl"

	expTime := time.Duration(24) * time.Hour

	url, err := a.Conn().PresignedGetObject(ctx, bucketName, fileName, expTime, nil)
	if err != nil {
		return "", richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return url.String(), nil
}

func (a Adapter) DeleteProfileImage(ctx context.Context, fileName string) error {
	const op = "MinioAdapter.DeleteProfileImage"

	if err := a.Conn().RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{}); err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
