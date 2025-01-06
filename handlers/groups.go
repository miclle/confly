package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type GetGroupsArgs struct {
	Q string `query:"q"`
}

func (ctrl *Handler) GetGroups(ctx *fox.Context, args *GetGroupsArgs) (*models.Pagination[*models.Group], error) {
	logger := ctx.Logger
	logger.Debugf("get groups args: %+v", args)

	pagination, err := ctrl.manager.GetGroups(ctx, &params.GetGroups{
		Q: args.Q,
	})

	if err != nil {
		logger.Errorf("get groups error: %+v", err)
		return nil, err
	}

	return pagination, nil
}

func (ctrl *Handler) GetGroup(ctx *fox.Context, args *GetGroupArgs) (*models.Group, error) {

	var group = ctx.MustGet("group").(*models.Group)

	return group, nil
}

type UpdategroupArgs struct {
	Description *string `json:"description"`
}

func (ctrl *Handler) UpdateGroup(ctx *fox.Context, args *UpdategroupArgs) error {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
		group  = ctx.MustGet("group").(*models.Group)
	)

	logger.Debugf("update group args: %+v", args)

	err := ctrl.manager.UpdateGroup(ctx, group.ID, &params.UpdateGroup{
		Description: args.Description,
		UpdatedBy:   user.Username,
	})

	if err != nil {
		logger.Errorf("update group error: %+v", err)
		return err
	}
	return nil
}

type CreateGroupArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (ctrl *Handler) CreateGroup(ctx *fox.Context, args *CreateGroupArgs) (*models.Group, error) {
	var (
		logger = ctx.Logger
		user   = ctrl.CurrentUser(ctx)
	)

	logger.Debugf("create group args: %+v", args)

	params := &params.CreateGroup{
		Name:        args.Name,
		Description: args.Description,
		CreatedBy:   user.Username,
	}

	group, err := ctrl.manager.CreateGroup(ctx, params)
	if err != nil {
		logger.Errorf("create group error: %+v", err)
		return nil, err
	}

	return group, nil
}

type DeleteGroupArgs struct {
	Name string `uri:"group_name"`
}

func (ctrl *Handler) DeleteGroup(ctx *fox.Context, args *DeleteGroupArgs) error {
	var (
		logger = ctx.Logger
		group  = ctx.MustGet("group").(*models.Group)
	)

	logger.Debugf("delete group arg: %+v", group.Name)

	err := ctrl.manager.DeleteGroup(ctx, group.ID)
	if err != nil {
		logger.Errorf("delete group error: %+v", err)
		return err
	}

	return nil
}

type GetGroupArgs struct {
	Name string `uri:"group_name"`
}

func (ctrl *Handler) SetGroup(ctx *fox.Context, args *GetGroupArgs) (res any) {

	group, err := ctrl.manager.GetGroup(ctx, &params.GetGroup{
		Name: args.Name,
	})
	if err != nil {
		ctx.Logger.Errorf("get group error: %+v", err)
		return err
	}

	ctx.Set("group", group)

	return
}
