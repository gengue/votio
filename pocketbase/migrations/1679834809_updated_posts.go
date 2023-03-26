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

		// add
		new_status := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ykisrlxq",
			"name": "status",
			"type": "select",
			"required": true,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"open",
					"under review",
					"planned",
					"in progress",
					"complete",
					"closed."
				]
			}
		}`), new_status)
		collection.Schema.AddField(new_status)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1ydrsvl5hd30ikc")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ykisrlxq")

		return dao.SaveCollection(collection)
	})
}
