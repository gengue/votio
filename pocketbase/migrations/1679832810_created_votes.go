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
			"id": "oxfju3gybxaymn2",
			"created": "2023-03-26 12:13:30.530Z",
			"updated": "2023-03-26 12:13:30.530Z",
			"name": "votes",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "93q6huej",
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
					"id": "fikjwoa8",
					"name": "voter",
					"type": "relation",
					"required": true,
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
					"id": "igwgwp3d",
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

		collection, err := dao.FindCollectionByNameOrId("oxfju3gybxaymn2")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
