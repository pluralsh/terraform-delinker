package v1alpha1

import (
	"github.com/sirupsen/logrus"

	analyzev1alpha1 "github.com/pluralsh/terraform-delinker/api/analyze/v1alpha1"
	execv1alpha1 "github.com/pluralsh/terraform-delinker/api/exec/v1alpha1"
)

func WithTerraform(options ...execv1alpha1.ExecutableOption) DelinkerOption {
	return func(p *Delinker) {
		p.terraform = execv1alpha1.NewExecutable(execv1alpha1.CommandTerraform, options...)
	}
}

func (this *Delinker) Run(report *analyzev1alpha1.Report) error {
	if report == nil || len(report.Resources) == 0 {
		logrus.Info("report is empty")
		return nil
	}

	logrus.Info("delinking resources")
	for _, r := range report.Resources {
		logrus.WithFields(logrus.Fields{"name": r.Address}).Info("Delinking resource from terraform state")
		output, err := this.terraform.Run("state", "rm", r.Address)
		if err != nil {
			return err
		}

		logrus.Debug(output)
	}

	logrus.Info("delinking completed successfully")
	return nil
}

func NewDelinker(options ...DelinkerOption) *Delinker {
	delinker := new(Delinker)

	for _, o := range options {
		o(delinker)
	}

	if delinker.terraform == nil {
		logrus.Fatal("terraform not defined, use WithTerraform option to configure delinker")
	}

	return delinker
}
