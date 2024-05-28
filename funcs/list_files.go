package funcs

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
)

func ListFiles(path string) map[string]int64 {
	// List all objects in the bucket
	if len(path)!=0{
		path = path[1:]
	}
	log.Println(path)
	ctx := context.Background()
	var objs = make(map[string]int64)
	for object := range minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: path, Recursive: false, WithVersions: false}) {
		log.Printf("Object: %s, Size: %d, LastModified: %s VersionID: %s\n", object.Key, object.Size, object.LastModified, object.VersionID)
		objs[object.Key] = object.Size
	}
	return objs
	
}
