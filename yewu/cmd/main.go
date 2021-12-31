package main

import (
	"yewu/api"
	"yewu/dao"
)

func main() {
	dao.InitDB();
	api.InitEngine()

}
