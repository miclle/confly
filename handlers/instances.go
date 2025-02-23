package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetInstancesArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`
	Environment     string `uri:"environment"`

	Page int `query:"page"`
	Size int `query:"size"`
}

func (ctrl *Handler) GetInstances(ctx *fox.Context, args *GetInstancesArgs) (*models.Pagination[*models.Instance], error) {
	var (
		logger = ctx.Logger

		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)
	logger.Debugf("get instances, args: %+v", args)

	params := &params.GetInstances{
		NamespaceName:     args.NamespaceName,
		ApplicationName:   args.ApplicationName,
		ConfigurationName: configuration.Name,
		Environment:       args.Environment,
		Pagination: models.Pagination[*models.Instance]{
			Page: args.Page,
			Size: args.Size,
		},
	}

	pagination, err := ctrl.manager.GetInstances(ctx, params)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}
