package v1beta1

type LoggingOperator struct {
	Deploy    bool     `json:"deploy,omitempty"`
	Prefix    string   `json:"prefix,omitempty"`
	Namespace string   `json:"namespace,omitempty"`
	Logging   *Logging `json:"logging,omitempty"`
}

type Logging struct {
	ControlNamespace string `json:"controlNamespace,omitempty"`
}