package api

type Organization struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}
