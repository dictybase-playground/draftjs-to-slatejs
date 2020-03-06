package minio

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
)

func NewMinioClient(c *cli.Context) (*minio.Client, error) {
	return minio.New(
		fmt.Sprintf("%s:%s", c.String("minio-host"), c.String("minio-port")),
		c.String("minio-access-key"),
		c.String("minio-secret-key"),
		false,
	)
}

func MakeBucket(c *cli.Context, minioClient *minio.Client) error {
	bucket := c.String("minio-bucket")
	location := c.String("minio-location")
	err := minioClient.MakeBucket(bucket, location)
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(bucket)
		if errBucketExists == nil && exists {
			log.Printf("%s already exists \n", bucket)
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created bucket %s\n", bucket)
	}
	return nil
}
