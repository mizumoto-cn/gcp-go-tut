package createproj

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mizumoto-cn/gcp-go-tut/create_proj/project"
)

const (
	url string = "https://cloudresourcemanager.googleapis.com/v3/projects"
)

type p_json struct {
	ProjectId string            `json:"projectId"`
	Name      string            `json:"name"`
	Labels    map[string]string `json:"labels"`
}

func Post(p project.Project) (*http.Response, error) {
	pj := p_json{
		ProjectId: p.ProjectId,
		Name:      p.Name,
		Labels:    p.Labels,
	}
	data, err := json.Marshal(pj)
	if err != nil {
		return nil, err
	}
	//for test
	log.Println(string(data))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.DefaultClient
	resp, err := client.Do(req)
	return resp, err
}
