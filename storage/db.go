// Copyright 2014 Krister Svanlund
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// The storage package manages setting up the database ond knows how to fetch
// and update all relevant data.
package storage

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "fmt"
)

var (
    db *gorm.DB
)


func GetDatabase() *gorm.DB {
//    if db == nil {
        if gdb, err := gorm.Open("postgres", "user=matrixorg dbname=matrixorg password=matrixorg"); err != nil {
            panic("matrix: could not open database: " + err.Error())
        } else {
            db = &gdb
        }
//    }
    return db
}

func Setup() error {

    GetDatabase()

    tx := db.Begin()
    tx.CreateTable(&User{})
    tx.CreateTable(&Profile{})

    tx.Model(&User{}).Related(&Profile{})
    tx.Model(&Profile{}).Related(&User{}, "UserId")

    tx.Create(&User{UserId:"foo",Password:"foo", Profiles: []Profile{Profile{DisplayName:"FooBAR"}}})

    tx.Commit()
    foo := User{}

    db.First(&foo)

    fmt.Println(foo.Profiles)

//    tx.Rollback()


    return nil
//    var db_tables map[string]string = map[string]string{
//        "user_table":         user_table,
//        "rooms_table":        rooms_table,
//        "events_table":       events_table,
//        "profile_table":      profile_table,
//        "presence_table":     presence_table,
//        "access_token_table": access_token_table,
//    }
//
//    db := GetDatabase()
//    tx, err := db.Begin()
//    if err != nil {
//        fmt.Println(err)
//        panic("Could not open database transaction")
//    }
//    for name, table := range db_tables {
//        fmt.Printf("matrix: setting up DB table %v\n", name)
//        if _, err := tx.Exec(table); err != nil {
//            tx.Rollback()
//            panic("Could not setup " + name + ": " + err.Error())
//        }
//    }
//    tx.Commit()
//    return nil
}
