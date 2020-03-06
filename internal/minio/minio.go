package minio

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
)

func SetUpMinio(c *cli.Context) (*minio.Client, error) {
	m := minio.Client{}
	minioClient, err := newMinioClient(c)
	if err != nil {
		return &m, cli.NewExitError(fmt.Sprintf("could not connect to minio %s", err), 2)
	}
	err = makeBucket(c, minioClient)
	if err != nil {
		return &m, cli.NewExitError(fmt.Sprintf("error making bucket %s", err), 2)
	}
	return minioClient, nil
}

func newMinioClient(c *cli.Context) (*minio.Client, error) {
	return minio.New(
		fmt.Sprintf("%s:%s", c.String("minio-host"), c.String("minio-port")),
		c.String("minio-access-key"),
		c.String("minio-secret-key"),
		false,
	)
}

func makeBucket(c *cli.Context, minioClient *minio.Client) error {
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
