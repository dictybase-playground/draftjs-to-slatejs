package download

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

type Content struct {
	Data Data `json:"data"`
}

type Data struct {
	Type       string     `json:"type"`
	Id         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	UpdatedBy string `json:"updated_by"`
	Content   string `json:"content"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func DownloadJSON(c *cli.Context) error {
	// connect to content grpc
	host := c.String("content-grpc-host")
	port := c.String("content-grpc-port")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	check(err)
	defer conn.Close()
	pbClient := pb.NewContentServiceClient(conn)

	// set up minio client
	minioClient, err := NewMinioClient(c)
	check(err)
	bucket := c.String("minio-bucket")
	location := c.String("us-east-1")
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

	// create json files for draftjs content
	dir, err := ioutil.TempDir(os.TempDir(), "draftjs")
	check(err)
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
		n, err := minioClient.FPutObject(bucket, filename, filenamePath, minio.PutObjectOptions{ContentType: "application/json"})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s of size %d\n", filename, n)
	}
	// f, err := os.Open("../slate")
	// check(err)
	// files, err := f.Readdir(-1)
	// f.Close()
	// check(err)

	// for _, file := range files {
	// 	jsonFile, err := os.Open("../slate/" + file.Name())
	// 	check(err)
	// 	defer jsonFile.Close()

	// 	c := &Content{}

	// 	byteVal, err := ioutil.ReadAll(jsonFile)
	// 	err = json.Unmarshal([]byte(byteVal), c)
	// 	check(err)
	// 	fmt.Println(c.Data.Attributes.Content)

	// 	id, err := strconv.ParseInt(c.Data.Id, 10, 64)
	// 	check(err)
	// 	uid, err := strconv.ParseInt(c.Data.Attributes.UpdatedBy, 10, 64)
	// 	check(err)

	// 	l, err := client.UpdateContent(context.Background(), &pb.UpdateContentRequest{
	// 		Id: id,
	// 		Data: &pb.UpdateContentRequest_Data{
	// 			Type: c.Data.Type,
	// 			Id:   id,
	// 			Attributes: &pb.ExistingContentAttributes{
	// 				UpdatedBy: uid,
	// 				Content:   c.Data.Attributes.Content,
	// 			},
	// 		},
	// 	})
	// 	check(err)
	// 	log.Printf("successfully updated page content with ID %v \n", l.Data.Id)
	// }

	return nil
}
