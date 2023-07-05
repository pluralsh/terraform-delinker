package v1alpha1

import (
	"encoding/json"
	"os"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/sirupsen/logrus"

	"github.com/pluralsh/terraform-delinker/api/exec/v1alpha1"
	"github.com/pluralsh/terraform-delinker/internal"
)

type PlannerOption func(*Planner)

func WithTerraform(options ...v1alpha1.ExecutableOption) PlannerOption {
	return func(p *Planner) {
		p.terraform = v1alpha1.NewExecutable(v1alpha1.CommandTerraform, options...)
	}
}

func WithDir(directory string) PlannerOption {
	return func(p *Planner) {
		p.directory = directory
	}
}

type Planner struct {
	terraform *v1alpha1.Executable
	directory string
}

func (this *Planner) Plan() (*tfjson.Plan, error) {
	logrus.Debug("creating temp file")
	f, err := os.CreateTemp("", "plural.tf.*.plan")
	if err != nil {
		return nil, err
	}

	defer internal.RemoveFileOrDie(f.Name())

	logrus.Infof("creating execution plan and saving to: %s", f.Name())
	_, err = this.terraform.Run("plan", "-out", f.Name())
	if err != nil {
		return nil, err
	}

	logrus.Info("converting execution plan to JSON format")
	output, err := this.terraform.Run("show", "-json", f.Name())
	if err != nil {
		return nil, err
	}

	plan := new(tfjson.Plan)
	err = json.Unmarshal(output, plan)
	if err != nil {
		return nil, err
	}

	logrus.Debug("successfully unmarshalled execution plan")
	logrus.Debug(plan)
	return plan, nil
}

func NewPlanner(options ...PlannerOption) *Planner {
	planner := new(Planner)

	for _, o := range options {
		o(planner)
	}

	return planner
}
