package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetConfigSetsArgs struct {
	GroupName string   `uri:"group_name"`
	AppName   string   `uri:"app_name"`
	Q         string   `query:"q"`
	IsDefault *bool    `query:"isDefault"`
	Names     []string `query:"name"`
	ClusterID string   `query:"clusterID"`
	Page      int      `query:"page"`
	Size      int      `query:"size"`
}

func (ctrl *Handler) GetConfigSets(ctx *fox.Context, args *GetConfigSetsArgs) (*models.Pagination[*models.ConfigSet], error) {

	var (
		logger = ctx.Logger
		app    = ctx.MustGet("app").(*models.App)
	)

	logger.Debugf("get config sets args: %+v", args)

	p := &params.GetConfigSets{
		AppID: app.ID,
		Q:     args.Q,
		Pagination: models.Pagination[*models.ConfigSet]{
			Page: args.Page,
			Size: args.Size,
		},
	}

	pagination, err := ctrl.manager.GetConfigSets(ctx, p)
	if err != nil {
		logger.Errorf("get config sets error: %+v", err)
		return nil, err
	}

	return pagination, nil
}

type GetConfigSetArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigSetID string `uri:"configset_id"`
}

func (ctrl *Handler) GetConfigSet(ctx *fox.Context, args *GetConfigSetArgs) (*models.ConfigSet, error) {

	var (
		logger    = ctx.Logger
		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)

	logger.Debugf("get config set args: %+v", args)

	return configSet, nil
}

type CreateConfigSetArgs struct {
	GroupName    string `uri:"group_name"`
	AppName      string `uri:"app_name"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ConfigFormat models.ConfigFormat
}

func (ctrl *Handler) CreateConfigSet(ctx *fox.Context, args *CreateConfigSetArgs) (*models.ConfigSet, error) {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
		app    = ctx.MustGet("app").(*models.App)
	)

	logger.Debugf("create config set args: %+v", args)

	p := &params.CreateConfigSet{
		AppID:        app.ID,
		Name:         args.Name,
		Description:  args.Description,
		ConfigFormat: args.ConfigFormat,
		CreatedBy:    user.Username,
	}

	configSets, err := ctrl.manager.CreateConfigSet(ctx, p)
	if err != nil {
		logger.Errorf("create config set error: %+v", err)
		return nil, err
	}

	return configSets, nil
}

type UpdateConfigSetArgs struct {
	ExtendDefault *bool   `json:"extendDefault"`
	Description   *string `json:"description"`
}

func (ctrl *Handler) UpdateConfigSet(ctx *fox.Context, args *UpdateConfigSetArgs) (*models.ConfigSet, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)

	logger.Debugf("update config set args: %+v", args)

	params := &params.UpdateConfigSet{
		Description: args.Description,
		UpdatedBy:   user.Username,
	}

	configSet, err := ctrl.manager.UpdateConfigSet(ctx, configSet.ID, params)
	if err != nil {
		logger.Errorf("update config set error %+v", err)
		return nil, err
	}

	return configSet, nil
}

type UpdateConfigContentArgs struct {
	ConfigContent         string `json:"configContent"`
	CurrentConfigChecksum string `json:"currentConfigChecksum"`
}

func (ctrl *Handler) UpdateConfigContent(ctx *fox.Context, args *UpdateConfigContentArgs) (*models.ConfigSet, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)

	logger.Debugf("update config content args: %+v", args)

	params := &params.UpdateConfigContent{
		ConfigContent:  args.ConfigContent,
		ConfigChecksum: args.CurrentConfigChecksum,
		UpdatedBy:      user.Username,
	}

	configSet, err := ctrl.manager.UpdateConfigContent(ctx, configSet.ID, params)
	if err != nil {
		logger.Errorf("update config content error %+v", err)
		return nil, err
	}

	return configSet, nil
}

type DeleteConfigSetArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigSetID string `uri:"configset_id"`
}

func (ctrl *Handler) DeleteConfigSet(ctx *fox.Context, args *DeleteConfigSetArgs) error {

	var (
		logger    = ctx.Logger
		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)

	logger.Debugf("delete config set args: %+v", args)

	err := ctrl.manager.DeleteConfigSet(ctx, configSet.ID)
	return err
}

func (ctrl *Handler) SetConfigSet(ctx *fox.Context, args *GetConfigSetArgs) (res any) {

	configSet, err := ctrl.manager.GetConfigSet(ctx, args.ConfigSetID)
	if err != nil {
		ctx.Logger.Errorf("get configset failed, err: %+v", err)
		return err
	}

	ctx.Set("configSet", configSet)

	return
}
