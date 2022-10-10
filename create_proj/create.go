package createproj

import (
	"os"
	"os/exec"

	"github.com/mizumoto-cn/gcp-go-tut/create_proj/project"
)

func CreateProj(p project.Project) error {
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
