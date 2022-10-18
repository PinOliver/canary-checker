/*
 * Canary Checker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1..1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Check struct {
	CanaryId string `json:"canary_id,omitempty"`
	CanaryName string `json:"canary_name,omitempty"`
	CheckStatuses []CheckStatus `json:"checkStatuses,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayType string `json:"displayType,omitempty"`
	Icon string `json:"icon,omitempty"`
	Id string `json:"id,omitempty"`
	Labels *interface{} `json:"labels,omitempty"`
	LastRuntime string `json:"lastRuntime,omitempty"`
	Latency *Latency `json:"latency,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	NextRuntime string `json:"nextRuntime,omitempty"`
	Owner string `json:"owner,omitempty"`
	Severity string `json:"severity,omitempty"`
	Status string `json:"status,omitempty"`
	Type_ string `json:"type,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Uptime *Uptime `json:"uptime,omitempty"`
}