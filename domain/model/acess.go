package model

type AccessModule struct {
	ModuleName string   `json:"module_name" bson:"module_name"`
	Actions    []string `json:"module_actions,omitempty" bson:"module_actions,omitempty"`
}
