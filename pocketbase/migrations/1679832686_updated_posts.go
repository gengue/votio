package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1ydrsvl5hd30ikc")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("a9m9pyfn")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1ydrsvl5hd30ikc")
		if err != nil {
			return err
		}

		// add
		del_votes := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "a9m9pyfn",
			"name": "votes",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": 0,
				"max": null
			}
		}`), del_votes)
		collection.Schema.AddField(del_votes)

		return dao.SaveCollection(collection)
	})
}
