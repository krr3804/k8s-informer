package services

import (
	"context"
	securityresource "k8s-informer/k8s/types/security-resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetServiceAccountInfo(clients *kubernetes.Clientset) []securityresource.ServiceAccountInfo {

	serviceaccounts, err := clients.CoreV1().ServiceAccounts("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var saInfoList []securityresource.ServiceAccountInfo

	for _, sa := range serviceaccounts.Items {
		namespace := sa.Namespace
		name := sa.Name
		secrets := len(sa.Secrets)

		saInfo := securityresource.SetServiceAccountInfo(namespace, name, secrets)
		saInfoList = append(saInfoList, *saInfo)

	}

	return saInfoList
}
