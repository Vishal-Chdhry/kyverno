package anchor

import (
	"fmt"
	"strconv"

	"github.com/go-logr/logr"
	"github.com/kyverno/kyverno/pkg/logging"
)

type resourceElementHandler = func(logr.Logger, interface{}, interface{}, interface{}, string, *AnchorKey) (string, error)

// ValidationHandler for element processes
type ValidationHandler interface {
	Handle(resourceElementHandler, map[string]interface{}, interface{}, *AnchorKey) (string, error)
}

// CreateElementHandler factory to process elements
func CreateElementHandler(element string, pattern interface{}, path string) ValidationHandler {
	anchor := Parse(element)
	if anchor == nil {
		return newDefaultHandler(element, pattern, path)
	}
	switch {
	case anchor.IsCondition():
		return newConditionAnchorHandler(element, pattern, path)
	case anchor.IsGlobal():
		return newGlobalAnchorHandler(element, pattern, path)
	case anchor.IsExistence():
		return newExistenceHandler(element, pattern, path)
	case anchor.IsEquality():
		return newEqualityHandler(element, pattern, path)
	case anchor.IsNegation():
		return newNegationHandler(element, pattern, path)
	default:
		return newDefaultHandler(element, pattern, path)
	}
}

// negationHandler provides handler for check if the tag in anchor is not defined
type negationHandler struct {
	anchor  string
	pattern interface{}
	path    string
}

// newNegationHandler returns instance of negation handler
func newNegationHandler(anchor string, pattern interface{}, path string) ValidationHandler {
	return negationHandler{
		anchor:  anchor,
		pattern: pattern,
		path:    path,
	}
}

// Handle process negation handler
func (nh negationHandler) Handle(handler resourceElementHandler, resourceMap map[string]interface{}, originPattern interface{}, ac *AnchorKey) (string, error) {
	anchorKey, _ := RemoveAnchor(nh.anchor)
	currentPath := nh.path + anchorKey + "/"
	// if anchor is present in the resource then fail
	if _, ok := resourceMap[anchorKey]; ok {
		// no need to process elements in value as key cannot be present in resource
		ac.AnchorError = newNegationAnchorError(fmt.Sprintf("%s is not allowed", currentPath))
		return currentPath, ac.AnchorError.Error()
	}
	// key is not defined in the resource
	return "", nil
}

// equalityHandler provides handler for non anchor element
type equalityHandler struct {
	anchor  string
	pattern interface{}
	path    string
}

// newEqualityHandler returens instance of equality handler
func newEqualityHandler(anchor string, pattern interface{}, path string) ValidationHandler {
	return equalityHandler{
		anchor:  anchor,
		pattern: pattern,
		path:    path,
	}
}

// Handle processed condition anchor
func (eh equalityHandler) Handle(handler resourceElementHandler, resourceMap map[string]interface{}, originPattern interface{}, ac *AnchorKey) (string, error) {
	anchorKey, _ := RemoveAnchor(eh.anchor)
	currentPath := eh.path + anchorKey + "/"
	// check if anchor is present in resource
	if value, ok := resourceMap[anchorKey]; ok {
		// validate the values of the pattern
		returnPath, err := handler(logging.GlobalLogger(), value, eh.pattern, originPattern, currentPath, ac)
		if err != nil {
			return returnPath, err
		}
		return "", nil
	}
	return "", nil
}

// defaultHandler provides handler for non anchor element
type defaultHandler struct {
	element string
	pattern interface{}
	path    string
}

// newDefaultHandler returns handler for non anchor elements
func newDefaultHandler(element string, pattern interface{}, path string) ValidationHandler {
	return defaultHandler{
		element: element,
		pattern: pattern,
		path:    path,
	}
}

// Handle process non anchor element
func (dh defaultHandler) Handle(handler resourceElementHandler, resourceMap map[string]interface{}, originPattern interface{}, ac *AnchorKey) (string, error) {
	currentPath := dh.path + dh.element + "/"
	if dh.pattern == "*" && resourceMap[dh.element] != nil {
		return "", nil
	} else if dh.pattern == "*" && resourceMap[dh.element] == nil {
		return dh.path, fmt.Errorf("%s/%s not found", dh.path, dh.element)
	} else {
		path, err := handler(logging.GlobalLogger(), resourceMap[dh.element], dh.pattern, originPattern, currentPath, ac)
		if err != nil {
			return path, err
		}
	}
	return "", nil
}

// conditionAnchorHandler provides handler for condition anchor
type conditionAnchorHandler struct {
	anchor  string
	pattern interface{}
	path    string
}

// newConditionAnchorHandler returns an instance of condition acnhor handler
func newConditionAnchorHandler(anchor string, pattern interface{}, path string) ValidationHandler {
	return conditionAnchorHandler{
		anchor:  anchor,
		pattern: pattern,
		path:    path,
	}
}

