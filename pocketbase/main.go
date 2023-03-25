package main

import (
    "log"

    "github.com/pocketbase/pocketbase"
    // "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/plugins/migratecmd"

    _ "github.com/gengue/votio/pocketbase/migrations"
)

func main() {
    app := pocketbase.New()

    migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
        Automigrate: true, // auto creates migration files when making collection changes
    })

    // app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
    //     record, err := app.Dao().FindFirstRecordByData("posts", "slug", "emojis-support")
    //     if err != nil {
    //         log.Fatal(err)
    //         return err;
    //     }
    //     println("checking if admin exists...")
    //     println("record", record.GetString("name"))
    //
    //     return nil
    // })
    
    // record, err := app.Dao().FindAdminByEmail("admin@votio.com")
    // if err != nil {
    //   log.Fatal(err)
    //   return
    // }
    //
    // println("record")
    // println(record)


    if err := app.Start(); err != nil {
        log.Fatal(err)
    }

}
