package common

import (
	"fmt"
	"reflect"
	"strings"

	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	kyvernov1beta1 "github.com/kyverno/kyverno/api/kyverno/v1beta1"
	"github.com/kyverno/kyverno/pkg/logging"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	pkglabels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

const (
	LabelKeyKind      = "kyverno.io/generated-by-kind"
	LabelKeyNamespace = "kyverno.io/generated-by-namespace"
	LabelKeyName      = "kyverno.io/generated-by-name"
)

type Object interface {
	GetName() string
	GetNamespace() string
	GetKind() string
	GetAPIVersion() string
}

func ManageLabels(unstr *unstructured.Unstructured, triggerResource unstructured.Unstructured, policy kyvernov1.PolicyInterface, ruleName string) {
	// add managedBY label if not defined
	labels := unstr.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}

	// handle managedBy label
	managedBy(labels)
	// handle generatedBy label
	generatedBy(labels, triggerResource)

	PolicyInfo(labels, policy, ruleName)

	TriggerInfo(labels, &triggerResource)
	// update the labels
	unstr.SetLabels(labels)
}

func MutateLabelsSet(policyKey string, trigger Object) pkglabels.Set {
	_, policyName, _ := cache.SplitMetaNamespaceKey(policyKey)

	set := pkglabels.Set{
		kyvernov1beta1.URMutatePolicyLabel: policyName,
	}
	isNil := trigger == nil || (reflect.ValueOf(trigger).Kind() == reflect.Ptr && reflect.ValueOf(trigger).IsNil())
	if !isNil {
		set[kyvernov1beta1.URMutateTriggerNameLabel] = trigger.GetName()
		set[kyvernov1beta1.URMutateTriggerNSLabel] = trigger.GetNamespace()
		set[kyvernov1beta1.URMutateTriggerKindLabel] = trigger.GetKind()
		if trigger.GetAPIVersion() != "" {
			set[kyvernov1beta1.URMutateTriggerAPIVersionLabel] = strings.ReplaceAll(trigger.GetAPIVersion(), "/", "-")
		}
	}
	return set
}

func GenerateLabelsSet(policyKey string, trigger Object) pkglabels.Set {
	_, policyName, _ := cache.SplitMetaNamespaceKey(policyKey)

	set := pkglabels.Set{
		kyvernov1beta1.URGeneratePolicyLabel: policyName,
	}
	isNil := trigger == nil || (reflect.ValueOf(trigger).Kind() == reflect.Ptr && reflect.ValueOf(trigger).IsNil())
	if !isNil {
		set[kyvernov1beta1.URGenerateResourceNameLabel] = trigger.GetName()
		set[kyvernov1beta1.URGenerateResourceNSLabel] = trigger.GetNamespace()
		set[kyvernov1beta1.URGenerateResourceKindLabel] = trigger.GetKind()
	}
	return set
}

func managedBy(labels map[string]string) {
	// ManagedBy label
	key := kyvernov1.LabelAppManagedBy
	value := kyvernov1.ValueKyvernoApp
	val, ok := labels[key]
	if ok {
		if val != value {
			logging.V(2).Info(fmt.Sprintf("resource managed by %s, kyverno wont over-ride the label", val))
			return
		}
	}
	if !ok {
		// add label
		labels[key] = value
	}
}

func generatedBy(labels map[string]string, triggerResource unstructured.Unstructured) {
	checkGeneratedBy(labels, LabelKeyKind, triggerResource.GetKind())
	checkGeneratedBy(labels, LabelKeyNamespace, triggerResource.GetNamespace())
	checkGeneratedBy(labels, LabelKeyName, triggerResource.GetName())
}

func checkGeneratedBy(labels map[string]string, key, value string) {
	value = trimByLength(value, 63)

	val, ok := labels[key]
	if ok {
		if val != value {
			logging.V(2).Info(fmt.Sprintf("kyverno wont over-ride the label %s", key))
			return
		}
	}
	if !ok {
		// add label
		labels[key] = value
	}
}

func PolicyInfo(labels map[string]string, policy kyvernov1.PolicyInterface, ruleName string) {
	labels[GeneratePolicyLabel] = policy.GetName()
	labels[GeneratePolicyNamespaceLabel] = policy.GetNamespace()
	labels[GenerateRuleLabel] = ruleName
}

func TriggerInfo(labels map[string]string, obj Object) {
	labels[GenerateTriggerAPIVersionLabel] = obj.GetAPIVersion()
	labels[GenerateTriggerKindLabel] = obj.GetKind()
	labels[GenerateTriggerNSLabel] = obj.GetNamespace()
	labels[GenerateTriggerNameLabel] = trimByLength(obj.GetName(), 63)
}

func trimByLength(value string, character int) string {
	if len(value) > character {
		return value[0:character]
	}
	return value
}
