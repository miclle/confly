package handlers

import (
	"github.com/fox-gonic/fox"

	"github.com/miclle/confly/models"
	"github.com/miclle/confly/params"
)

type AuthenticationState string

const (
	AuthenticationStateAuthorized   AuthenticationState = "authorized"
	AuthenticationStateUnauthorized AuthenticationState = "unauthorized"
)

type OverviewState struct {
	Authentication AuthenticationState `json:"authentication,omitempty"`
	User           *models.User        `json:"user,omitempty"`
}

type SigninArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ctrl *Handler) Signin(c *fox.Context, args *SigninArgs) (res any) {
	// TODO(m)
	return nil
}

func (ctrl *Handler) Boot(c *fox.Context) any {
	// TODO(m)
	return nil
}

type LogoutArgs struct{}

func (ctrl *Handler) Logout(c *fox.Context, args *LogoutArgs) (res interface{}) {
	// TODO(m)
	return nil
}

func (ctrl *Handler) CurrentUser(c *fox.Context) *models.User {
	// TODO(m)
	return nil
}

type GetUsersArgs struct {
	models.Pagination[*models.User]

	Q string `query:"q"`
}

func (ctrl *Handler) GetUsers(c *fox.Context, args *GetUsersArgs) (*models.Pagination[*models.User], error) {

	pagination, err := ctrl.manager.GetUsers(c, &params.GetUsers{
		Pagination: args.Pagination,
		Q:          args.Q,
	})

	if err != nil {
		c.Logger.Errorf("get users failed, error: %+v", err)
	}

	return pagination, err
}

type UpdateUserArgs struct {
	Username string  `uri:"username"`
	Name     *string `json:"name"`
}

func (ctrl *Handler) UpdateUser(c *fox.Context, args *UpdateUserArgs) (*models.User, error) {

	var user = ctrl.CurrentUser(c)

	user, err := ctrl.manager.UpdateUser(c, args.Username, &params.UpdateUser{
		Name:      args.Name,
		UpdatedBy: user.Username,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
