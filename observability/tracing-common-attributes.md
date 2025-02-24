# Common attributes

Short list of common resource and span attributes.

## Resource Attributes

## Span Attributes

| Attribute               | Type   | Description                                                                                | Examples                               |
|-------------------------|--------|--------------------------------------------------------------------------------------------|----------------------------------------|
| `cs.cluster.uid`        | string | A unique identifier for the cluster, created, deleted, or modified by the current operation. | `d1082221-078a-44ca-9ec3-8e6b5a3a51ad` |
| `opid`                  | string | The operation ID which can be generated by the service or provided by the calling entity using HTTP headers. | `d1082221-078a-44ca-9ec3-8e6b5a3a51ad` |
| `aro.correlation.id`    | string | Identifier provided by the Azure Resource Manager (ARM) which helps tracing the request flow between services. Each operation has a unique Correlation ID that aids in troubleshooting issues by correlating them with other signals across multiple services.  | `d1082221-078a-44ca-9ec3-8e6b5a3a51ad` |
| `aro.client.request.id` | string | Identifier for the azure client’s request, used to track a specific request from the client through different services. | `d1082221-078a-44ca-9ec3-8e6b5a3a51ad` |
| `aro.request.id`        | string | A unique identifier generated by the Resource Provider which can be used for tracking and troubleshooting of the request across multiple services. | `d1082221-078a-44ca-9ec3-8e6b5a3a51ad` |

