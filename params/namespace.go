package params

import "github.com/miclle/confly/models"

type CreateNamespace struct {
	Name        string
	Description string
	CreatedBy   string
}

type GetNamespaces struct {
	models.Pagination[*models.Namespace]

	Q string
}

type GetNamespace struct {
	Name string
	ID   string
}

type UpdateNamespace struct {
	Description *string
	UpdatedBy   string
}
