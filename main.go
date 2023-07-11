package main

import api "shortLink/src"

func main() {
	db := api.InitDB()
	var Engine = api.Create(db)
	Engine.InitRoute()
	Engine.Run("localhost:8080")
}
