package v1alpha1

import "github.com/pluralsh/terraform-delinker/api/exec/v1alpha1"

type Planner struct {
	terraform *v1alpha1.Executable
	directory string
}

type PlannerOption func(*Planner)
