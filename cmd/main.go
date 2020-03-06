package main

import (
	"log"
	"os"

	"github.com/dictybase-playground/draftjs-to-slatejs/internal/app/download"
	"github.com/dictybase-playground/draftjs-to-slatejs/internal/app/update"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "draftjs-to-slate"
	app.Usage = "cli for replacing draft.js data through grpc"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-format",
			Usage: "format of the logging out, either of json or text.",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "log level for the application",
			Value: "error",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "download-json",
			Usage:  "downloads draft.js content and saves as json files",
			Action: download.DownloadJSON,
			Flags:  getServerFlags(),
		},
		{
			Name:   "update-content",
			Usage:  "updates API with downloaded slate.js content",
			Action: update.UpdateContent,
			Flags:  getServerFlags(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error in running command %s", err)
	}
}

func getServerFlags() []cli.Flag {
	var f []cli.Flag
	f = append(f, minioFlags()...)
	return append(f, grpcFlags()...)
}

func grpcFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "content-grpc-host",
			EnvVar: "CONTENT_API_SERVICE_HOST",
			Usage:  "content grpc host",
		},
		cli.StringFlag{
			Name:   "content-grpc-port",
			EnvVar: "CONTENT_API_SERVICE_PORT",
			Usage:  "content grpc port",
		},
	}
}

func minioFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "minio-host",
			EnvVar: "MINIO_SERVICE_HOST",
			Usage:  "minio host",
		},
		cli.StringFlag{
			Name:   "minio-port",
			EnvVar: "MINIO_SERVICE_PORT",
			Usage:  "minio port",
		},
		cli.StringFlag{
			Name:  "minio-access-key",
			Usage: "minio access key",
		},
		cli.StringFlag{
			Name:  "minio-secret-key",
			Usage: "minio secret key",
		},
		cli.StringFlag{
			Name:  "minio-bucket",
			Usage: "minio bucket",
			Value: "draftjs",
		},
		cli.StringFlag{
			Name:  "minio-location",
			Usage: "minio location",
			Value: "us-east-1",
		},
	}
}
