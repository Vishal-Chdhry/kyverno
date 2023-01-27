package anchor

// GetAnchorsResourcesFromMap returns map of anchors
func GetAnchorsResourcesFromMap(patternMap map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	anchors := map[string]interface{}{}
	resources := map[string]interface{}{}
	for key, value := range patternMap {
		anchor := Parse(key)
		if anchor.IsCondition() || anchor.IsExistence() || anchor.IsEquality() || anchor.IsNegation() {
			anchors[key] = value
			continue
		}
		resources[key] = value
	}
	return anchors, resources
}
