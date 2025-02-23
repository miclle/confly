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

	router.Use(ctrl.platformAuthMiddleware)

	router.GET("/confly/v1/signin", ctrl.Signin)
	router.GET("/confly/v1/logout", ctrl.Logout)
	router.GET("/confly/v1/boot", ctrl.Boot)

	system := router.Group("/confly/v1/system")
	system.GET("/users", ctrl.GetUsers)
	system.PATCH("/users/:username", ctrl.UpdateUser)

	namespaces := router.Group("/confly/v1/namespaces")
	namespaces.GET("", ctrl.GetNamespaces)
	namespaces.POST("", ctrl.CreateNamespace)

	namespace := router.Group("/confly/v1/namespaces/:namespace_name", ctrl.SetNamespace)
	namespace.GET("", ctrl.GetNamespace)
	namespace.PATCH("", ctrl.UpdateNamespace)
	namespace.DELETE("", ctrl.DeleteNamespace)

	apps := namespace.Group("/apps")
	apps.GET("", ctrl.GetApplications)
	apps.POST("", ctrl.CreateApplication)

	app := namespace.Group("/apps/:application_name", ctrl.SetApplication)
	app.GET("", ctrl.GetApplication)
	app.PATCH("", ctrl.UpdateApplication)
	app.DELETE("", ctrl.DeleteApplication)

	configurations := app.Group("/configurations")
	configurations.GET("", ctrl.GetConfigurations)
	configurations.POST("", ctrl.CreateConfiguration)

	configuration := app.Group("/configurations/:configuration_id", ctrl.SetConfiguration)
	configuration.GET("", ctrl.GetConfiguration)
	configuration.PATCH("", ctrl.UpdateConfiguration)
	configuration.PUT("/content", ctrl.UpdateConfigurationContent)
	configuration.DELETE("", ctrl.DeleteConfiguration)

	publishes := configuration.Group("")
	publishes.POST("/publish", ctrl.PublishConfiguration)
	publishes.POST("/revert", ctrl.RevertConfiguration)
	publishes.GET("/publishes", ctrl.GetPublishes)

	publish := publishes.Group("/publishes/:version_id", ctrl.SetPublish)
	publish.GET("", ctrl.GetPublish)

	instances := configuration.Group("/instances")
	instances.GET("", ctrl.GetInstances)

	return router
}
