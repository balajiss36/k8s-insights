package models

import "time"

// PodInsightsRequest represents the request for the pod insights
type PodInsightsRequest struct {
	PodName       string    `json:"pod" bson:"pod"`
	Namespace     string    `json:"namespace" bson:"namespace"`
	CPURequest    string    `json:"cpuRequest" bson:"cpuRequest"`
	MemoryRequest string    `json:"memoryRequest" bson:"memoryRequest"`
	CPULimit      string    `json:"cpuLimit" bson:"cpuLimit"`
	MemoryLimit   string    `json:"memoryLimit" bson:"memoryLimit"`
	CPUUsage      string    `json:"cpuUsage" bson:"cpuUsage"`
	MemoryUsage   string    `json:"memoryUsage" bson:"memoryUsage"`
	RequestTime   time.Time `json:"requestTime" bson:"requestTime"`
}

// PodInsightsResponse represents the response for the pod insights
type PodInsightsResponse struct {
	PodName        string `json:"podName"`
	Namespace      string `json:"namespace"`
	Recommendation string `json:"recommendation"`
}
