package createproj

import (
	"io"
	"os"
	"os/exec"

	"github.com/mizumoto-cn/gcp-go-tut/create_proj/project"
)

type CreatorType uint

const (
	CLI CreatorType = iota
	API
)

func (t CreatorType) String() string {
	switch t {
	case CLI:
		return "CLI"
	default:
		return "API"
	}
}

func CreateProj(p project.Project, t CreatorType) error {
	switch t {
	case CLI:
		return createProjWithCLI(p)
	default:
		return createProjWithAPI(p)
	}
}

func createProjWithAPI(p project.Project) error {
	resp, err := Post(p)
	if err != nil {
		return err
	}
	// print resp for test
	println(resp.Status)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	println(string(data))
	return nil
}

func createProjWithCLI(p project.Project) error {
	var labels string
	for k, v := range p.Labels {
		labels += k + "=" + v + ","
	}
	// remove last comma
	labels = labels[:len(labels)-1]
	cmd := exec.Command("gcloud", "projects", "create", p.ProjectId, "--name="+p.Name, "--labels="+labels)
	println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
