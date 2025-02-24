package params

import "github.com/miclle/confly/models"

type CreateService struct {
	NamespaceID string
	Name        string
	Description string
	CreatedBy   string
}

type GetServices struct {
	models.Pagination[*models.Service]

	Q           string
	NamespaceID string
	Page        int
	Size        int
}

type GetService struct {
	NamespaceID string
	ID          string
	Name        string
}

type UpdateService struct {
	Description *string
	UpdatedBy   string
}
