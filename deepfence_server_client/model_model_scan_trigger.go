/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelScanTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanTrigger{}

// ModelScanTrigger struct for ModelScanTrigger
type ModelScanTrigger struct {
	NodeId string `json:"node_id"`
	ResourceId string `json:"resource_id"`
	ResourceType ControlsScanResource `json:"resource_type"`
}

// NewModelScanTrigger instantiates a new ModelScanTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanTrigger(nodeId string, resourceId string, resourceType ControlsScanResource) *ModelScanTrigger {
	this := ModelScanTrigger{}
	this.NodeId = nodeId
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewModelScanTriggerWithDefaults instantiates a new ModelScanTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanTriggerWithDefaults() *ModelScanTrigger {
	this := ModelScanTrigger{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *ModelScanTrigger) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelScanTrigger) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelScanTrigger) SetNodeId(v string) {
	o.NodeId = v
}

// GetResourceId returns the ResourceId field value
func (o *ModelScanTrigger) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ModelScanTrigger) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ModelScanTrigger) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *ModelScanTrigger) GetResourceType() ControlsScanResource {
	if o == nil {
		var ret ControlsScanResource
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ModelScanTrigger) GetResourceTypeOk() (*ControlsScanResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ModelScanTrigger) SetResourceType(v ControlsScanResource) {
	o.ResourceType = v
}

func (o ModelScanTrigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_id"] = o.NodeId
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	return toSerialize, nil
}

type NullableModelScanTrigger struct {
	value *ModelScanTrigger
	isSet bool
}

func (v NullableModelScanTrigger) Get() *ModelScanTrigger {
	return v.value
}

func (v *NullableModelScanTrigger) Set(val *ModelScanTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanTrigger(val *ModelScanTrigger) *NullableModelScanTrigger {
	return &NullableModelScanTrigger{value: val, isSet: true}
}

func (v NullableModelScanTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


