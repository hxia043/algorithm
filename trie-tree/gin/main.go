package main

import (
	"gin-trie/router"
	"gin-trie/types"
)

func main() {
	r := router.NewRouter()
	r.AddRoute(types.GETROUTEMETHOD, "/", nil)
	r.AddRoute(types.GETROUTEMETHOD, "/ping", nil)
	r.AddRoute(types.GETROUTEMETHOD, "/pong", nil)
	r.AddRoute(types.GETROUTEMETHOD, "/ping/:name", nil)
	r.AddRoute(types.GETROUTEMETHOD, "/pong/:name", nil)
	r.AddRoute(types.POSTROUTERMETHOD, "/ping/linux", nil)

	r.GetRoute(types.GETROUTEMETHOD, "/ping")
	r.GetRoute(types.GETROUTEMETHOD, "/ping/linux")
	r.GetRoute(types.GETROUTEMETHOD, "/pong")
	r.GetRoute(types.GETROUTEMETHOD, "/pong/linux")
	r.GetRoute(types.POSTROUTERMETHOD, "/ping/linux")

	//fmt.Println(r.RouteCount(types.GETROUTEMETHOD))
}
