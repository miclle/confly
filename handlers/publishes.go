package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type PublishConfigSetArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigsetID string `uri:"configset_id"`
	FullRelease bool   `query:"fullRelease"`

	Title       string           `json:"title"`
	Description string           `json:"description"`
	GrayRule    *models.GrayRule `json:"grayRule"`

	MergedConfigChecksum string `json:"mergedConfigChecksum"`
}

func (ctrl *Handler) PublishConfigSet(ctx *fox.Context, args *PublishConfigSetArgs) (*models.Publish, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)

	logger.Debugf("publish configset, args: %+v", args)

	publishType := models.PublishTypeGrayRelease
	if args.FullRelease {
		publishType = models.PublishTypeFullRelease
	}

	publishConfigSetParams := &params.PublishConfigSet{
		ConfigSetID: configSet.ID,
		Type:        publishType,
		Title:       args.Title,
		Description: args.Description,
		GrayRule:    args.GrayRule,
		CreatedBy:   user.Username,
	}

	publish, err := ctrl.manager.PublishConfigSet(ctx, publishConfigSetParams)
	if err != nil {
		return nil, err
	}

	return publish, nil
}

type RevertConfigSetArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigsetID string `uri:"configset_id"`
	ClusterName string `uri:"cluster_name"`
}

func (ctrl *Handler) RevertConfigSet(ctx *fox.Context, args *RevertConfigSetArgs) (*models.Publish, error) {
	var (
		logger    = ctx.Logger
		user      = ctrl.CurrentUser(ctx)
		configSet = ctx.MustGet("configSet").(*models.ConfigSet)
	)

	logger.Debugf("revert configset, args: %+v", args)

	params := &params.RevertConfigSet{
		GroupName:     args.GroupName,
		AppName:       args.AppName,
		ClusterName:   args.ClusterName,
		ConfigSetName: configSet.Name,
		ConfigSetID:   configSet.ID,
		CreatedBy:     user.Username,
	}

	publish, err := ctrl.manager.RevertConfigSet(ctx, params)
	if err != nil {
		return nil, err
	}

	return publish, nil
}

type GetPublishesArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigsetID string `uri:"configset_id"`

	Q           string `query:"q"`
	IncludeGray bool   `query:"includeGray"`
	Page        int    `query:"page"`
	Size        int    `query:"size"`
}

func (ctrl *Handler) GetPublishes(ctx *fox.Context, args *GetPublishesArgs) (*models.Pagination[*models.Publish], error) {
	var (
		logger = ctx.Logger

		configSet  = ctx.MustGet("configSet").(*models.ConfigSet)
		err        error
		pagination *models.Pagination[*models.Publish]
	)

	logger.Debugf("get publishes, args: %+v", args)

	p := &params.GetPublishes{
		ConfigSetID: configSet.ID,
		Q:           args.Q,

		Pagination: models.Pagination[*models.Publish]{
			Page: args.Page,
			Size: args.Size,
		},
	}

	pagination, err = ctrl.manager.GetPublishes(ctx, p)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

type GetPublishArgs struct {
	GroupName   string `uri:"group_name"`
	AppName     string `uri:"app_name"`
	ConfigsetID string `uri:"configset_id"`
	PublishID   string `uri:"publish_id"`
}

func (ctrl *Handler) GetPublish(ctx *fox.Context, args *GetPublishArgs) (*models.Publish, error) {
	var (
		logger = ctx.Logger

		publish = ctx.MustGet("publish").(*models.Publish)
	)

	logger.Debugf("get publish, args: %+v", args)

	return publish, nil
}

func (ctrl *Handler) SetPublish(ctx *fox.Context, args *GetPublishArgs) (res any) {

	publish, err := ctrl.manager.GetPublish(ctx, &params.GetPublish{
		PublishID: args.PublishID,
	})

	if err != nil {
		ctx.Logger.Errorf("get publish failed, err: %+v", err)
		return err
	}

	ctx.Set("publish", publish)

	return
}
