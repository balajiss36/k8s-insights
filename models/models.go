package models

import "time"

// PodInsightsRequest represents the request for the pod insights
type PodInsightsRequest struct {
	PodName       string    `json:"podName"`
	Namespace     string    `json:"namespace"`
	CPURequest    string    `json:"cpuRequest"`
	MemoryRequest string    `json:"memoryRequest"`
	CPULimit      string    `json:"cpuLimit"`
	MemoryLimit   string    `json:"memoryLimit"`
	ContainerName string    `json:"containerName"`
	CPUUsage      string    `json:"cpuUsage"`
	MemoryUsage   string    `json:"memoryUsage"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// PodInsightsResponse represents the response for the pod insights
type PodInsightsResponse struct {
	PodName        string `json:"podName"`
	Namespace      string `json:"namespace"`
	Recommendation string `json:"recommendation"`
}
