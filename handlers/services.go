package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetServicesArgs struct {
	NamespaceName string   `uri:"namespace_name"`
	Q             string   `query:"q"`
	Names         []string `query:"name"`
}

func (ctrl *Handler) GetServices(ctx *fox.Context, args *GetServicesArgs) (*models.Pagination[*models.Service], error) {
	var (
		logger    = ctx.Logger
		namespace = ctx.MustGet("namespace").(*models.Namespace)
		err       error
	)

	logger.Debugf("get services args: %+v", args)

	var params = &params.GetServices{
		Q:           args.Q,
		NamespaceID: namespace.ID,
	}

	pagination, err := ctrl.manager.GetServices(ctx, params)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

type GetServiceArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `uri:"service_name"`
}

func (ctrl *Handler) GetService(ctx *fox.Context, args *GetServiceArgs) (*models.Service, error) {

	var (
		namespace = ctx.MustGet("namespace").(*models.Namespace)
		service   = ctx.MustGet("service").(*models.Service)
	)

	service.Namespace = namespace

	return service, nil
}

type UpdateServiceArgs struct {
	Description *string `json:"description"`
}

func (ctrl *Handler) UpdateService(ctx *fox.Context, args *UpdateServiceArgs) error {
	var (
		logger  = ctx.Logger
		user    = ctrl.CurrentUser(ctx)
		service = ctx.MustGet("service").(*models.Service)
	)

	logger.Debugf("update service args: %+v", args)

	err := ctrl.manager.UpdateService(ctx, service.ID, &params.UpdateService{
		Description: args.Description,
		UpdatedBy:   user.Username,
	})

	if err != nil {
		return err
	}

	return nil
}

type CreateServiceArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

func (ctrl *Handler) CreateService(ctx *fox.Context, args *CreateServiceArgs) (*models.Service, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		namespace = ctx.MustGet("namespace").(*models.Namespace)
	)

	logger.Debugf("create service args: %+v", args)

	service, err := ctrl.manager.CreateService(ctx, &params.CreateService{
		NamespaceID: namespace.ID,
		Name:        args.Name,
		Description: args.Description,
		CreatedBy:   user.Username,
	})

	if err != nil {
		return nil, err
	}

	return service, nil
}

type DeleteServiceArgs struct {
	NamespaceName string `uri:"namespace_name"`
	Name          string `uri:"service_name"`
}

func (ctrl *Handler) DeleteService(ctx *fox.Context, args *DeleteServiceArgs) error {
	var (
		logger  = ctx.Logger
		service = ctx.MustGet("service").(*models.Service)
	)

	logger.Debugf("delete service args: %+v", args)

	err := ctrl.manager.DeleteService(ctx, service.ID)
	if err != nil {
		logger.Errorf("delete service error: %+v", err)
		return err
	}

	return nil
}

func (ctrl *Handler) SetService(ctx *fox.Context, args *GetServiceArgs) (res any) {

	var namespace = ctx.MustGet("namespace").(*models.Namespace)

	service, err := ctrl.manager.GetService(ctx, &params.GetService{
		NamespaceID: namespace.ID,
		Name:        args.Name,
	})
	if err != nil {
		return err
	}

	ctx.Set("service", service)

	return
}
