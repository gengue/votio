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

		collection, err := dao.FindCollectionByNameOrId("yp7jxnp1c4tw51c")
		if err != nil {
			return err
		}

		// add
		new_color := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "olvb9sd0",
			"name": "color",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": 20,
				"pattern": ""
			}
		}`), new_color)
		collection.Schema.AddField(new_color)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yp7jxnp1c4tw51c")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("olvb9sd0")

		return dao.SaveCollection(collection)
	})
}
