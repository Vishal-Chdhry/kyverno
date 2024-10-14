package test

import (
	"fmt"

	"github.com/go-git/go-billy/v5"
	"github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func getAndCompareResource(actualResources []*unstructured.Unstructured, fs billy.Filesystem, path string) (bool, error) {
	expectedResources, err := resource.GetResourceFromPath(fs, path)
	if err != nil {
		return false, fmt.Errorf("error: failed to load resource (%s)", err)
	}

	expectedResourcesMap := map[string]unstructured.Unstructured{}
	for _, expectedResource := range expectedResources {
		r := *expectedResource
		resource.FixupGenerateLabels(r)
		expectedResourcesMap[expectedResource.GetNamespace()+"/"+expectedResource.GetName()] = r
	}

	for _, actualResource := range actualResources {
		r := *actualResource
		resource.FixupGenerateLabels(r)
		equals, err := resource.Compare(r, expectedResourcesMap[r.GetNamespace()+"/"+r.GetName()], true)
		if err != nil {
			return false, fmt.Errorf("error: failed to compare resources (%s)", err)
		}
		if !equals {
			return false, nil
		}
	}
	return true, nil
}
