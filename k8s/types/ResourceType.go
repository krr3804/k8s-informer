package types

import (
	cluster "k8s-informer/k8s/types/cluster-resource"
	configuration "k8s-informer/k8s/types/configuration-resource"
	networking "k8s-informer/k8s/types/network-resource"
	security "k8s-informer/k8s/types/security-resource"
	storage "k8s-informer/k8s/types/storage-resource"
	workload "k8s-informer/k8s/types/workload"
)

type ResourceType interface {
	workload.PodInfo |
		workload.DeploymentInfo |
		workload.ContainerInfo |
		workload.DaemonSetInfo |
		workload.StatefulSetInfo |
		workload.JobInfo |
		workload.CronJobInfo |
		networking.ServiceInfo |
		networking.IngressInfo |
		storage.StorageClassInfo |
		storage.PersistentVolumeClaimInfo |
		storage.PersistentVolumeInfo |
		configuration.ConfigMapInfo |
		configuration.SecretInfo |
		security.ServiceAccountInfo |
		cluster.NodeInfo
}
