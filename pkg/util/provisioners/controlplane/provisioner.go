/*
Copyright 2022 EscherCloud.

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

package controlplane

import (
	"context"

	unikornv1alpha1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	"github.com/eschercloudai/unikorn/pkg/util/provisioners/clusterapi"
	provisioner "github.com/eschercloudai/unikorn/pkg/util/provisioners/generic"
	"github.com/eschercloudai/unikorn/pkg/util/provisioners/vcluster"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provisioner encapsulates control plane provisioning.
type Provisioner struct {
	// client provides access to Kubernetes.
	client client.Client

	// controlPlane is the control plane CR this deployment relates to
	controlPlane *unikornv1alpha1.ControlPlane
}

// New returns a new initialized provisioner object.
func New(client client.Client, controlPlane *unikornv1alpha1.ControlPlane) *Provisioner {
	return &Provisioner{
		client:       client,
		controlPlane: controlPlane,
	}
}

// Ensure the Provisioner interface is implemented.
var _ provisioner.Provisioner = &Provisioner{}

// Provision implements the Provision interface.
func (p *Provisioner) Provision(ctx context.Context) error {
	// Provision a virtual cluster for CAPI to live in.
	vclusterProvisioner := vcluster.New(p.client, p.controlPlane)

	if err := vclusterProvisioner.Provision(ctx); err != nil {
		return err
	}

	// Create a new client that's able to talk to the vcluster.
	vclusterConfig, err := vcluster.RESTConfig(ctx, p.client, p.controlPlane.Namespace, p.controlPlane.Name)
	if err != nil {
		return err
	}

	vclusterClient, err := client.New(vclusterConfig, client.Options{Scheme: p.client.Scheme(), Mapper: p.client.RESTMapper()})
	if err != nil {
		return err
	}

	// Provision CAPI in the vcluster.
	clusterAPIProvisioner := clusterapi.New(vclusterClient)

	if err := clusterAPIProvisioner.Provision(ctx); err != nil {
		return err
	}

	return nil
}