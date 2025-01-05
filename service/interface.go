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

	// group
	CreateGroup(ctx context.Context, params *params.CreateGroup) (*models.Group, error)
	GetGroups(ctx context.Context, params *params.GetGroups) ([]*models.Group, error)
	GetGroup(ctx context.Context, params *params.GetGroup) (*models.Group, error)
	UpdateGroup(ctx context.Context, GroupID string, params *params.UpdateGroup) error
	DeleteGroup(ctx context.Context, GroupID string) error

	// app
	CreateApp(ctx context.Context, params *params.CreateApp) (*models.App, error)
	GetApps(ctx context.Context, params *params.GetApps) ([]*models.App, error)
	GetApp(ctx context.Context, params *params.GetApp) (*models.App, error)
	UpdateApp(ctx context.Context, appID string, params *params.UpdateApp) error
	DeleteApp(ctx context.Context, appID string) error

	// configSet
	CreateConfigSet(ctx context.Context, params *params.CreateConfigSet) (*models.ConfigSet, error)
	GetConfigSets(ctx context.Context, params *params.GetConfigSets) (*models.Pagination[*models.ConfigSet], error)
	GetConfigSet(ctx context.Context, configSetID string) (*models.ConfigSet, error)
	UpdateConfigSet(ctx context.Context, configSetID string, params *params.UpdateConfigSet) (*models.ConfigSet, error)
	UpdateConfigContent(ctx context.Context, cconfigSetID string, params *params.UpdateConfigContent) (*models.ConfigSet, error)
	DeleteConfigSet(ctx context.Context, configSetID string) error

	// publish
	PublishConfigSet(ctx context.Context, params *params.PublishConfigSet) (*models.Publish, error)
	RevertConfigSet(ctx context.Context, params *params.RevertConfigSet) (*models.Publish, error)
	RollbackConfigSet(ctx context.Context, params *params.RollbackConfigSet) (*models.Publish, error)
	GetPublishes(ctx context.Context, params *params.GetPublishes) (*models.Pagination[*models.Publish], error)
	GetPublish(ctx context.Context, params *params.GetPublish) (*models.Publish, error)

	// instance
	UpsertInstance(ctx context.Context, params *params.UpsertInstance) (*models.Instance, error)
	GetInstances(ctx context.Context, params *params.GetInstances) (*models.Pagination[*models.Instance], error)
}
