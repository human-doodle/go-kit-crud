package main
import (
"database/sql"
"fmt"
"os"
 "github.com/go-sql-driver/mysql"
mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database
func GetMongoDB() *mgo.Database {
host := "localhost"
dbName := "users"
fmt.Println("connection info:", host, dbName)
session, err := mgo.Dial("mongodb://localhost:27017")
if err != nil {
fmt.Println("session err:", err)
os.Exit(2)
}
db = session.DB(dbName)
return db
}
