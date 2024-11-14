package models

import "time"

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

type PodInsightsResponse struct {
	PodName        string `json:"podName"`
	Namespace      string `json:"namespace"`
	Recommendation string `json:"recommendation"`
}
