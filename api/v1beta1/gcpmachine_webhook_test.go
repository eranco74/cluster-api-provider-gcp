/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGCPMachine_ValidateCreate(t *testing.T) {
	g := NewWithT(t)
	confidentialComputeEnabled := ConfidentialComputePolicyEnabled
	onHostMaintenanceTerminate := HostMaintenancePolicyTerminate
	onHostMaintenanceMigrate := HostMaintenancePolicyMigrate
	tests := []struct {
		name string
		*GCPMachine
		wantErr bool
	}{
		{
			name: "GCPMachined with OnHostMaintenance set to TERMINATE - valid",
			GCPMachine: &GCPMachine{
				Spec: GCPMachineSpec{
					OnHostMaintenance: &onHostMaintenanceTerminate,
				},
			},
			wantErr: false,
		},
		{
			name: "GCPMachined with ConfidentialCompute enabled and OnHostMaintenance set to TERMINATE - valid",
			GCPMachine: &GCPMachine{
				Spec: GCPMachineSpec{
					OnHostMaintenance:   &onHostMaintenanceTerminate,
					ConfidentialCompute: &confidentialComputeEnabled,
				},
			},
			wantErr: false,
		},
		{
			name: "GCPMachined with ConfidentialCompute enabled and OnHostMaintenance set to MIGRATE - invalid",
			GCPMachine: &GCPMachine{
				Spec: GCPMachineSpec{
					OnHostMaintenance:   &onHostMaintenanceMigrate,
					ConfidentialCompute: &confidentialComputeEnabled,
				},
			},
			wantErr: true,
		},
		{
			name: "GCPMachined with ConfidentialCompute enabled and default OnHostMaintenance (MIGARTE) - invalid",
			GCPMachine: &GCPMachine{
				Spec: GCPMachineSpec{
					ConfidentialCompute: &confidentialComputeEnabled,
				},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.GCPMachine.ValidateCreate()
			if test.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}
		})
	}
}
