package requests

type RollbackDeploymentRequest struct {
	Id            int    `json:"id"`
	Commands      []int  `json:"commands"`
	BeforeRelease string `json:"before_release"`
	AfterRelease  string `json:"after_release"`
}
