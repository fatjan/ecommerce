package main

import (
	"myproject/project/router"
)

// M map
type M map[string]interface{}

func main() {

	userConfig, _ := CreateUserConfig()
	r := router.New()
	v1 := r.Group("")

	userConfig.Register(v1)

	r.Start(":" + "3133")
}
