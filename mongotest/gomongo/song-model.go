package main

import "gopkg.in/mgo.v2/bson"

// Song struct
type Song struct {
	ID     bson.ObjectId   `bson:"_id" json:"id"`
	Title  string          `bson:"title" json:"title"`
	Lyrics []bson.ObjectId `bson:"lyrics" json:"lyrics"`
}
