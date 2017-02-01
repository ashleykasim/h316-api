package db

import (
    mgo "gopkg.in/mgo.v2"
    "log"
)

var Session *mgo.Session
var Db *mgo.Database

func Connect(conn string, db string) {
  session, err := mgo.Dial(conn)
  if err != nil {
      panic(err)
  }
  Session = session
  Db = session.DB(db)
  log.Println("Connected to Mongodb instance at", conn, "using default database", db)
}
