package lexical

import (
	"fmt"

	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli"
)

type SlateOutput struct {
	Content string
	Slug    string
}

const query = `SELECT content,slug from content`

func LexicalContent(c *cli.Context) error {
	dburl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.String("user"),
		c.String("pass"),
		c.String("host"),
		c.Int("port"),
		c.String("database"),
	)
	pool, err := pgxpool.New(context.Background(), dburl)
	if err != nil {
		cli.NewExitError(
			fmt.Sprintf("error in creating database connection pool %s", err),
			2,
		)
	}
	so := make([]*SlateOutput, 0)
	if err := pgxscan.Select(context.Background(), pool, &so, query); err != nil {
		return cli.NewExitError(
			fmt.Sprintf("error in running query %s", err),
			2,
		)
	}
	for _, row := range so {
		fmt.Println(row.Slug)
	}
	return nil
}

func LexicalContentFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:     "user,u",
			Usage:    "postgresql database user",
			Required: true,
		},
		cli.StringFlag{
			Name:     "pass,p",
			Usage:    "postgresql database password",
			Required: true,
		},
		cli.StringFlag{
			Name:     "host",
			Usage:    "postgresql database host",
			Required: true,
		},
		cli.IntFlag{
			Name:  "port",
			Usage: "postgresql database port",
			Value: 5432,
		},
		cli.StringFlag{
			Name:     "database,d",
			Usage:    "postgresql database name",
			Required: true,
		},
	}
}
