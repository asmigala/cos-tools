/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public
// ConnectorNamespaceEvalRequest An evaluation connector namespace create request
type ConnectorNamespaceEvalRequest struct {
	Name string `json:"name"`
	Annotations []ConnectorNamespaceRequestMetaAnnotations `json:"annotations,omitempty"`
}
