package v1alpha1

import (
	execv1alpha1 "github.com/pluralsh/terraform-delinker/api/exec/v1alpha1"
)

type Delinker struct {
	terraform *execv1alpha1.Executable
}

type DelinkerOption func(*Delinker)
