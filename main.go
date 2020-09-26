package main

import (
	"github.com/ohmyray/my-blog/model"
	"github.com/ohmyray/my-blog/route"
)

func main() {
	model.InitConnection()
	route.InitRouter()
}
