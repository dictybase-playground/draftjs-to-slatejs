package update

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
	m "github.com/dictybase-playground/draftjs-to-slatejs/internal/minio"
	"github.com/dictybase-playground/draftjs-to-slatejs/internal/slugs"
	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

type Content struct {
	Data Data `json:"data"`
}

type Data struct {
	Type       string     `json:"type"`
	ID         int64      `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	UpdatedBy int64  `json:"updated_by"`
	Content   string `json:"content"`
}

func UpdateContent(c *cli.Context) error {
	minioClient, err := m.SetUpMinio(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not set up minio %s", err), 2)
	}
	dir, err := ioutil.TempDir(os.TempDir(), "slatejs")
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not create temp dir %s", err), 2)
	}
	for _, slug := range slugs.DSCSlugs {
		err = downloadFiles(minioClient, slug, dir, c.String("minio-bucket"))
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("error downloading files %s", err), 2)
		}
	}
	files, err := getFilesList(dir)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("getting file content %s", err), 2)
	}
	for _, file := range files {
		err = updateWithSlateContent(c, file.Name(), dir)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("error updating content %s", err), 2)
		}
	}
	return nil
}

func getFilesList(dir string) ([]os.FileInfo, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return files, nil
}

func downloadFiles(minioClient *minio.Client, slug string, dir string, bucket string) error {
	filename := fmt.Sprintf("%s.json", slug)
	filenamePath := fmt.Sprintf("%s/%s", dir, filename)
	objectName := fmt.Sprintf("%s/%s", "slatejs", filename)
	// download files from minio
	err := minioClient.FGetObject(bucket, objectName, filenamePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	log.Printf("Successfully downloaded %s", filename)
	return nil
}

func updateWithSlateContent(c *cli.Context, filename string, dir string) error {
	addr := fmt.Sprintf("%s:%s", c.String("content-grpc-host"), c.String("content-grpc-port"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	pbClient := pb.NewContentServiceClient(conn)
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", dir, filename))
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not open file %s", err), 2)
	}
	defer jsonFile.Close()
	content := &Content{}
	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteVal, content)
	if err != nil {
		return err
	}
	l, err := pbClient.UpdateContent(context.Background(), &pb.UpdateContentRequest{
		Id: content.Data.ID,
		Data: &pb.UpdateContentRequest_Data{
			Type: content.Data.Type,
			Id:   content.Data.ID,
			Attributes: &pb.ExistingContentAttributes{
				UpdatedBy: content.Data.Attributes.UpdatedBy,
				Content:   content.Data.Attributes.Content,
			},
		},
	})
	if err != nil {
		return err
	}
	log.Printf("successfully updated page content with ID %v \n", l.Data.Id)
	return nil
}
