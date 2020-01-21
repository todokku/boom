# CRD boom.caos.ch/v1beta1

## Structure

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `kubeVersion`                      | Version of the kuberentes version of the cluster                                |                                   |
| `prometheus-operator`              | Spec for the Prometheus-Operator                                                |                                   |
| `logging-operator`                 | Spec for the Banzaicloud Logging-Operator                                       |                                   |
| `prometheus-node-exporter`         | Spec for the Prometheus-NOde-Exporter                                           |                                   |
| `grafana`                          | Spec for the Grafana                                                            |                                   |
| `ambassador`                       | Spec for the Ambassador                                                         |                                   |
| `kube-state-metrics`               | Spec for the Kube-State-Metrics                                                 |                                   |
| `argocd`                           | Spec for the Argo-CD                                                            |                                   |
| `prometheus`                       | Spec for the Prometheus instance                                                |                                   |
| `loki`                             | Spec for the Loki instance                                                      |                                   |

### Prometheus-Operator

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |

### Logging-Operator

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |

### Prometheus-Node-Exporter

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |

### Grafana

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |
| `admin`                            | Spec for the definition of the admin account                                    | nil                               |
| `admin.existingSecret`             | Name of the secret which contains the admin account                             |                                   |
| `admin.userKey`                    | Key of the username in the secret                                               |                                   |
| `admin.PasswordKey`                | Key of the password in the secret                                               |                                   |
| `admin`                            | Spec for the definition of the admin account                                    |                                   |
| `datasources`                      | Spec for additional datasources                                                 | nil                               |
| `datasources.name`                 | Name of the datasource                                                          |                                   |
| `datasources.type`                 | Type of the datasource (for example prometheus)                                 |                                   |
| `datasources.url`                  | URL to the datasource                                                           |                                   |
| `datasources.access`               | Access defintion of the datasource                                              |                                   |
| `datasources.isDefault`            | Boolean if datasource should be used as default                                 |                                   |
| `dashboardproviders`               | Spec for additional Dashboardproviders                                          | nil                               |
| `dashboardproviders.configMaps`    | ConfigMaps in which the dashboards are stored                                   |                                   |
| `dashboardproviders.folder`        | Local folder in which the dashboards are mounted                                |                                   |

### Ambassador

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |
| `replicaCount`                     | Number of replicas used for deployment                                          | 1                                 |
| `loadBalancerIP`                   | Used IP for loadbalancing for ambassador                                        | nil                               |

### Kube-State-Metrics

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |
| `replicaCount`                     | Number of replicas used for deployment                                          | 1                                 |

### Argo-CD

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |

### Prometheus

When the metrics spec is nil all metrics will get scraped.

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |
| `metrics`                          | Spec to define which metrics should get scraped                                 | nil                               |
| `metrics.ambassador`               | Bool if metrics should get scraped                                              | false                             |
| `metrics.argocd`                   | Bool if metrics should get scraped                                              | false                             |
| `metrics.kube-state-metrics`       | Bool if metrics should get scraped                                              | false                             |
| `metrics.prometheus-node-exporter` | Bool if metrics should get scraped                                              | false                             |
| `metrics.api-server`               | Bool if metrics should get scraped                                              | false                             |
| `metrics.prometheus-operator`      | Bool if metrics should get scraped                                              | false                             |
| `metrics.logging-operator`         | Bool if metrics should get scraped                                              | false                             |
| `metrics.loki`                     | Bool if metrics should get scraped                                              | false                             |
| `storage`                          | Spec to define how the persistency should be handled                            | nil                               |
| `storage.size`                     | Defined size of the PVC                                                         |                                   |
| `storage.storageClass`             | Storageclass used by the PVC                                                    |                                   |
| `storage.accessModes`              | Accessmodes used by the PVC                                                     |                                   |

### Loki

When the logs spec is nil all logs will get persisted in loki.

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `deploy`                           | Flag if tool should be deployed                                                 | false                             |
| `logs`                             | Spec to define which logs will get persisted                                    | nil                               |
| `logs.ambassador`                  | Bool if logs will get persisted                                                 | false                             |
| `logs.argocd`                      | Bool if logs will get persisted                                                 | false                             |
| `logs.kube-state-metrics`          | Bool if logs will get persisted                                                 | false                             |
| `logs.prometheus-node-exporter`    | Bool if logs will get persisted                                                 | false                             |
| `logs.grafana`                     | Bool if logs will get persisted                                                 | false                             |
| `logs.prometheus-operator`         | Bool if logs will get persisted                                                 | false                             |
| `logs.logging-operator`            | Bool if logs will get persisted                                                 | false                             |
| `logs.loki`                        | Bool if logs will get persisted                                                 | false                             |
| `storage`                          | Spec to define how the persistency should be handled                            | nil                               |
| `storage.size`                     | Defined size of the PVC                                                         |                                   |
| `storage.storageClass`             | Storageclass used by the PVC                                                    |                                   |
| `storage.accessModes`              | Accessmodes used by the PVC                                                     |                                   |