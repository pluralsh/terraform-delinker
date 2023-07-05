package v1alpha1

import (
	tfjson "github.com/hashicorp/terraform-json"

	"github.com/pluralsh/terraform-delinker/internal"
)

type PlanAnalyzer struct {
	plan *tfjson.Plan
}

func (this *PlanAnalyzer) Analyze(filter ...tfjson.Action) *Report {
	report := &Report{
		Resources: make([]Resource, 0),
	}

	for _, resource := range this.plan.ResourceChanges {
		if includes := internal.IncludesArray(resource.Change.Actions, filter); !includes {
			continue
		}

		report.Resources = append(report.Resources, Resource{
			Address: resource.Address,
			Name:    resource.Name,
			Actions: resource.Change.Actions,
			Type:    resource.Type,
		})
	}

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
