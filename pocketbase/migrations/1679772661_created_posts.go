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
			"id": "1ydrsvl5hd30ikc",
			"created": "2023-03-25 19:31:01.138Z",
			"updated": "2023-03-25 19:31:01.138Z",
			"name": "posts",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "aesk9blc",
					"name": "title",
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
					"id": "1vtnazsw",
					"name": "details",
					"type": "editor",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "bpceygrj",
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
				},
				{
					"system": false,
					"id": "vdacsj9s",
					"name": "author",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": [
							"email"
						]
					}
				},
				{
					"system": false,
					"id": "amdpsnd0",
					"name": "slug",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
					}
				},
				{
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

		collection, err := dao.FindCollectionByNameOrId("1ydrsvl5hd30ikc")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
