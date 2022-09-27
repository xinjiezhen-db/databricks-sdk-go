// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewClusterPolicies(client *client.DatabricksClient) ClusterPoliciesService {
	return &ClusterPoliciesAPI{
		client: client,
	}
}

type ClusterPoliciesAPI struct {
	client *client.DatabricksClient
}

// Create a new policy
//
// Creates a new policy with prescribed settings.
func (a *ClusterPoliciesAPI) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {
	var createPolicyResponse CreatePolicyResponse
	path := "/api/2.0/policies/clusters/create"
	err := a.client.Post(ctx, path, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

// Delete a cluster policy
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) Delete(ctx context.Context, request DeletePolicy) error {
	path := "/api/2.0/policies/clusters/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete a cluster policy
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) DeleteByPolicyId(ctx context.Context, policyId string) error {
	return a.Delete(ctx, DeletePolicy{
		PolicyId: policyId,
	})
}

// Update a cluster policy
//
// Update an existing policy for cluster. This operation may make some clusters
// governed by the previous policy invalid.
func (a *ClusterPoliciesAPI) Edit(ctx context.Context, request EditPolicy) error {
	path := "/api/2.0/policies/clusters/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Get entity
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) Get(ctx context.Context, request GetRequest) (*Policy, error) {
	var policy Policy
	path := "/api/2.0/policies/clusters/get"
	err := a.client.Get(ctx, path, request, &policy)
	return &policy, err
}

// Get entity
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) GetByPolicyId(ctx context.Context, policyId string) (*Policy, error) {
	return a.Get(ctx, GetRequest{
		PolicyId: policyId,
	})
}

// Get a cluster policy
//
// Returns a list of policies accessible by the requesting user.
func (a *ClusterPoliciesAPI) List(ctx context.Context) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	err := a.client.Get(ctx, path, nil, &listPoliciesResponse)
	return &listPoliciesResponse, err
}