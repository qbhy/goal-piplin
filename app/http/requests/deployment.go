package requests

type DeploymentRequest struct {
	ProjectId    int             `json:"project_id"`
	Version      string          `json:"version"`
	Comment      string          `json:"comment"`
	Params       map[string]bool `json:"params"`
	Environments []int           `json:"environments"`
}
