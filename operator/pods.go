package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateCountryPod(clientset *kubernetes.Clientset, country string) error {
	pod := &v1.Pod{
		ObjectMeta: meta.ObjectMeta{
			Name: fmt.Sprintf("country-%s", country),
			Labels: map[string]string{
				"app":     "beanworld",
				"country": country,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "country-sleeper",
					Image:   "busybox", // or your own placeholder image
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}

	_, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, meta.CreateOptions{})
	return err
}
