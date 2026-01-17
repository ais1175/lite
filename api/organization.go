package api

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}

type OrganizationStats struct {
	TotalFiles   int   `json:"totalFiles"`
	TotalSize    int64 `json:"totalSize"`
	TotalLogs    int   `json:"totalLogs"`
	TotalTokens  int   `json:"totalTokens"`
	TotalDataset int   `json:"totalDataset"`
}
