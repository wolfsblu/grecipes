package db

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"context"
	"embed"
	"log"
)

//go:embed migrations
var migrationFS embed.FS

func Migrate(dbUrl string) error {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(migrationFS),
	)
	if err != nil {
		return err
	}
	defer func(workdir *atlasexec.WorkingDir) {
		_ = workdir.Close()
	}(workdir)

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return err
	}

	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: dbUrl,
	})
	if err != nil {
		return err
	}

	if len(res.Applied) > 0 {
		log.Printf("applied %d migrations\n", len(res.Applied))
	}
	return nil
}
