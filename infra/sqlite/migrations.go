package sqlite

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"context"
	"fmt"
	"io/fs"
)

func (s *Store) Migrate() error {
	subFS, err := fs.Sub(migrationFS, "migrations")
	if err != nil {
		return err
	}
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(subFS),
	)
	if err != nil {
		return err
	}

	defer func(workdir *atlasexec.WorkingDir) {
		err = workdir.Close()
	}(workdir)

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return err
	}

	_, err = client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: fmt.Sprintf("sqlite://%s", s.path),
	})
	return err
}
