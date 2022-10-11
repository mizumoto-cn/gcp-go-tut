package operation

type Operation struct {
	Name     string            `json:"op"`
	Metadata map[string]string `json:"metadata"`
	Done     bool              `json:"done"`
	Error    map[string]string `json:"error"`
	Response map[string]string `json:"response"`
}
