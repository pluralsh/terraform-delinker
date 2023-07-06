package v1alpha1

import tfjson "github.com/hashicorp/terraform-json"

type Report struct {
	Resources []Resource
}

type Resource struct {
	Address string
	Name    string
	Actions tfjson.Actions
	Type    string
}

type PlanAnalyzer struct {
	plan *tfjson.Plan
}
