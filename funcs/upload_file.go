package funcs

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
)


func UploadFile(filedata io.Reader, objectSize int64,objectName string) error {
	ctx := context.Background()
    contentType := "application/octet-stream"
	
    // Upload the test file with FPutObject
    info, err := minioClient.PutObject(ctx, bucketName, objectName, filedata, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println(err)
		return err
	}
    log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return nil
}