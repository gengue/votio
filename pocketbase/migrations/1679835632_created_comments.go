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
			"id": "s612yvfryrt9596",
			"created": "2023-03-26 13:00:32.364Z",
			"updated": "2023-03-26 13:00:32.364Z",
			"name": "comments",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jtkm3sai",
					"name": "author",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "g8pf6xpy",
					"name": "post",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "1ydrsvl5hd30ikc",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": [
							"title"
						]
					}
				},
				{
					"system": false,
					"id": "egtvvpio",
					"name": "content",
					"type": "editor",
					"required": true,
					"unique": false,
					"options": {}
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

		collection, err := dao.FindCollectionByNameOrId("s612yvfryrt9596")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
