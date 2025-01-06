package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetInstancesArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigsetID string `uri:"configset_id"`
	ClusterName string `uri:"cluster_name"`

	Page int `query:"page"`
	Size int `query:"size"`
}

func (ctrl *Handler) GetInstances(ctx *fox.Context, args *GetInstancesArgs) (*models.Pagination[*models.Instance], error) {
	var (
		logger = ctx.Logger

		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)
	logger.Debugf("get instances, args: %+v", args)

	params := &params.GetInstances{
		GroupName:     args.GroupName,
		AppName:       args.AppName,
		ConfigSetName: configSet.Name,
		ClusterName:   args.ClusterName,
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
