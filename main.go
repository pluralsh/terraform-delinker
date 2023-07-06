package main

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/sirupsen/logrus"

	analyzev1alpha1 "github.com/pluralsh/terraform-delinker/api/analyze/v1alpha1"
	"github.com/pluralsh/terraform-delinker/api/delink/v1alpha1"
	execv1alpha1 "github.com/pluralsh/terraform-delinker/api/exec/v1alpha1"
	planv1alpha1 "github.com/pluralsh/terraform-delinker/api/plan/v1alpha1"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

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

	report := analyzev1alpha1.NewAnalyzer(plan).Analyze(tfjson.ActionDelete)

	delinker := v1alpha1.NewDelinker(
		v1alpha1.WithTerraform(
			execv1alpha1.WithDir("/home/floreks/gcp-capi/bootstrap/terraform"),
		),
	)

	if err := delinker.Run(report); err != nil {
		logrus.Fatal(err)
	}
}
