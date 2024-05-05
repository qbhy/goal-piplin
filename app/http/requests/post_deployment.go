package requests

type PostDeploymentRequest struct {
	UUID         string          `json:"uuid"`
	Version      string          `json:"version"`
	Comment      string          `json:"comment"`
	Params       map[string]bool `json:"params"`
	Environments []int           `json:"environments"`
}
