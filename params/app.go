package params

import "github.com/miclle/confly/models"

type CreateApp struct {
	GroupID     string
	Name        string
	Description string
	CreatedBy   string
}

type GetApps struct {
	models.Pagination[*models.App]

	Q       string
	GroupID string
}

type GetApp struct {
	GroupID string
	ID      string
	Name    string
}

type UpdateApp struct {
	Description *string
	UpdatedBy   string
}
