package models

import "time"

// PodInsightsRequest represents the request for the pod insights
type PodInsightsRequest struct {
	PodName       string    `json:"pod" bson:"pod,omitempty"`
	Namespace     string    `json:"namespace" bson:"namespace,omitempty"`
	CPURequest    int64     `json:"cpu_request" bson:"cpu_request,omitempty"`
	MemoryRequest int64     `json:"memory_request" bson:"memory_request,omitempty"`
	CPULimit      int64     `json:"cpu_limit" bson:"cpu_limit,omitempty"`
	MemoryLimit   int64     `json:"memory_limit" bson:"memory_limit,omitempty"`
	CPUUsage      int64     `json:"cpu_usage" bson:"cpu_usage,omitempty"`
	MemoryUsage   int64     `json:"memory_usage" bson:"memory_usage,omitempty"`
	RequestTime   time.Time `json:"request_time" bson:"request_time,omitempty"`
}

// PodInsightsResponse represents the response for the pod insights
type PodInsightsResponse struct {
	PodName        string `json:"podName"`
	Namespace      string `json:"namespace"`
	Recommendation string `json:"recommendation"`
}
