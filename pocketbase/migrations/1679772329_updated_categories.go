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

		collection, err := dao.FindCollectionByNameOrId("f0w9htlglwc378l")
		if err != nil {
			return err
		}

		// add
		new_parent := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "h7fyxt5v",
			"name": "parent",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "f0w9htlglwc378l",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"board",
					"name"
				]
			}
		}`), new_parent)
		collection.Schema.AddField(new_parent)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("f0w9htlglwc378l")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("h7fyxt5v")

		return dao.SaveCollection(collection)
	})
}
