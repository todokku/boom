package logs

import (
	toolsetsv1beta1 "github.com/caos/boom/api/v1beta1"
	amlogs "github.com/caos/boom/internal/bundle/application/applications/ambassador/logs"
	aglogs "github.com/caos/boom/internal/bundle/application/applications/argocd/logs"
	glogs "github.com/caos/boom/internal/bundle/application/applications/grafana/logs"
	ksmlogs "github.com/caos/boom/internal/bundle/application/applications/kubestatemetrics/logs"
	"github.com/caos/boom/internal/bundle/application/applications/loggingoperator/logging"
	pnelogs "github.com/caos/boom/internal/bundle/application/applications/prometheusnodeexporter/logs"
	pologs "github.com/caos/boom/internal/bundle/application/applications/prometheusoperator/logs"
)

func GetAllResources(toolsetCRDSpec *toolsetsv1beta1.ToolsetSpec) []interface{} {

	// output to loki
	outputNames, outputs := getOutputs()

	// add flows for each application
	flows := getAllFlows(toolsetCRDSpec, outputNames)

	ret := make([]interface{}, 0)
	if len(flows) > 0 {
		for _, flow := range flows {
			ret = append(ret, flow)
		}
		for _, output := range outputs {
			ret = append(ret, output)
		}

		//logging resource so that fluentd and fluentbit are deployed
		ret = append(ret, getLogging())
	}

	return ret
}

func getLogging() *logging.Logging {

	conf := &logging.Config{
		Name:             "logging",
		Namespace:        "caos-system",
		ControlNamespace: "caos-system",
	}

	return logging.New(conf)
}

func getAllFlows(toolsetCRDSpec *toolsetsv1beta1.ToolsetSpec, outputNames []string) []*logging.Flow {

	flows := make([]*logging.Flow, 0)
	if toolsetCRDSpec.Ambassador != nil && toolsetCRDSpec.Ambassador.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.Ambassador) {
		flows = append(flows, logging.NewFlow(amlogs.GetFlow(outputNames)))
	}

	if toolsetCRDSpec.Grafana != nil && toolsetCRDSpec.Grafana.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.Grafana) {
		flows = append(flows, logging.NewFlow(glogs.GetFlow(outputNames)))
	}

	if toolsetCRDSpec.PrometheusOperator != nil && toolsetCRDSpec.PrometheusOperator.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.PrometheusOperator) {
		flows = append(flows, logging.NewFlow(pologs.GetFlow(outputNames)))
	}

	if toolsetCRDSpec.PrometheusNodeExporter != nil && toolsetCRDSpec.PrometheusNodeExporter.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.PrometheusNodeExporter) {
		flows = append(flows, logging.NewFlow(pnelogs.GetFlow(outputNames)))
	}

	if toolsetCRDSpec.KubeStateMetrics != nil && toolsetCRDSpec.KubeStateMetrics.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.KubeStateMetrics) {
		flows = append(flows, logging.NewFlow(ksmlogs.GetFlow(outputNames)))
	}

	if toolsetCRDSpec.Argocd != nil && toolsetCRDSpec.Argocd.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.Argocd) {
		flows = append(flows, logging.NewFlow(aglogs.GetFlow(outputNames)))
	}

	if toolsetCRDSpec.Loki != nil && toolsetCRDSpec.Loki.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.Loki) {
		flows = append(flows, logging.NewFlow(getLokiFlow(outputNames)))
	}

	if toolsetCRDSpec.Prometheus != nil && toolsetCRDSpec.Prometheus.Deploy &&
		(toolsetCRDSpec.Loki.Logs == nil || toolsetCRDSpec.Loki.Logs.Prometheus) {
		flows = append(flows, logging.NewFlow(getLokiFlow(outputNames)))
	}

	return flows
}

func getLokiFlow(outputs []string) *logging.FlowConfig {
	lables := map[string]string{"release": "loki", "app": "loki"}

	return &logging.FlowConfig{
		Name:         "flow-loki",
		Namespace:    "caos-system",
		SelectLabels: lables,
		Outputs:      outputs,
		ParserType:   "none",
	}
}

func getOutputs() ([]string, []*logging.Output) {
	conf := &logging.ConfigOutput{
		Name:      "output-loki",
		Namespace: "caos-system",
		URL:       "http://loki.caos-system:3100",
	}

	outputs := make([]*logging.Output, 0)
	outputs = append(outputs, logging.NewOutput(conf))
	outputNames := make([]string, 0)
	outputNames = append(outputNames, conf.Name)

	return outputNames, outputs
}
