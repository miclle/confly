package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetAppsArgs struct {
	GroupName  string   `uri:"group_name"`
	ClusterIDs []string `query:"clusterID"`
	Q          string   `query:"q"`
	Names      []string `query:"name"`
}

func (ctrl *Handler) GetApps(ctx *fox.Context, args *GetAppsArgs) (*models.Pagination[*models.App], error) {
	var (
		logger = ctx.Logger
		group  = ctx.MustGet("group").(*models.Group)
		err    error
	)

	logger.Debugf("get apps args: %+v", args)

	var params = &params.GetApps{
		Q:       args.Q,
		GroupID: group.ID,
	}

	pagination, err := ctrl.manager.GetApps(ctx, params)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

type GetAppArgs struct {
	GroupName string `uri:"group_name"`
	Name      string `uri:"app_name"`
}

func (ctrl *Handler) GetApp(ctx *fox.Context, args *GetAppArgs) (*models.App, error) {

	var (
		group = ctx.MustGet("group").(*models.Group)
		app   = ctx.MustGet("app").(*models.App)
	)

	app.Group = group

	return app, nil
}

type UpdateAppArgs struct {
	Description *string `json:"description"`
}

func (ctrl *Handler) UpdateApp(ctx *fox.Context, args *UpdateAppArgs) error {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
		app    = ctx.MustGet("app").(*models.App)
	)

	logger.Debugf("update app args: %+v", args)

	err := ctrl.manager.UpdateApp(ctx, app.ID, &params.UpdateApp{
		Description: args.Description,
		UpdatedBy:   user.Username,
	})

	if err != nil {
		return err
	}

	return nil
}

type CreateAppArgs struct {
	GroupName   string `uri:"group_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (ctrl *Handler) CreateApp(ctx *fox.Context, args *CreateAppArgs) (*models.App, error) {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
		group  = ctx.MustGet("group").(*models.Group)
	)

	logger.Debugf("create app args: %+v", args)

	app, err := ctrl.manager.CreateApp(ctx, &params.CreateApp{
		GroupID:     group.ID,
		Name:        args.Name,
		Description: args.Description,
		CreatedBy:   user.Username,
	})

	if err != nil {
		return nil, err
	}

	return app, nil
}

type DeleteAppArgs struct {
	GroupName string `uri:"group_name"`
	Name      string `uri:"app_name"`
}

func (ctrl *Handler) DeleteApp(ctx *fox.Context, args *DeleteAppArgs) error {
	var (
		logger = ctx.Logger
		app    = ctx.MustGet("app").(*models.App)
	)

	logger.Debugf("delete app args: %+v", args)

	err := ctrl.manager.DeleteApp(ctx, app.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *Handler) SetApp(ctx *fox.Context, args *GetAppArgs) (res any) {

	var group = ctx.MustGet("group").(*models.Group)

	app, err := ctrl.manager.GetApp(ctx, &params.GetApp{
		GroupID: group.ID,
		Name:    args.Name,
	})
	if err != nil {
		return err
	}

	ctx.Set("app", app)

	return
}
