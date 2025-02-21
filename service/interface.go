package service

import (
	"context"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type Service interface {
	// user
	UpsertUser(ctx context.Context, username, email string) (*models.User, error)
	GetUsers(ctx context.Context, params *params.GetUsers) (*models.Pagination[*models.User], error)
	GetUser(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, username string, params *params.UpdateUser) (*models.User, error)

	// namespaces
	CreateNamespace(ctx context.Context, params *params.CreateNamespace) (*models.Namespace, error)
	GetNamespaces(ctx context.Context, params *params.GetNamespaces) (*models.Pagination[*models.Namespace], error)
	GetNamespace(ctx context.Context, params *params.GetNamespace) (*models.Namespace, error)
	UpdateNamespace(ctx context.Context, id string, params *params.UpdateNamespace) error
	DeleteNamespace(ctx context.Context, id string) error

	// applications
	CreateApplication(ctx context.Context, params *params.CreateApp) (*models.Application, error)
	GetApplications(ctx context.Context, params *params.GetApps) (*models.Pagination[*models.Application], error)
	GetApplication(ctx context.Context, params *params.GetApp) (*models.Application, error)
	UpdateApplication(ctx context.Context, appID string, params *params.UpdateApp) error
	DeleteApplication(ctx context.Context, appID string) error

	// configurations
	CreateConfiguration(ctx context.Context, params *params.CreateConfiguration) (*models.Configuration, error)
	GetConfigurations(ctx context.Context, params *params.GetConfigurations) (*models.Pagination[*models.Configuration], error)
	GetConfiguration(ctx context.Context, id string) (*models.Configuration, error)
	UpdateConfiguration(ctx context.Context, id string, params *params.UpdateConfiguration) (*models.Configuration, error)
	UpdateConfigurationContent(ctx context.Context, id string, params *params.UpdateConfigurationContent) (*models.Configuration, error)
	DeleteConfiguration(ctx context.Context, id string) error

	// publishes
	PublishConfiguration(ctx context.Context, params *params.PublishConfiguration) (*models.Publish, error)
	RevertConfiguration(ctx context.Context, params *params.RevertConfiguration) (*models.Publish, error)
	RollbackConfiguration(ctx context.Context, params *params.RollbackConfiguration) (*models.Publish, error)
	GetPublishes(ctx context.Context, params *params.GetPublishes) (*models.Pagination[*models.Publish], error)
	GetPublish(ctx context.Context, params *params.GetPublish) (*models.Publish, error)

	// instances
	UpsertInstance(ctx context.Context, params *params.UpsertInstance) (*models.Instance, error)
	GetInstances(ctx context.Context, params *params.GetInstances) (*models.Pagination[*models.Instance], error)
}
