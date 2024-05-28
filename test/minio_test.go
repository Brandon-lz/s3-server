package test

import (
    "context"
    "log"

    // "log/slog/internal/buffer"
    "testing"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "github.com/stretchr/testify/require"
)

func TestMinio(t *testing.T) {
    require := require.New(t)

    ctx := context.Background()

    endpoint := "123.56.186.207:50003"
    accessKeyID := "AyDByj0p9g7z4Sfo6vl2"
    secretAccessKey := "aUedVd98Tf8WVjNbvllhlt9Cz8nAC87KGSbB07j9"
    useSSL := false

    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    require.NoError(err)

    // Make a new bucket called testbucket.
    bucketName := "testbucket"
    // location := "us-east-1"

    err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
    if err != nil {
        // Check to see if we already own this bucket (which happens if you run this twice)
        exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
        if errBucketExists == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalln(err)
        }
    } else {
        log.Printf("Successfully created %s\n", bucketName)
    }

    // Upload the test file
    // Change the value of filePath if the file is in another location
    objectName := "/test/testdata.txt"
    filePath := "./testdata.txt"
    contentType := "application/octet-stream"

    // Upload the test file with FPutObject
    info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
    require.NoError(err)

    log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

    // List all objects in the bucket
    // for object := range minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true, WithVersions: false}) {
    //     log.Printf("Object: %s, Size: %d, LastModified: %s VersionID: %s\n", object.Key, object.Size, object.LastModified, object.VersionID)
    // }

	for object := range minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: "/",Recursive: false, WithVersions: false}) {
        log.Printf("Object: %s, Size: %d, LastModified: %s VersionID: %s\n", object.Key, object.Size, object.LastModified, object.VersionID)
    }

    // for object := range minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true, WithVersions: true}) {
    //     log.Printf("Object: %s, Size: %d, LastModified: %s VersionID: %s\n", object.Key, object.Size, object.LastModified, object.VersionID)
    // }


    // Download the file from the bucket to local
    filePath = "./file.txt"
    err = minioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
    require.NoError(err)

    // get file obj
    fileobj, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
    require.NoError(err)
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
    
    t.Logf("file contents: %s\n", string(data))

    // Remove the file from the bucket
    // err = minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{VersionID: "1b90908a-7127-43f9-9239-7a049d50947f"}) // ForceDelete: true is used to delete all versions of the object.
    // err = minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{ForceDelete: true})                                 // ForceDelete: true is used to delete all versions of the object.
    // require.NoError(err)

    // Remove the bucket
    // err = minioClient.RemoveBucket(ctx, bucketName)
    // require.NoError(err)

}
