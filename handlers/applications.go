package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetApplicationsArgs struct {
	NamespaceName string   `uri:"namespace_name"`
	Q             string   `query:"q"`
	Names         []string `query:"name"`
}

func (ctrl *Handler) GetApplications(ctx *fox.Context, args *GetApplicationsArgs) (*models.Pagination[*models.Application], error) {
	var (
		logger    = ctx.Logger
		namespace = ctx.MustGet("namespace").(*models.Namespace)
		err       error
	)

	logger.Debugf("get applications args: %+v", args)

	var params = &params.GetApplications{
		Q:           args.Q,
		NamespaceID: namespace.ID,
	}

	pagination, err := ctrl.manager.GetApplications(ctx, params)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

type GetApplicationArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `uri:"application_name"`
}

func (ctrl *Handler) GetApplication(ctx *fox.Context, args *GetApplicationArgs) (*models.Application, error) {

	var (
		namespace   = ctx.MustGet("namespace").(*models.Namespace)
		application = ctx.MustGet("application").(*models.Application)
	)

	application.Namespace = namespace

	return application, nil
}

type UpdateApplicationArgs struct {
	Description *string `json:"description"`
}

func (ctrl *Handler) UpdateApplication(ctx *fox.Context, args *UpdateApplicationArgs) error {
	var (
		logger      = ctx.Logger
		user        = ctrl.CurrentUser(ctx)
		application = ctx.MustGet("application").(*models.Application)
	)

	logger.Debugf("update application args: %+v", args)

	err := ctrl.manager.UpdateApplication(ctx, application.ID, &params.UpdateApplication{
		Description: args.Description,
		UpdatedBy:   user.Username,
	})

	if err != nil {
		return err
	}

	return nil
}

type CreateApplicationArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

func (ctrl *Handler) CreateApplication(ctx *fox.Context, args *CreateApplicationArgs) (*models.Application, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		namespace = ctx.MustGet("namespace").(*models.Namespace)
	)

	logger.Debugf("create application args: %+v", args)

	application, err := ctrl.manager.CreateApplication(ctx, &params.CreateApplication{
		NamespaceID: namespace.ID,
		Name:        args.Name,
		Description: args.Description,
		CreatedBy:   user.Username,
	})

	if err != nil {
		return nil, err
	}

	return application, nil
}

type DeleteApplicationArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `uri:"application_name"`
}

func (ctrl *Handler) DeleteApplication(ctx *fox.Context, args *DeleteApplicationArgs) error {
	var (
		logger      = ctx.Logger
		application = ctx.MustGet("application").(*models.Application)
	)

	logger.Debugf("delete application args: %+v", args)

	err := ctrl.manager.DeleteApplication(ctx, application.ID)
	if err != nil {
		logger.Errorf("delete application error: %+v", err)
		return err
	}

	return nil
}

func (ctrl *Handler) SetApplication(ctx *fox.Context, args *GetApplicationArgs) (res any) {

	var namespace = ctx.MustGet("namespace").(*models.Namespace)

	application, err := ctrl.manager.GetApplication(ctx, &params.GetApplication{
		NamespaceID: namespace.ID,
		Name:        args.Name,
	})
	if err != nil {
		return err
	}

	ctx.Set("application", application)

	return
}
