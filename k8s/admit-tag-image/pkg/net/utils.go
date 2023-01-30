package net

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

func hasIgnoreNotLatestImageLabels(labels map[string]string) bool {
	_, ok := labels[ignoreNotLastImagePolicy]
	return ok
}

func haveSomeImageLatestTag(containers []corev1.Container) (string, bool) {
	for _, container := range containers {
		if strings.Contains(container.Image, ":latest") || !strings.Contains(container.Image, ":") {
			return fmt.Sprintf("Container %s contains image with latest tag", container.Name), true
		}
	}
	return "Containers do not contain image with latest tag", false
}
