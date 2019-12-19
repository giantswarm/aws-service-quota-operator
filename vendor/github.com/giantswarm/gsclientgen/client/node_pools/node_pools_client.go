// Code generated by go-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new node pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddNodePool creates node pool

This allows to add a node pool to a cluster.

Some, but not all, node pool configuration can be changed after
creation. The following settings will have a permanent effect:

- `availability_zones`
- `node_spec.aws.instance_type`

*/
func (a *Client) AddNodePool(params *AddNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*AddNodePoolCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addNodePool",
		Method:             "POST",
		PathPattern:        "/v5/clusters/{cluster_id}/nodepools/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddNodePoolCreated), nil

}

/*
DeleteNodePool deletes node pool

Triggers the deletion of a node pool.

Nodes in the pool will first be cordoned and drained. Note that it is
your responsibililty to make sure that workloads using the node pool can
be scheduled elsewhere. We recommend to double-check the available
capacity of remaining node pools, as well as any node selectors and
taints. Also you can do the draining yourself before issuing the
delete request, to observe the outcome. Use

```
kubectl drain nodes -l giantswarm.io/machine-deployment=<nodepool_id> ...
```

*/
func (a *Client) DeleteNodePool(params *DeleteNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodePoolAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodePool",
		Method:             "DELETE",
		PathPattern:        "/v5/clusters/{cluster_id}/nodepools/{nodepool_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodePoolAccepted), nil

}

/*
GetNodePool gets node pool details

Returns all available details on a specific node pool.

*/
func (a *Client) GetNodePool(params *GetNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodePool",
		Method:             "GET",
		PathPattern:        "/v5/clusters/{cluster_id}/nodepools/{nodepool_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodePoolOK), nil

}

/*
GetNodePools gets node pools

Returns a list of node pools from a given cluster.

*/
func (a *Client) GetNodePools(params *GetNodePoolsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodePoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodePoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodePools",
		Method:             "GET",
		PathPattern:        "/v5/clusters/{cluster_id}/nodepools/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNodePoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodePoolsOK), nil

}

/*
ModifyNodePool modifies node pool

Allows to rename a node pool or change its scaling settings.

*/
func (a *Client) ModifyNodePool(params *ModifyNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyNodePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyNodePool",
		Method:             "PATCH",
		PathPattern:        "/v5/clusters/{cluster_id}/nodepools/{nodepool_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyNodePoolOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
