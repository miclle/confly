package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetNamespacesArgs struct {
	Q string `query:"q"`
}

func (ctrl *Handler) GetNamespaces(ctx *fox.Context, args *GetNamespacesArgs) (*models.Pagination[*models.Namespace], error) {
	logger := ctx.Logger
	logger.Debugf("get namespaces args: %+v", args)

	pagination, err := ctrl.manager.GetNamespaces(ctx, &params.GetNamespaces{
		Q: args.Q,
	})

	if err != nil {
		logger.Errorf("get namespaces error: %+v", err)
		return nil, err
	}

	return pagination, nil
}

func (ctrl *Handler) GetNamespace(ctx *fox.Context, args *GetNamespaceArgs) (*models.Namespace, error) {

	var namespace = ctx.MustGet("namespace").(*models.Namespace)

	return namespace, nil
}

type UpdateNamespaceArgs struct {
	Description *string `json:"description"`
}

func (ctrl *Handler) UpdateNamespace(ctx *fox.Context, args *UpdateNamespaceArgs) error {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		namespace = ctx.MustGet("namespace").(*models.Namespace)
	)

	logger.Debugf("update namespace args: %+v", args)

	err := ctrl.manager.UpdateNamespace(ctx, namespace.ID, &params.UpdateNamespace{
		Description: args.Description,
		UpdatedBy:   user.Username,
	})

	if err != nil {
		logger.Errorf("update namespace error: %+v", err)
		return err
	}
	return nil
}

type CreateNamespaceArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (ctrl *Handler) CreateNamespace(ctx *fox.Context, args *CreateNamespaceArgs) (*models.Namespace, error) {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
	)

	logger.Debugf("create namespace args: %+v", args)

	params := &params.CreateNamespace{
		Name:        args.Name,
		Description: args.Description,
		CreatedBy:   user.Username,
	}

	namespace, err := ctrl.manager.CreateNamespace(ctx, params)
	if err != nil {
		logger.Errorf("create namespace error: %+v", err)
		return nil, err
	}

	return namespace, nil
}

type DeleteNamespaceArgs struct {
	Name string `uri:"namespace_name"`
}

func (ctrl *Handler) DeleteNamespace(ctx *fox.Context, args *DeleteNamespaceArgs) error {
	var (
		logger    = ctx.Logger
		namespace = ctx.MustGet("namespace").(*models.Namespace)
	)

	logger.Debugf("delete namespace arg: %+v", namespace.Name)

	err := ctrl.manager.DeleteNamespace(ctx, namespace.ID)
	if err != nil {
		logger.Errorf("delete namespace error: %+v", err)
		return err
	}

	return nil
}

type GetNamespaceArgs struct {
	Name string `uri:"namespace_name"`
}

func (ctrl *Handler) SetNamespace(ctx *fox.Context, args *GetNamespaceArgs) (res any) {

	namespace, err := ctrl.manager.GetNamespace(ctx, &params.GetNamespace{
		Name: args.Name,
	})
	if err != nil {
		ctx.Logger.Errorf("get namespace error: %+v", err)
		return err
	}

	ctx.Set("namespace", namespace)

	return
}
