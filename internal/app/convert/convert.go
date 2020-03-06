package convert

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func ConvertContent(c *cli.Context) error {
	host := c.String("minio-host")
	port := c.String("minio-port")
	ak := c.String("minio-access-key")
	sk := c.String("minio-secret-key")
	b := c.String("minio-bucket")
	id := c.String("user-id")
	command := exec.Command("node-cli", "convert", "--minioHost", host, "--minioPort", port, "--accessKey", ak, "--secretKey", sk, "--bucket", b, "--userId", id)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}
