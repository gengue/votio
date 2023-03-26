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

		collection, err := dao.FindCollectionByNameOrId("s612yvfryrt9596")
		if err != nil {
			return err
		}

		// add
		new_parent := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kkic8rrn",
			"name": "parent",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "s612yvfryrt9596",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_parent)
		collection.Schema.AddField(new_parent)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("s612yvfryrt9596")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kkic8rrn")

		return dao.SaveCollection(collection)
	})
}
