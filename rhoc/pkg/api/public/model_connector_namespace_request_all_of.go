/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public
// ConnectorNamespaceRequestAllOf struct for ConnectorNamespaceRequestAllOf
type ConnectorNamespaceRequestAllOf struct {
	ClusterId string `json:"cluster_id,omitempty"`
	Kind ConnectorNamespaceTenantKind `json:"kind,omitempty"`
}
