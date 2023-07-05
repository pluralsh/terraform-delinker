package main

import (
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"

	analyzev1alpha1 "github.com/pluralsh/terraform-delinker/api/analyze/v1alpha1"
	execv1alpha1 "github.com/pluralsh/terraform-delinker/api/exec/v1alpha1"
	planv1alpha1 "github.com/pluralsh/terraform-delinker/api/plan/v1alpha1"
)

func main() {
	planner := planv1alpha1.NewPlanner(
		planv1alpha1.WithTerraform(
			execv1alpha1.WithDir("/home/floreks/gcp-capi/bootstrap/terraform"),
		),
	)

	plan, err := planner.Plan()
	if err != nil {
		panic(err)
	}

	fmt.Println(analyzev1alpha1.NewAnalyzer(plan).Analyze(tfjson.ActionCreate))
}
