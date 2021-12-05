package request

type ApiInfoRequest struct {
	ID      uint     `json:"ID"`
	Name    string   `json:"name"`
	Method  string   `json:"method"`
	Url     string   `json:"url"`
	Params  []string `json:"params"`
	Project string   `json:"project"`
	Module  string   `json:"module"`
}
