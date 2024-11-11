package project

import "embed"

//go:embed internal/db/migrations/*.sql
var EmbedMigrations embed.FS
