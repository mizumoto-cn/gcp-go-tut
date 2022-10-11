package project

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

// Project contains the information about a project
//
//	{
//		"name": string,
//		"parent": string,
//		"projectId": string,
//		"state": enum (State),
//		"displayName": string,
//		"createTime": string,
//		"updateTime": string,
//		"deleteTime": string,
//		"etag": string,
//		"labels": {
//		  string: string,
//		  ...
//		}
//	}
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

type ProjectBuilder struct {
	Project
	err error
}

func NewProjectBuilder() *ProjectBuilder {
	p := new(ProjectBuilder)
	p.Name = "Default-Proj-Name"
	p.ProjectId = "Default-Proj-Id"
	return p
}

func (p *ProjectBuilder) Build() (Project, error) {
	if p.err != nil {
		return Project{}, p.err
	}
	return p.Project, nil
}

func (p *ProjectBuilder) SetName(name string) *ProjectBuilder {
	p.Name = name
	return p
}

func (p *ProjectBuilder) SetProjectId(id string) *ProjectBuilder {
	p.ProjectId = id
	return p
}

func (p *ProjectBuilder) SetParent(parent string) *ProjectBuilder {
	p.Parent = parent
	return p
}

func (p *ProjectBuilder) SetState(state State) *ProjectBuilder {
	p.State = state.String()
	return p
}

func (p *ProjectBuilder) SetDisplayName(name string) *ProjectBuilder {
	p.DisplayName = name
	return p
}

func (p *ProjectBuilder) SetCreateTime(time string) *ProjectBuilder {
	p.CreateTime = time
	return p
}

func (p *ProjectBuilder) SetUpdateTime(time string) *ProjectBuilder {

	p.UpdateTime = time
	return p
}

func (p *ProjectBuilder) SetDeleteTime(time string) *ProjectBuilder {
	p.DeleteTime = time
	return p
}

func (p *ProjectBuilder) SetEtag(etag string) *ProjectBuilder {
	p.Etag = etag
	return p
}

func (p *ProjectBuilder) SetLabels(labels map[string]string) *ProjectBuilder {
	p.Labels = labels
	return p
}
