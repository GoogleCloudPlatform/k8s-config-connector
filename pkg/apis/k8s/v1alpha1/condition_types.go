package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

const (
	ReadyConditionType = "Ready"
)

type Condition struct {
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last
	// transition.
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown.
	Status v1.ConditionStatus `json:"status,omitempty"`

	// Type is the type of the condition.
	Type string `json:"type,omitempty"`
}

type ResourceRef struct {
	/* The external name of the referenced resource */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}
