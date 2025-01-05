package params

import "github.com/miclle/confly/models"

type GetUsers struct {
	models.Pagination[*models.User]

	Q string
}

type UpdateUser struct {
	Name      *string
	UpdatedBy string
}
