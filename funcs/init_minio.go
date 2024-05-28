package funcs

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var bucketName = "testbucket"

func InitMinio() {

	endpoint := "123.56.186.207:50003"
	accessKeyID := "AyDByj0p9g7z4Sfo6vl2"
	secretAccessKey := "aUedVd98Tf8WVjNbvllhlt9Cz8nAC87KGSbB07j9"
	useSSL := false

	var err error
	// Initialize minio client object.
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalf("Error creating minio client: %s", err)
	}

}
