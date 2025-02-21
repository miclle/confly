package params

import "github.com/miclle/confly/models"

type CreateApp struct {
	NamespaceID string
	Name        string
	Description string
	CreatedBy   string
}

type GetApps struct {
	models.Pagination[*models.Application]

	Q           string
	NamespaceID string
}

type GetApp struct {
	NamespaceID string
	ID          string
	Name        string
}

type UpdateApp struct {
	Description *string
	UpdatedBy   string
}
