package handlers

import (
	"image/png"
	"os"

	"github.com/fox-gonic/fox"
	"github.com/pquerna/otp/totp"

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

func (ctrl *Handler) TOTP(c *fox.Context) (res any) {

	// 生成 TOTP 密钥
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "YourPlatformName", // 你的平台名称
		AccountName: "user@example.com", // 用户的唯一标识（如邮箱）
	})
	if err != nil {
		panic(err)
	}

	// 打印密钥（用户可手动输入到 Google Authenticator）
	println("TOTP Secret:", key.Secret())

	// 生成二维码图片（供 Google Authenticator 扫描）
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}

	file, _ := os.Create("qr.png")
	defer file.Close()
	png.Encode(file, img)
	println("QR code saved as qr.png")

	return nil
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
