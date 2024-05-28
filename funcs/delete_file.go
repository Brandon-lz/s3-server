package funcs

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func DeleteFile(objectName string) error {
	ctx := context.Background()
    return minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{ForceDelete: true})                                 // ForceDelete: true is used to delete all versions of the object.
}