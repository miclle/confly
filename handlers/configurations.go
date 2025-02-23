package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetConfigurationsArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	Q               string `query:"q"`
	IsDefault       *bool  `query:"isDefault"`
	Page            int    `query:"page"`
	Size            int    `query:"size"`
}

func (ctrl *Handler) GetConfigurations(ctx *fox.Context, args *GetConfigurationsArgs) (*models.Pagination[*models.Configuration], error) {

	var (
		logger      = ctx.Logger
		application = ctx.MustGet("application").(*models.Application)
	)

	logger.Debugf("get config sets args: %+v", args)

	p := &params.GetConfigurations{
		ApplicationID: application.ID,
		Q:             args.Q,
		Pagination: models.Pagination[*models.Configuration]{
			Page: args.Page,
			Size: args.Size,
		},
	}

	pagination, err := ctrl.manager.GetConfigurations(ctx, p)
	if err != nil {
		logger.Errorf("get config sets error: %+v", err)
		return nil, err
	}

	return pagination, nil
}

type GetConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`
}

func (ctrl *Handler) GetConfiguration(ctx *fox.Context, args *GetConfigurationArgs) (*models.Configuration, error) {

	var (
		logger        = ctx.Logger
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("get config set args: %+v", args)

	return configuration, nil
}

type CreateConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ConfigFormat    models.ConfigurationFormat
}

func (ctrl *Handler) CreateConfiguration(ctx *fox.Context, args *CreateConfigurationArgs) (*models.Configuration, error) {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
		app    = ctx.MustGet("application").(*models.Application)
	)

	logger.Debugf("create config set args: %+v", args)

	p := &params.CreateConfiguration{
		ApplicationID: app.ID,
		Name:          args.Name,
		Description:   args.Description,
		ConfigFormat:  args.ConfigFormat,
		CreatedBy:     user.Username,
	}

	configuration, err := ctrl.manager.CreateConfiguration(ctx, p)
	if err != nil {
		logger.Errorf("create config set error: %+v", err)
		return nil, err
	}

	return configuration, nil
}

type UpdateConfigurationArgs struct {
	ExtendDefault *bool   `json:"extendDefault"`
	Description   *string `json:"description"`
}

func (ctrl *Handler) UpdateConfiguration(ctx *fox.Context, args *UpdateConfigurationArgs) (*models.Configuration, error) {
	var (
		logger        = ctx.Logger
		user          = ctrl.CurrentUser(ctx)
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("update config set args: %+v", args)

	params := &params.UpdateConfiguration{
		Description: args.Description,
		UpdatedBy:   user.Username,
	}

	configuration, err := ctrl.manager.UpdateConfiguration(ctx, configuration.ID, params)
	if err != nil {
		logger.Errorf("update config set error %+v", err)
		return nil, err
	}

	return configuration, nil
}

type UpdateConfigurationContentArgs struct {
	ConfigContent         string `json:"configContent"`
	CurrentConfigChecksum string `json:"currentConfigChecksum"`
}

func (ctrl *Handler) UpdateConfigurationContent(ctx *fox.Context, args *UpdateConfigurationContentArgs) (*models.Configuration, error) {
	var (
		logger        = ctx.Logger
		user          = ctrl.CurrentUser(ctx)
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("update config content args: %+v", args)

	params := &params.UpdateConfigurationContent{
		Content:   args.ConfigContent,
		Checksum:  args.CurrentConfigChecksum,
		UpdatedBy: user.Username,
	}

	configuration, err := ctrl.manager.UpdateConfigurationContent(ctx, configuration.ID, params)
	if err != nil {
		logger.Errorf("update config content error %+v", err)
		return nil, err
	}

	return configuration, nil
}

type DeleteConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`
}

func (ctrl *Handler) DeleteConfiguration(ctx *fox.Context, args *DeleteConfigurationArgs) error {

	var (
		logger        = ctx.Logger
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("delete config set args: %+v", args)

	err := ctrl.manager.DeleteConfiguration(ctx, configuration.ID)
	return err
}

func (ctrl *Handler) SetConfiguration(ctx *fox.Context, args *GetConfigurationArgs) (res any) {

	configuration, err := ctrl.manager.GetConfiguration(ctx, args.ConfigurationID)
	if err != nil {
		ctx.Logger.Errorf("get configuration failed, err: %+v", err)
		return err
	}

	ctx.Set("configuration", configuration)

	return
}
