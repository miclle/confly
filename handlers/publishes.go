package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type PublishConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	AppName         string `uri:"app_name"`
	ConfigurationID string `uri:"configuration_id"`
	FullRelease     bool   `query:"fullRelease"`

	Title       string           `json:"title"`
	Description string           `json:"description"`
	GrayRule    *models.GrayRule `json:"grayRule"`

	MergedConfigChecksum string `json:"mergedConfigChecksum"`
}

func (ctrl *Handler) PublishConfiguration(ctx *fox.Context, args *PublishConfigurationArgs) (*models.Publish, error) {
	var (
		logger        = ctx.Logger
		user          = ctrl.CurrentUser(ctx)
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("publish configuration, args: %+v", args)

	publishType := models.PublishTypeGrayRelease
	if args.FullRelease {
		publishType = models.PublishTypeFullRelease
	}

	publishConfigurationParams := &params.PublishConfiguration{
		ConfigurationID: configuration.ID,
		Type:            publishType,
		Title:           args.Title,
		Description:     args.Description,
		GrayRule:        args.GrayRule,
		CreatedBy:       user.Username,
	}

	publish, err := ctrl.manager.PublishConfiguration(ctx, publishConfigurationParams)
	if err != nil {
		return nil, err
	}

	return publish, nil
}

type RevertConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	AppName         string `uri:"app_name"`
	ConfigurationID string `uri:"configuration_id"`
	ClusterName     string `uri:"cluster_name"`
}

func (ctrl *Handler) RevertConfiguration(ctx *fox.Context, args *RevertConfigurationArgs) (*models.Publish, error) {
	var (
		logger        = ctx.Logger
		user          = ctrl.CurrentUser(ctx)
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("revert configuration, args: %+v", args)

	params := &params.RevertConfiguration{
		NamespaceName:     args.NamespaceName,
		AppName:           args.AppName,
		ClusterName:       args.ClusterName,
		ConfigurationName: configuration.Name,
		ConfigurationID:   configuration.ID,
		CreatedBy:         user.Username,
	}

	publish, err := ctrl.manager.RevertConfiguration(ctx, params)
	if err != nil {
		return nil, err
	}

	return publish, nil
}

type GetPublishesArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	AppName         string `uri:"app_name"`
	ConfigurationID string `uri:"configuration_id"`

	Q           string `query:"q"`
	IncludeGray bool   `query:"includeGray"`
	Page        int    `query:"page"`
	Size        int    `query:"size"`
}

func (ctrl *Handler) GetPublishes(ctx *fox.Context, args *GetPublishesArgs) (*models.Pagination[*models.Publish], error) {
	var (
		logger = ctx.Logger

		configuration = ctx.MustGet("configuration").(*models.Configuration)
		err           error
		pagination    *models.Pagination[*models.Publish]
	)

	logger.Debugf("get publishes, args: %+v", args)

	p := &params.GetPublishes{
		ConfigurationID: configuration.ID,
		Q:               args.Q,

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
	NamespaceName   string `uri:"namespace_name"`
	AppName         string `uri:"app_name"`
	ConfigurationID string `uri:"configuration_id"`
	PublishID       string `uri:"publish_id"`
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
