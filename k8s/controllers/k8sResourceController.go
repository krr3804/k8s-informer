package controllers

import (
	"k8s-informer/k8s/services"
	clusterresource "k8s-informer/k8s/types/cluster-resource"
	configurationresource "k8s-informer/k8s/types/configuration-resource"
	networkresource "k8s-informer/k8s/types/network-resource"
	securityresource "k8s-informer/k8s/types/security-resource"
	storageresource "k8s-informer/k8s/types/storage-resource"
	"k8s-informer/k8s/types/workload"
	"k8s-informer/view/output"
	"reflect"

	xlsx "github.com/tealeg/xlsx/v3"
	"k8s.io/client-go/kubernetes"
)

func GenerateNodeInfo(file *xlsx.File, clientset *kubernetes.Clientset) {
	nodeInfoList := services.GetNodeInfo(clientset)
	output.WriteResourceData(file, "NODE", nodeInfoList, reflect.TypeOf(clusterresource.NodeInfo{}))
}

func GenerateK8SResources(file *xlsx.File, clientset *kubernetes.Clientset) {

	createWorkloadSheets(file, clientset)
	createNetworkResourceSheets(file, clientset)
	createStorageResourceSheets(file, clientset)
	createConfigurationResourceSheets(file, clientset)
	createSecurityResourceSheets(file, clientset)
}

func createWorkloadSheets(file *xlsx.File, clientset *kubernetes.Clientset) {

	podInfoList := services.GetPodsInfo(clientset)
	deployInfoList := services.GetDeploymentsInfo(clientset)
	statefulsetInfoList := services.GetStatefulSetsInfo(clientset)
	daemonsetInfoList := services.GetDaemonSetsInfo(clientset)
	jobInfoList := services.GetJobsInfo(clientset)
	cronjobInfoList := services.GetCronJobsInfo(clientset)

	output.WriteResourceData(file, "POD", podInfoList, reflect.TypeOf(workload.PodInfo{}))
	output.WriteResourceData(file, "DEPLOYMENT", deployInfoList, reflect.TypeOf(workload.DeploymentInfo{}))
	output.WriteResourceData(file, "STATEFULSET", statefulsetInfoList, reflect.TypeOf(workload.StatefulSetInfo{}))
	output.WriteResourceData(file, "DAEMONSET", daemonsetInfoList, reflect.TypeOf(workload.DaemonSetInfo{}))
	output.WriteResourceData(file, "JOB", jobInfoList, reflect.TypeOf(workload.JobInfo{}))
	output.WriteResourceData(file, "CRONJOB", cronjobInfoList, reflect.TypeOf(workload.CronJobInfo{}))
}

func createNetworkResourceSheets(file *xlsx.File, clientset *kubernetes.Clientset) {
	serviceInfoList := services.GetServiceInfo(clientset)
	ingressInfoList := services.GetIngressInfo(clientset)
	output.WriteResourceData(file, "SERVICE", serviceInfoList, reflect.TypeOf(networkresource.ServiceInfo{}))
	output.WriteResourceData(file, "INGRESS", ingressInfoList, reflect.TypeOf(networkresource.IngressInfo{}))
}

func createStorageResourceSheets(file *xlsx.File, clientset *kubernetes.Clientset) {
	scInfoList := services.GetStorageClassInfo(clientset)
	pvcInfoList := services.GetPersistentVolumeClaimInfo(clientset)
	pvInfoList := services.GetPersistentVolumeInfo(clientset)

	output.WriteResourceData(file, "STORAGE CLASS", scInfoList, reflect.TypeOf(storageresource.StorageClassInfo{}))
	output.WriteResourceData(file, "PERSISTENT VOLUME CLAIM", pvcInfoList, reflect.TypeOf(storageresource.PersistentVolumeClaimInfo{}))
	output.WriteResourceData(file, "PESISTENT VOLUME", pvInfoList, reflect.TypeOf(storageresource.PersistentVolumeInfo{}))
}

func createConfigurationResourceSheets(file *xlsx.File, clientset *kubernetes.Clientset) {
	cmInfoList := services.GetConfigMapInfo(clientset)
	secretInfoList := services.GetSecretInfo(clientset)

	output.WriteResourceData(file, "CONFIGMAP", cmInfoList, reflect.TypeOf(configurationresource.ConfigMapInfo{}))
	output.WriteResourceData(file, "SECRET", secretInfoList, reflect.TypeOf(configurationresource.SecretInfo{}))
}

func createSecurityResourceSheets(file *xlsx.File, clientset *kubernetes.Clientset) {
	saInfoList := services.GetServiceAccountInfo(clientset)

	output.WriteResourceData(file, "SERVICE ACCOUNT", saInfoList, reflect.TypeOf(securityresource.ServiceAccountInfo{}))
}
