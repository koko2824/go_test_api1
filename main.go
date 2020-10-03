package main

import (
	"github.com/koko2824/first_postgresql/db"
	"github.com/koko2824/first_postgresql/server"
)

func main() {
	db.Init()
	server.Init()
}