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

	services := namespace.Group("/services")
	services.GET("", ctrl.GetServices)
	services.POST("", ctrl.CreateService)

	service := namespace.Group("/services/:service_name", ctrl.SetService)
	service.GET("", ctrl.GetService)
	service.PATCH("", ctrl.UpdateService)
	service.DELETE("", ctrl.DeleteService)

	applications := namespace.Group("/applications")
	applications.GET("", ctrl.GetApplications)
	applications.POST("", ctrl.CreateApplication)

	application := namespace.Group("/applications/:application_name", ctrl.SetApplication)
	application.GET("", ctrl.GetApplication)
	application.PATCH("", ctrl.UpdateApplication)
	application.DELETE("", ctrl.DeleteApplication)

	configurations := application.Group("/configurations")
	configurations.GET("", ctrl.GetConfigurations)
	configurations.POST("", ctrl.CreateConfiguration)

	configuration := application.Group("/configurations/:configuration_id", ctrl.SetConfiguration)
	configuration.GET("", ctrl.GetConfiguration)
	configuration.PATCH("", ctrl.UpdateConfiguration)
	configuration.PUT("/content", ctrl.UpdateConfigurationContent)
	configuration.DELETE("", ctrl.DeleteConfiguration)

	releases := configuration.Group("")
	releases.POST("/publish", ctrl.PublishConfiguration)
	releases.POST("/revert", ctrl.RevertConfiguration)
	releases.GET("/releases", ctrl.GetReleases)

	release := releases.Group("/releases/:version_id", ctrl.SetRelease)
	release.GET("", ctrl.GetRelease)

	instances := configuration.Group("/instances")
	instances.GET("", ctrl.GetInstances)

	return router
}
