package params

import "github.com/miclle/confly/models"

type CreateApplication struct {
	NamespaceID string
	Name        string
	Description string
	CreatedBy   string
}

type GetApplications struct {
	models.Pagination[*models.Application]

	Q           string
	NamespaceID string
	Page        int
	Size        int
}

type GetApplication struct {
	NamespaceID string
	ID          string
	Name        string
}

type UpdateApplication struct {
	Description *string
	UpdatedBy   string
}
