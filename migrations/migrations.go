package migrations

import "embed"

var (
	//go:embed *.up.sql
	Up embed.FS

	//go:embed *.down.sql
	Down embed.FS
)
