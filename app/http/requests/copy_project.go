package requests

type CopyProjectForm struct {
	TargetProject int    `json:"target_project"`
	Name          string `json:"name"`
	KeyId         int    `json:"key_id"`
	RepoAddress   string `json:"repo_address"`
	DefaultBranch string `json:"default_branch"`
	GroupId       int    `json:"group_id"`
}
