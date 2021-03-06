package draftjs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictybase-playground/draftjs-to-slatejs/internal/app/convert"
	m "github.com/dictybase-playground/draftjs-to-slatejs/internal/minio"
	"github.com/dictybase-playground/draftjs-to-slatejs/internal/slugs"
	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func GetDraftjsContent(c *cli.Context) error {
	minioClient, err := m.SetUpMinio(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not set up minio %s", err), 2)
	}
	dir, err := ioutil.TempDir(os.TempDir(), "draftjs")
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not create temp directory %s", err), 2)
	}
	bucket := c.String("minio-bucket")
	for _, slug := range slugs.DSCSlugs {
		if err := getContentJSON(c, slug, dir); err != nil {
			return cli.NewExitError(err, 2)
		}
		if err := uploadFiles(minioClient, slug, dir, bucket); err != nil {
			return cli.NewExitError(err, 2)
		}
	}
	if c.Bool("convert") {
		if err := convert.ConvertContent(c); err != nil {
			return cli.NewExitError(err, 2)
		}
	}
	return nil
}

func uploadFiles(minioClient *minio.Client, slug string, dir string, bucket string) error {
	filename := fmt.Sprintf("%s.json", slug)
	filenamePath := fmt.Sprintf("%s/%s", dir, filename)
	objectName := fmt.Sprintf("%s/%s", "draftjs", filename)
	options := minio.PutObjectOptions{ContentType: "application/json"}
	// upload file to minio
	n, err := minioClient.FPutObject(bucket, objectName, filenamePath, options)
	if err != nil {
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", filename, n)
	return nil
}

func getContentJSON(c *cli.Context, slug string, dir string) error {
	addr := fmt.Sprintf("%s:%s", c.String("content-grpc-host"), c.String("content-grpc-port"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	pbClient := pb.NewContentServiceClient(conn)
	content, err := pbClient.GetContentBySlug(context.Background(), &pb.ContentRequest{Slug: slug})
	if err != nil {
		return err
	}
	jsonString, err := json.Marshal(content)
	if err != nil {
		return err
	}
	filenamePath := fmt.Sprintf("%s/%s.json", dir, slug)
	// write json files to temp directory
	if err := ioutil.WriteFile(filenamePath, jsonString, 0644); err != nil {
		return err
	}
	return nil
}
