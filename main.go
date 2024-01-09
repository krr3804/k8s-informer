package main

import (
	ekscontrollers "k8s-informer/aws/controllers"
	"k8s-informer/aws/services"
	k8scontrollers "k8s-informer/k8s/controllers"
	k8sservices "k8s-informer/k8s/services"
	"k8s-informer/view/input"
	"k8s-informer/view/output"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	xlsx "github.com/tealeg/xlsx/v3"
)

func main() {

	contexts, kubeconfig := k8sservices.GetContexts()

	provider, context, filename := input.ScreenPrompt(contexts)

	filename = filename + ".xlsx"
	file := xlsx.NewFile()

	indexSheet, _ := file.AddSheet("INDEX")

	indexDataMap := navigateByProvider(provider, file, context, kubeconfig)

	output.WriteIndexSheet(indexSheet, indexDataMap)

	err := file.Save(filename)
	if err != nil {
		panic(err)
	}
}

func navigateByProvider(provider string, file *xlsx.File, context string, kubeconfig string) linkedhashmap.Map {
	clientset := k8sservices.GetClientset(context, kubeconfig)

	indexDataMap := linkedhashmap.New()

	switch provider {
	case "EKS[AWS]":
		clusterName := k8sservices.GetCurrentCluster(context, kubeconfig)
		eksclient := services.GetEksClient(clusterName)

		ekscontrollers.GetCluster(file, eksclient, clusterName)
		ekscontrollers.GetNodeGroups(file, eksclient, clusterName, clientset)

		k8scontrollers.GenerateK8SResources(file, clientset)

		indexDataMap.Put("EKS", []string{"CLUSTER", "NODE GROUP"})
	default:
		indexDataMap.Put("CLUSTER", []string{"NODE"})
		k8scontrollers.GenerateNodeInfo(file, clientset)
		k8scontrollers.GenerateK8SResources(file, clientset)
	}

	indexDataMap.Put("WORKLOAD", []string{"POD", "DEPLOYMENT", "STATEFULSET", "DAEMONSET", "JOB", "CRONJOB"})
	indexDataMap.Put("NETWORKING", []string{"SERVICE", "INGRESS"})
	indexDataMap.Put("STORAGE", []string{"STORAGE CLASS", "PERSISTENT VOLUME CLAIM", "PESISTENT VOLUME"})
	indexDataMap.Put("CONFIGURATION", []string{"CONFIGMAP", "SECRET"})
	indexDataMap.Put("SECURITY", []string{"SERVICE ACCOUNT"})

	return *indexDataMap
}
