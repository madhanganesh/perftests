package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SongDAO struct
type SongDAO struct {
	db *mgo.Database
}

// FindAll function
func (dao *SongDAO) FindAll() ([]Song, error) {
	var songs []Song
	err := dao.db.C("songs").Find(bson.M{}).All(&songs)
	return songs, err
}
