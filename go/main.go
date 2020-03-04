package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	pb "github.com/dictyBase/go-genproto/dictybaseapis/content"
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

func main() {
	host := os.Getenv("CONTENT_API_SERVICE_HOST")
	port := os.Getenv("CONTENT_API_SERVICE_PORT")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	check(err)
	defer conn.Close()
	client := pb.NewContentServiceClient(conn)

	f, err := os.Open("../slate")
	check(err)
	files, err := f.Readdir(-1)
	f.Close()
	check(err)

	for _, file := range files {
		jsonFile, err := os.Open("../slate/" + file.Name())
		check(err)
		defer jsonFile.Close()

		c := &Content{}

		byteVal, err := ioutil.ReadAll(jsonFile)
		err = json.Unmarshal([]byte(byteVal), c)
		check(err)
		fmt.Println(c.Data.Attributes.Content)

		id, err := strconv.ParseInt(c.Data.Id, 10, 64)
		check(err)
		uid, err := strconv.ParseInt(c.Data.Attributes.UpdatedBy, 10, 64)
		check(err)

		l, err := client.UpdateContent(context.Background(), &pb.UpdateContentRequest{
			Id: id,
			Data: &pb.UpdateContentRequest_Data{
				Type: c.Data.Type,
				Id:   id,
				Attributes: &pb.ExistingContentAttributes{
					UpdatedBy: uid,
					Content:   c.Data.Attributes.Content,
				},
			},
		})
		check(err)
		log.Printf("successfully updated page content with ID %v \n", l.Data.Id)
	}
}
