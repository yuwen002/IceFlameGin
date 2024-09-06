package main

import "ice_flame_gin/migrations"

func main() {
	migrations.CreateDatabase()
	migrations.SyncData()
}
