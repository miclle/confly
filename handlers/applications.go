package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetAppsArgs struct {
	NamespaceName string   `uri:"namespace_name"`
	ClusterIDs    []string `query:"clusterID"`
	Q             string   `query:"q"`
	Names         []string `query:"name"`
}

func (ctrl *Handler) GetApps(ctx *fox.Context, args *GetAppsArgs) (*models.Pagination[*models.Application], error) {
	var (
		logger    = ctx.Logger
		namespace = ctx.MustGet("namespace").(*models.Namespace)
		err       error
	)

	logger.Debugf("get apps args: %+v", args)

	var params = &params.GetApps{
		Q:           args.Q,
		NamespaceID: namespace.ID,
	}

	pagination, err := ctrl.manager.GetApplications(ctx, params)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

type GetAppArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `uri:"app_name"`
}

func (ctrl *Handler) GetApp(ctx *fox.Context, args *GetAppArgs) (*models.Application, error) {

	var (
		namespace = ctx.MustGet("namespace").(*models.Namespace)
		app       = ctx.MustGet("app").(*models.Application)
	)

	app.Namespace = namespace

	return app, nil
}

type UpdateAppArgs struct {
	Description *string `json:"description"`
}

func (ctrl *Handler) UpdateApp(ctx *fox.Context, args *UpdateAppArgs) error {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
		app    = ctx.MustGet("app").(*models.Application)
	)

	logger.Debugf("update app args: %+v", args)

	err := ctrl.manager.UpdateApplication(ctx, app.ID, &params.UpdateApp{
		Description: args.Description,
		UpdatedBy:   user.Username,
	})

	if err != nil {
		return err
	}

	return nil
}

type CreateAppArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

func (ctrl *Handler) CreateApp(ctx *fox.Context, args *CreateAppArgs) (*models.Application, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		namespace = ctx.MustGet("namespace").(*models.Namespace)
	)

	logger.Debugf("create app args: %+v", args)

	app, err := ctrl.manager.CreateApplication(ctx, &params.CreateApp{
		NamespaceID: namespace.ID,
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
	NamespaceName string `uri:"namespace_name"`
	Name          string `uri:"app_name"`
}

func (ctrl *Handler) DeleteApp(ctx *fox.Context, args *DeleteAppArgs) error {
	var (
		logger = ctx.Logger
		app    = ctx.MustGet("app").(*models.Application)
	)

	logger.Debugf("delete app args: %+v", args)

	err := ctrl.manager.DeleteApplication(ctx, app.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *Handler) SetApp(ctx *fox.Context, args *GetAppArgs) (res any) {

	var namespace = ctx.MustGet("namespace").(*models.Namespace)

	app, err := ctrl.manager.GetApplication(ctx, &params.GetApp{
		NamespaceID: namespace.ID,
		Name:        args.Name,
	})
	if err != nil {
		return err
	}

	ctx.Set("app", app)

	return
}
