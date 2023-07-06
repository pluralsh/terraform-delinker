package v1alpha1

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/sirupsen/logrus"

	"github.com/pluralsh/terraform-delinker/internal"
)

func (this *PlanAnalyzer) Analyze(filter ...tfjson.Action) *Report {
	report := &Report{
		Resources: make([]Resource, 0),
	}

	for _, resource := range this.plan.ResourceChanges {
		logrus.Debugf("checking if resource %s matches %v action", resource.Address, filter)
		if includes := internal.IncludesArray(resource.Change.Actions, filter); len(filter) > 0 && !includes {
			continue
		}

		report.Resources = append(report.Resources, Resource{
			Address: resource.Address,
			Name:    resource.Name,
			Actions: resource.Change.Actions,
			Type:    resource.Type,
		})
	}

	logrus.Debugf("found %d resources", len(report.Resources))
	return report
}

func NewAnalyzer(plan *tfjson.Plan) *PlanAnalyzer {
	if plan == nil {
		panic("plan cannot be nil")
	}

	return &PlanAnalyzer{
		plan: plan,
	}
}
