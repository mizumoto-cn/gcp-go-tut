package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type State uint8

const (
	STATE_UNSPECIFIED = iota
	ACTIVE
	DELETE_REQUESTED
)

func (s State) String() string {
	switch s {
	case ACTIVE:
		return "ACTIVE"
	case DELETE_REQUESTED:
		return "DELETE_REQUESTED"
	default:
		return "STATE_UNSPECIFIED"
	}
}

type Project struct {
	Name        string            `json:"name"`
	Parent      string            `json:"parent"`
	ProjectId   string            `json:"projectId"`
	State       string            `json:"state"`
	DisplayName string            `json:"displayName"`
	CreateTime  string            `json:"createTime"`
	UpdateTime  string            `json:"updateTime"`
	DeleteTime  string            `json:"deleteTime"`
	Etag        string            `json:"etag"`
	Labels      map[string]string `json:"labels"`
}

func (p *Project) MarshalJSON() ([]byte, error) {
	type Alias Project
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(*p),
	})
}

func main() {
	m := make(map[string]string)
	m["roses"] = "red"
	m["violets"] = "blue"
	p := Project{
		Name:        "Default-Proj-Name",
		Parent:      "",
		ProjectId:   "Default-Proj-Id",
		State:       "",
		DisplayName: "",
		CreateTime:  "",
		UpdateTime:  "",
		DeleteTime:  "",
		Etag:        "",
		Labels:      m,
	}
	b, err := json.Marshal(p)
	if err != nil {
		log.Println("Error marshalling project to json:", err)
	}
	//for test
	fmt.Println(string(b))
}