// Handle processed condition anchor
func (ch conditionAnchorHandler) Handle(handler resourceElementHandler, resourceMap map[string]interface{}, originPattern interface{}, ac *AnchorKey) (string, error) {
	anchorKey, _ := RemoveAnchor(ch.anchor)
	currentPath := ch.path + anchorKey + "/"
	// check if anchor is present in resource
	if value, ok := resourceMap[anchorKey]; ok {
		// validate the values of the pattern
		returnPath, err := handler(logging.GlobalLogger(), value, ch.pattern, originPattern, currentPath, ac)
		if err != nil {
			ac.AnchorError = newConditionalAnchorError(err.Error())
			return returnPath, ac.AnchorError.Error()
		}
		return "", nil
	} else {
		msg := "conditional anchor key doesn't exist in the resource"
		return currentPath, newConditionalAnchorError(msg).Error()
	}
}

// globalAnchorHandler provides handler for global condition anchor
type globalAnchorHandler struct {
	anchor  string
	pattern interface{}
	path    string
}

// newGlobalAnchorHandler returns an instance of condition acnhor handler
func newGlobalAnchorHandler(anchor string, pattern interface{}, path string) ValidationHandler {
	return globalAnchorHandler{
		anchor:  anchor,
		pattern: pattern,
		path:    path,
	}
}

// Handle processed global condition anchor
func (gh globalAnchorHandler) Handle(handler resourceElementHandler, resourceMap map[string]interface{}, originPattern interface{}, ac *AnchorKey) (string, error) {
	anchorKey, _ := RemoveAnchor(gh.anchor)
	currentPath := gh.path + anchorKey + "/"
	// check if anchor is present in resource
	if value, ok := resourceMap[anchorKey]; ok {
		// validate the values of the pattern
		returnPath, err := handler(logging.GlobalLogger(), value, gh.pattern, originPattern, currentPath, ac)
		if err != nil {
			ac.AnchorError = newGlobalAnchorError(err.Error())
			return returnPath, ac.AnchorError.Error()
		}
		return "", nil
	}
	return "", nil
}

// existenceHandler provides handlers to process exitence anchor handler
type existenceHandler struct {
	anchor  string
	pattern interface{}
	path    string
}

// newExistenceHandler returns existence handler
func newExistenceHandler(anchor string, pattern interface{}, path string) ValidationHandler {
	return existenceHandler{
		anchor:  anchor,
		pattern: pattern,
		path:    path,
	}
}

// Handle processes the existence anchor handler
func (eh existenceHandler) Handle(handler resourceElementHandler, resourceMap map[string]interface{}, originPattern interface{}, ac *AnchorKey) (string, error) {
	// skip is used by existence anchor to not process further if condition is not satisfied
	anchorKey, _ := RemoveAnchor(eh.anchor)
	currentPath := eh.path + anchorKey + "/"
	// check if anchor is present in resource
	if value, ok := resourceMap[anchorKey]; ok {
		// Existence anchor can only exist on resource value type of list
		switch typedResource := value.(type) {
		case []interface{}:
			typedPattern, ok := eh.pattern.([]interface{})
			if !ok {
				return currentPath, fmt.Errorf("invalid pattern type %T: Pattern has to be of list to compare against resource", eh.pattern)
			}
			// loop all item in the pattern array
			errorPath := ""
			var err error
			for _, patternMap := range typedPattern {
				typedPatternMap, ok := patternMap.(map[string]interface{})
				if !ok {
					return currentPath, fmt.Errorf("invalid pattern type %T: Pattern has to be of type map to compare against items in resource", eh.pattern)
				}
				errorPath, err = validateExistenceListResource(handler, typedResource, typedPatternMap, originPattern, currentPath, ac)
				if err != nil {
					return errorPath, err
				}
			}
			return errorPath, err
		default:
			return currentPath, fmt.Errorf("invalid resource type %T: Existence ^ () anchor can be used only on list/array type resource", value)
		}
	}
	return "", nil
}

func validateExistenceListResource(handler resourceElementHandler, resourceList []interface{}, patternMap map[string]interface{}, originPattern interface{}, path string, ac *AnchorKey) (string, error) {
	// the idea is all the element in the pattern array should be present atleast once in the resource list
	// if non satisfy then throw an error
	for i, resourceElement := range resourceList {
		currentPath := path + strconv.Itoa(i) + "/"
		_, err := handler(logging.GlobalLogger(), resourceElement, patternMap, originPattern, currentPath, ac)
		if err == nil {
			// condition is satisfied, dont check further
			return "", nil
		}
	}
	// none of the existence checks worked, so thats a failure sceanario
	return path, fmt.Errorf("existence anchor validation failed at path %s", path)
}
