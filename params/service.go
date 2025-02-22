package params

import "github.com/miclle/confly/models"

type CreateService struct {
	NamespaceID string `json:"namespace_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string
}

type GetServices struct {
	models.Pagination[*models.Service]

	Q           string
	NamespaceID string `json:"namespace_id"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}

type GetService struct {
	ID string `json:"id"`
}

type UpdateService struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
