package v1beta1

type PrometheusOperator struct {
	Deploy    bool   `json:"deploy,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}
