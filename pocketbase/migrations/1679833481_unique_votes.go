package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
        _, err := db.CreateUniqueIndex("votes", "unique_votes_idx", "board", "voter", "post").Execute()
		return err 
	}, func(db dbx.Builder) error {
        _, err := db.DropIndex("votes", "unique_votes_idx").Execute()
		return err 
	})
}
