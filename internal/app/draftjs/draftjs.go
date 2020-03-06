package draftjs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var slugs = []string{"dsc-intro",
	"dsc-about",
	"dsc-other-materials",
	"dsc-order",
	"dsc-payment",
	"dsc-deposit",
	"dsc-faq",
	"dsc-nomenclature-guidelines",
	"dsc-other-stock-centers"}

func NewMinioClient(c *cli.Context) (*minio.Client, error) {
	return minio.New(
		fmt.Sprintf("%s:%s", c.String("minio-host"), c.String("minio-port")),
		c.String("minio-access-key"),
		c.String("minio-secret-key"),
		false,
	)
}

func GetDraftjsContent(c *cli.Context) error {
	addr := fmt.Sprintf("%s:%s", c.String("content-grpc-host"), c.String("content-grpc-port"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not connect to grpc %s", err), 2)
	}
	defer conn.Close()
	pbClient := pb.NewContentServiceClient(conn)
	minioClient, err := NewMinioClient(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not connect to minio %s", err), 2)
	}
	bucket := c.String("minio-bucket")
	location := c.String("minio-location")
	err = minioClient.MakeBucket(bucket, location)
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(bucket)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucket)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucket)
	}

	// get draftjs content and save as json files
	dir, err := ioutil.TempDir(os.TempDir(), "draftjs")
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not create temp directory %s", err), 2)
	}
	for _, slug := range slugs {
		content, err := pbClient.GetContentBySlug(context.Background(), &pb.ContentRequest{Slug: slug})
		if err != nil {
			return err
		}
		jsonString, err := json.Marshal(content)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("%s.json", slug)
		filenamePath := fmt.Sprintf("%s/%s", dir, filename)
		// write json files to temp directory
		if err := ioutil.WriteFile(filenamePath, jsonString, 0644); err != nil {
			return err
		}
		// upload file to minio
		n, err := minioClient.FPutObject(bucket, fmt.Sprintf("%s/%s", "draftjs", filename), filenamePath, minio.PutObjectOptions{ContentType: "application/json"})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s of size %d\n", filename, n)
	}
	return nil
}
