package request

type ApiInfoRequest struct {
	ID      string   `json:"ID"`
	Name    string   `json:"name"`
	Method  string   `json:"method"`
	Url     string   `json:"url"`
	Params  []string `json:"params"`
	Project string   `json:"project"`
	Module  string   `json:"module"`
}
