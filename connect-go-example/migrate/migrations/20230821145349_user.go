package migrations

import (
	"context"
	"example/repository"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*repository.User)(nil)).
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*repository.User)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
