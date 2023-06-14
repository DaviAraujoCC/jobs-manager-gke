package models

type GetCronjobsResponse struct {
	Total int      `json:"total"`
	Data  []string `json:"data"`
}

type GetJobsResponse struct {
	Total int       `json:"total"`
	Data  []JobInfo `json:"data"`
}

type JobInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
}
