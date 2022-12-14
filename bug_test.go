package bug

import (
	atlas "ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/miladibra10/ent-versioned-migration-bug/ent/migrate"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func TestBugPostgres(t *testing.T) {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithGlobalUniqueID(true),
	}

	for version, port := range map[string]int{"14": 5434} {
		t.Run(version, func(t *testing.T) {
			// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
			err = migrate.NamedDiff(ctx, fmt.Sprintf("postgresql://postgres:pass@localhost:%d/test?sslmode=disable", port), "add_last_name_users", opts...)
			if err != nil {
				log.Fatalf("failed generating migration file: %v", err)
			}
		})
	}
}
