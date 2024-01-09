package services

import (
	"context"
	"fmt"
	configurationresource "k8s-informer/k8s/types/configuration-resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func GetConfigMapInfo(clientset *kubernetes.Clientset) []configurationresource.ConfigMapInfo {
	configmaps, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var cmInfoList []configurationresource.ConfigMapInfo

	for _, cm := range configmaps.Items {
		namespace := cm.Namespace
		name := cm.Name
		data := len(cm.Data)
		cmInfo := configurationresource.SetConfigMapInfo(namespace, name, data)

		cmInfoList = append(cmInfoList, *cmInfo)
	}

	return cmInfoList
}

func GetSecretInfo(clientset *kubernetes.Clientset) []configurationresource.SecretInfo {
	secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var secretInfoList []configurationresource.SecretInfo

	for _, secret := range secrets.Items {
		namespace := secret.Namespace
		name := secret.Name
		secretType := fmt.Sprint(secret.Type)
		data := len(secret.Data)
		secretInfo := configurationresource.SetSecretInfo(namespace, name, secretType, data)

		secretInfoList = append(secretInfoList, *secretInfo)
	}

	return secretInfoList
}
