// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

const (
	Oauth2AuthenticationScopes = "oauth2Authentication.Scopes"
)

// ClusterManagerRead A cluster manager.
type ClusterManagerRead struct {
	Metadata externalRef0.ProjectScopedResourceReadMetadata `json:"metadata"`
}

// ClusterManagerWrite A cluster manager.
type ClusterManagerWrite struct {
	// Metadata Resource metadata valid for all API resource reads and writes.
	Metadata externalRef0.ResourceWriteMetadata `json:"metadata"`
}

// ClusterManagers A list of cluster managers.
type ClusterManagers = []ClusterManagerRead

// KubernetesClusterAPI Kubernetes API settings.
type KubernetesClusterAPI struct {
	// AllowedPrefixes Set of address prefixes to allow access to the Kubernetes API.
	AllowedPrefixes *[]string `json:"allowedPrefixes,omitempty"`

	// SubjectAlternativeNames Set of non-standard X.509 SANs to add to the API certificate.
	SubjectAlternativeNames *[]string `json:"subjectAlternativeNames,omitempty"`
}

// KubernetesClusterAutoscaling A Kubernetes cluster workload pool autoscaling configuration. Cluster autoscaling
// must also be enabled in the cluster features.
type KubernetesClusterAutoscaling struct {
	// MinimumReplicas The minimum number of replicas to allow. Must be less than the maximum.
	MinimumReplicas int `json:"minimumReplicas"`
}

// KubernetesClusterNetwork A kubernetes cluster network settings.
type KubernetesClusterNetwork struct {
	// DnsNameservers A list of DNS name server to use.
	DnsNameservers *[]string `json:"dnsNameservers,omitempty"`

	// NodePrefix Network prefix to provision nodes in. Must be a valid CIDR block.
	NodePrefix *string `json:"nodePrefix,omitempty"`

	// PodPrefix Network prefix to provision pods in. Must be a valid CIDR block.
	PodPrefix *string `json:"podPrefix,omitempty"`

	// ServicePrefix Network prefix to provision services in. Must be a valid CIDR block.
	ServicePrefix *string `json:"servicePrefix,omitempty"`
}

// KubernetesClusterRead Kubernetes cluster read.
type KubernetesClusterRead struct {
	Metadata externalRef0.ProjectScopedResourceReadMetadata `json:"metadata"`

	// Spec Kubernetes cluster creation parameters.
	Spec KubernetesClusterSpec `json:"spec"`
}

// KubernetesClusterSpec Kubernetes cluster creation parameters.
type KubernetesClusterSpec struct {
	// Api Kubernetes API settings.
	Api *KubernetesClusterAPI `json:"api,omitempty"`

	// ClusterManagerId The name of the cluster manager to use, if one is not specified
	// the system will create one for you.
	ClusterManagerId *string `json:"clusterManagerId,omitempty"`

	// Networking A kubernetes cluster network settings.
	Networking *KubernetesClusterNetwork `json:"networking,omitempty"`

	// RegionId The region to provision the cluster in.
	RegionId string `json:"regionId"`

	// Version The Kuebernetes version.  This should be derived from image metadata.
	Version string `json:"version"`

	// WorkloadPools A list of Kubernetes cluster workload pools.
	WorkloadPools KubernetesClusterWorkloadPools `json:"workloadPools"`
}

// KubernetesClusterWorkloadPool A Kuberntes cluster workload pool.
type KubernetesClusterWorkloadPool struct {
	// Autoscaling A Kubernetes cluster workload pool autoscaling configuration. Cluster autoscaling
	// must also be enabled in the cluster features.
	Autoscaling *KubernetesClusterAutoscaling `json:"autoscaling,omitempty"`

	// Labels Workload pool key value labels to apply on node creation.
	Labels *map[string]string `json:"labels,omitempty"`

	// Machine A Kubernetes cluster machine.
	Machine MachinePool `json:"machine"`

	// Name Workload pool name.
	Name string `json:"name"`
}

// KubernetesClusterWorkloadPools A list of Kubernetes cluster workload pools.
type KubernetesClusterWorkloadPools = []KubernetesClusterWorkloadPool

// KubernetesClusterWrite Kubernetes cluster create or update.
type KubernetesClusterWrite struct {
	// Metadata Resource metadata valid for all API resource reads and writes.
	Metadata externalRef0.ResourceWriteMetadata `json:"metadata"`

	// Spec Kubernetes cluster creation parameters.
	Spec KubernetesClusterSpec `json:"spec"`
}

// KubernetesClusters A list of Kubernetes clusters.
type KubernetesClusters = []KubernetesClusterRead

// KubernetesNameParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type KubernetesNameParameter = string

// MachinePool A Kubernetes cluster machine.
type MachinePool struct {
	// Disk A volume.
	Disk *Volume `json:"disk,omitempty"`

	// FlavorId Flavor ID.
	FlavorId *string `json:"flavorId,omitempty"`

	// Replicas Number of machines for a statically sized pool or the maximum for an auto-scaled pool.
	Replicas *int `json:"replicas,omitempty"`
}

// Volume A volume.
type Volume struct {
	// Size Disk size in GiB.
	Size int `json:"size"`
}

// ClusterIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type ClusterIDParameter = KubernetesNameParameter

// ClusterManagerIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type ClusterManagerIDParameter = KubernetesNameParameter

// OrganizationIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type OrganizationIDParameter = KubernetesNameParameter

// ProjectIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type ProjectIDParameter = KubernetesNameParameter

// RegionIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type RegionIDParameter = KubernetesNameParameter

// ClusterManagerResponse A cluster manager.
type ClusterManagerResponse = ClusterManagerRead

// ClusterManagersResponse A list of cluster managers.
type ClusterManagersResponse = ClusterManagers

// KubernetesClusterResponse Kubernetes cluster read.
type KubernetesClusterResponse = KubernetesClusterRead

// KubernetesClustersResponse A list of Kubernetes clusters.
type KubernetesClustersResponse = KubernetesClusters

// CreateControlPlaneRequest A cluster manager.
type CreateControlPlaneRequest = ClusterManagerWrite

// CreateKubernetesClusterRequest Kubernetes cluster create or update.
type CreateKubernetesClusterRequest = KubernetesClusterWrite

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDClustermanagersJSONRequestBody defines body for PostApiV1OrganizationsOrganizationIDProjectsProjectIDClustermanagers for application/json ContentType.
type PostApiV1OrganizationsOrganizationIDProjectsProjectIDClustermanagersJSONRequestBody = ClusterManagerWrite

// PutApiV1OrganizationsOrganizationIDProjectsProjectIDClustermanagersClusterManagerIDJSONRequestBody defines body for PutApiV1OrganizationsOrganizationIDProjectsProjectIDClustermanagersClusterManagerID for application/json ContentType.
type PutApiV1OrganizationsOrganizationIDProjectsProjectIDClustermanagersClusterManagerIDJSONRequestBody = ClusterManagerWrite

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDClustersJSONRequestBody defines body for PostApiV1OrganizationsOrganizationIDProjectsProjectIDClusters for application/json ContentType.
type PostApiV1OrganizationsOrganizationIDProjectsProjectIDClustersJSONRequestBody = KubernetesClusterWrite

// PutApiV1OrganizationsOrganizationIDProjectsProjectIDClustersClusterIDJSONRequestBody defines body for PutApiV1OrganizationsOrganizationIDProjectsProjectIDClustersClusterID for application/json ContentType.
type PutApiV1OrganizationsOrganizationIDProjectsProjectIDClustersClusterIDJSONRequestBody = KubernetesClusterWrite
