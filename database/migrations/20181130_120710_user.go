package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20181130_120710 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20181130_120710{}
	m.Created = "20181130_120710"

	migration.Register("User_20181130_120710", m)
}

// Run the migrations
func (m *User_20181130_120710) Up() {
	m.SQL(`
		CREATE TABLE IF NOT EXISTS tbl_user (
			id integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
			email varchar(500) NOT NULL DEFAULT '' ,
			password varchar(500) NOT NULL DEFAULT '' ,
			profile_id integer NOT NULL UNIQUE
		) ENGINE=INNODB;
    `)

	m.SQL(`
		CREATE TABLE IF NOT EXISTS tbl_profile (
			id integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
			birth_date date NOT NULL,
			gender varchar(10) NOT NULL DEFAULT ''
		) ENGINE=INNODB;
	`)

}

// Reverse the migrations
func (m *User_20181130_120710) Down() {
	m.SQL(" DROP TABLE IF EXISTS `tbl_profile` ")
	m.SQL(" DROP TABLE IF EXISTS `tbl_user` ")

}
