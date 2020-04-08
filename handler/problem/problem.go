package problem

import "experiment/model"

type UploadRequest struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Example     model.Example `json:"example"`
	Solution    string        `json:"solution"`
	Output      model.Output  `json:"output"`
	Poster      string        `json:"poster"`
}

type DetailResponse struct {
	ProblemId   int           `json:"problem_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Example     model.Example `json:"example"`
	Output      model.Output  `json:"output"`
	Poster      string        `json:"poster"`
}
