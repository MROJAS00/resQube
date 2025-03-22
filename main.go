package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Recommendation struct to store CPU & memory suggestions
type Recommendation struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

// Default recommendations based on language and workload
var recommendations = map[string]map[string]Recommendation{
	"java": {
		"low":    {CPU: "250m", Memory: "512Mi"},
		"medium": {CPU: "500m", Memory: "1Gi"},
		"high":   {CPU: "1", Memory: "2Gi"},
	},
	"python": {
		"low":    {CPU: "200m", Memory: "256Mi"},
		"medium": {CPU: "400m", Memory: "512Mi"},
		"high":   {CPU: "800m", Memory: "1Gi"},
	},
	"javascript": {
		"low":    {CPU: "150m", Memory: "256Mi"},
		"medium": {CPU: "300m", Memory: "512Mi"},
		"high":   {CPU: "600m", Memory: "1Gi"},
	},
}

// Generate YAML for a Kubernetes Deployment
func generateYAML(language, workload string) string {
	rec, exists := recommendations[language][workload]
	if !exists {
		return "Error: No recommendation available"
	}

	deployment := map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata":   map[string]string{"name": "example-app"},
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{
							"name": "app-container",
							"resources": map[string]interface{}{
								"requests": map[string]string{"cpu": rec.CPU, "memory": rec.Memory},
								"limits":   map[string]string{"cpu": rec.CPU, "memory": rec.Memory},
							},
						},
					},
				},
			},
		},
	}
	yamlData, _ := yaml.Marshal(deployment)
	return string(yamlData)
}

// Cobra CLI command
var rootCmd = &cobra.Command{
	Use:   "kubesizer",
	Short: "Recommend CPU & memory for Kubernetes pods",
	Run: func(cmd *cobra.Command, args []string) {
		language, _ := cmd.Flags().GetString("language")
		workload, _ := cmd.Flags().GetString("workload")

		if language == "" || workload == "" {
			fmt.Println("Usage: kubesizer --language=java --workload=medium")
			return
		}

		yamlOutput := generateYAML(language, workload)
		fmt.Println(yamlOutput)
	},
}

func main() {
	rootCmd.Flags().StringP("language", "l", "", "Programming language (java, python, javascript)")
	rootCmd.Flags().StringP("workload", "w", "", "Workload type (low, medium, high)")
	rootCmd.MarkFlagRequired("language")
	rootCmd.MarkFlagRequired("workload")
	rootCmd.Execute()
}
