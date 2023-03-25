package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "f0w9htlglwc378l",
			"created": "2023-03-25 19:24:52.323Z",
			"updated": "2023-03-25 19:24:52.323Z",
			"name": "categories",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "egtqluxh",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "i68ujaru",
					"name": "board",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "b2md91cas926z5q",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": [
							"name"
						]
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("f0w9htlglwc378l")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
