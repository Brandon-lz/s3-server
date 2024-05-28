package funcs

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func Get_file(objectName string) ([]byte, error) {
	// get file obj
	ctx := context.Background()
	fileobj, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return []byte{}, err
	}
	defer fileobj.Close()

	// Read the contents of the file
	var data = make([]byte, 1024)
	var d = make([]byte, 1024)
	for {
		n, err := fileobj.Read(d)
		data = append(data, d[:n]...)
		if err != nil {
			break
		}
	}
	return data, nil
}
