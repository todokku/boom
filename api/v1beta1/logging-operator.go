package v1beta1

type LoggingOperator struct {
	Deploy     bool         `json:"deploy,omitempty"`
	FluentdPVC *StorageSpec `json:"fluentdStorage,omitempty" yaml:"fluentdStorage,omitempty"`
}
