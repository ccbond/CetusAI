package server

import "github.com/gin-contrib/cors"

func setupAPIRouters(srv *Server) {
	routerGroup := srv.apiServer.GetAPIRouteGroup()
	routerGroup.Use(cors.Default())
	{
		routerGroup.GET("/logo.png", srv.getLogo)
		routerGroup.GET("/favicon.ico", srv.)
		routerGroup.GET("/opemai.yaml", srv.openaiSpace)
	}

	api := routerGroup.Group("/api/v1")
	cetusApi := api.Group("/cetus")
	{
		cetusApi.GET("/pools_info", srv.getPoolList)
	}
}
