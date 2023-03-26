package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/gengue/votio/pocketbase/migrations"
)


func goDotEnvVariable(key string) string {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
    app := pocketbase.New()

    migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
        Automigrate: true, // auto creates migration files when making collection changes
    })

    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        println("checking if default admin exists...")

        record, err := app.Dao().FindFirstRecordByData("posts", "slug", "emojis-support")
        if err != nil {
            email := goDotEnvVariable("DEFAULT_ADMIN_EMAIL")
            pwd := goDotEnvVariable("DEFAULT_ADMIN_PASSWORD")

            admin := &models.Admin{ Email: email }
            admin.SetPassword(pwd)

            admErr := app.Dao().SaveAdmin(admin)
            if admErr != nil {
                println("Error creating admin")
                log.Fatal(admErr)
            } else {
                println("Admin created")
                return nil;
            }

            log.Fatal(err)
            return err;
        }
        println("record", &record)
        println("admin found...skiping...")

        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }

}
