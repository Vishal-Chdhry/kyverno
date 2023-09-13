package values

type Values struct {
	Policies           []Policy               `json:"policies"`
	GlobalValues       map[string]interface{} `json:"globalValues"`
	NamespaceSelectors []NamespaceSelector    `json:"namespaceSelector"`
	Subresources       []Subresource          `json:"subresources"`
}
