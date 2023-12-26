/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2

// AnyAllConditionsApplyConfiguration represents an declarative configuration of the AnyAllConditions type for use
// with apply.
type AnyAllConditionsApplyConfiguration struct {
	AnyConditions []ConditionApplyConfiguration `json:"any,omitempty"`
	AllConditions []ConditionApplyConfiguration `json:"all,omitempty"`
}

// AnyAllConditionsApplyConfiguration constructs an declarative configuration of the AnyAllConditions type for use with
// apply.
func AnyAllConditions() *AnyAllConditionsApplyConfiguration {
	return &AnyAllConditionsApplyConfiguration{}
}

// WithAnyConditions adds the given value to the AnyConditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AnyConditions field.
func (b *AnyAllConditionsApplyConfiguration) WithAnyConditions(values ...*ConditionApplyConfiguration) *AnyAllConditionsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAnyConditions")
		}
		b.AnyConditions = append(b.AnyConditions, *values[i])
	}
	return b
}

// WithAllConditions adds the given value to the AllConditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllConditions field.
func (b *AnyAllConditionsApplyConfiguration) WithAllConditions(values ...*ConditionApplyConfiguration) *AnyAllConditionsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllConditions")
		}
		b.AllConditions = append(b.AllConditions, *values[i])
	}
	return b
}
