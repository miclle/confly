package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type PublishConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`
	FullRelease     bool   `query:"fullRelease"`

	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	GrayRule    *models.ReleaseGrayRule `json:"grayRule"`

	MergedConfigChecksum string `json:"mergedConfigChecksum"`
}

func (ctrl *Handler) PublishConfiguration(ctx *fox.Context, args *PublishConfigurationArgs) (*models.ConfigurationRelease, error) {
	var (
		logger        = ctx.Logger
		user          = ctrl.CurrentUser(ctx)
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("publish configuration, args: %+v", args)

	releaseType := models.ReleaseTypeGray
	if args.FullRelease {
		releaseType = models.ReleaseTypeFull
	}

	params := &params.PublishConfiguration{
		ConfigurationID: configuration.ID,
		Type:            releaseType,
		Title:           args.Title,
		Description:     args.Description,
		GrayRule:        args.GrayRule,
		CreatedBy:       user.Username,
	}

	release, err := ctrl.manager.PublishConfiguration(ctx, params)
	if err != nil {
		return nil, err
	}

	return release, nil
}

type RevertConfigurationArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`
	Environment     string `uri:"environment"`
}

func (ctrl *Handler) RevertConfiguration(ctx *fox.Context, args *RevertConfigurationArgs) (*models.ConfigurationRelease, error) {
	var (
		logger        = ctx.Logger
		user          = ctrl.CurrentUser(ctx)
		configuration = ctx.MustGet("configuration").(*models.Configuration)
	)

	logger.Debugf("revert configuration, args: %+v", args)

	params := &params.RevertConfiguration{
		NamespaceName:     args.NamespaceName,
		ApplicationName:   args.ApplicationName,
		Environment:       args.Environment,
		ConfigurationName: configuration.Name,
		ConfigurationID:   configuration.ID,
		CreatedBy:         user.Username,
	}

	release, err := ctrl.manager.RevertConfiguration(ctx, params)
	if err != nil {
		return nil, err
	}

	return release, nil
}

type GetReleasesArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`

	Q           string `query:"q"`
	IncludeGray bool   `query:"includeGray"`
	Page        int    `query:"page"`
	Size        int    `query:"size"`
}

func (ctrl *Handler) GetReleases(ctx *fox.Context, args *GetReleasesArgs) (*models.Pagination[*models.ConfigurationRelease], error) {
	var (
		logger = ctx.Logger

		configuration = ctx.MustGet("configuration").(*models.Configuration)
		err           error
		pagination    *models.Pagination[*models.ConfigurationRelease]
	)

	logger.Debugf("get releases, args: %+v", args)

	p := &params.GetReleases{
		ConfigurationID: configuration.ID,
		Q:               args.Q,

		Pagination: models.Pagination[*models.ConfigurationRelease]{
			Page: args.Page,
			Size: args.Size,
		},
	}

	pagination, err = ctrl.manager.GetReleases(ctx, p)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

type GetReleaseArgs struct {
	NamespaceName   string `uri:"namespace_name"`
	ApplicationName string `uri:"application_name"`
	ConfigurationID string `uri:"configuration_id"`
	VersionID       string `uri:"version_id"`
}

func (ctrl *Handler) GetRelease(ctx *fox.Context, args *GetReleaseArgs) (*models.ConfigurationRelease, error) {
	var (
		logger  = ctx.Logger
		release = ctx.MustGet("release").(*models.ConfigurationRelease)
	)

	logger.Debugf("get release, args: %+v", args)

	return release, nil
}

func (ctrl *Handler) SetRelease(ctx *fox.Context, args *GetReleaseArgs) (res any) {

	release, err := ctrl.manager.GetRelease(ctx, &params.GetRelease{
		VersionID: args.VersionID,
	})

	if err != nil {
		ctx.Logger.Errorf("get release failed, err: %+v", err)
		return err
	}

	ctx.Set("release", release)

	return
}
