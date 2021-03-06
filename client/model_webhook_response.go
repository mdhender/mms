/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type WebhookResponse struct {
	Messages []string                 `json:"messages,omitempty"`
	Rejected []Rejection              `json:"rejected,omitempty"`
	Webhooks []map[string]interface{} `json:"webhooks,omitempty"`
}
