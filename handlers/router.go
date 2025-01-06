package handlers

import (
	"net/http"

	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/service"
)

type Handler struct {
	manager service.Service
}

func New(manager service.Service) *Handler {
	return &Handler{
		manager: manager,
	}
}

func (ctrl *Handler) platformAuthMiddleware(c *fox.Context) (res any) {
	// TODO(m)
	return nil
}

func (ctrl *Handler) clientAuthMiddleware(c *fox.Context) error {
	// TODO(m)
	return nil
}

func (ctrl *Handler) RegisterHandler() http.Handler {

	router := fox.Default()

	{
		configs := router.Group("/confly/v1/configs", ctrl.clientAuthMiddleware)
		configs.GET("/groups/:group_name/apps/:app_name/configsets/:configset_name", ctrl.GetConfiguration)
		configs.POST("/groups/:group_name/apps/:app_name/configsets/:configset_name/instances", ctrl.ReportInstance)
	}

	router.Use(ctrl.platformAuthMiddleware)

	router.GET("/confly/v1/signin", ctrl.Signin)
	router.GET("/confly/v1/logout", ctrl.Logout)
	router.GET("/confly/v1/boot", ctrl.Boot)

	system := router.Group("/confly/v1/system")
	system.GET("/users", ctrl.GetUsers)
	system.PATCH("/users/:username", ctrl.UpdateUser)

	groups := router.Group("/confly/v1/groups")
	groups.GET("", ctrl.GetGroups)
	groups.POST("", ctrl.CreateGroup)

	group := router.Group("/confly/v1/groups/:group_name", ctrl.SetGroup)
	group.GET("", ctrl.GetGroup)
	group.PATCH("", ctrl.UpdateGroup)
	group.DELETE("", ctrl.DeleteGroup)

	apps := group.Group("/apps")
	apps.GET("", ctrl.GetApps)
	apps.POST("", ctrl.CreateApp)

	app := group.Group("/apps/:app_name", ctrl.SetApp)
	app.GET("", ctrl.GetApp)
	app.PATCH("", ctrl.UpdateApp)
	app.DELETE("", ctrl.DeleteApp)

	configSets := app.Group("/configsets")
	configSets.GET("", ctrl.GetConfigSets)
	configSets.POST("", ctrl.CreateConfigSet)

	configSet := app.Group("/configsets/:configset_id", ctrl.SetConfigSet)
	configSet.GET("", ctrl.GetConfigSet)
	configSet.PATCH("", ctrl.UpdateConfigSet)
	configSet.PUT("/content", ctrl.UpdateConfigContent)
	configSet.DELETE("", ctrl.DeleteConfigSet)

	publishes := configSet.Group("")
	publishes.POST("/publish", ctrl.PublishConfigSet)
	publishes.POST("/revert", ctrl.RevertConfigSet)
	publishes.GET("/publishes", ctrl.GetPublishes)

	publish := publishes.Group("/publishes/:publish_id", ctrl.SetPublish)
	publish.GET("", ctrl.GetPublish)

	instances := configSet.Group("/instances")
	instances.GET("", ctrl.GetInstances)

	return router
}
