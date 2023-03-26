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
			"id": "yp7jxnp1c4tw51c",
			"created": "2023-03-26 12:10:23.808Z",
			"updated": "2023-03-26 12:10:23.808Z",
			"name": "tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "udchju2c",
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
					"id": "yuarrwhu",
					"name": "slug",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
					}
				},
				{
					"system": false,
					"id": "tukganju",
					"name": "board",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "b2md91cas926z5q",
						"cascadeDelete": true,
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

		collection, err := dao.FindCollectionByNameOrId("yp7jxnp1c4tw51c")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
