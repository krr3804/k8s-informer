package controllers

import (
	awsservices "k8s-informer/aws/services"
	eksresource "k8s-informer/aws/types/eks"
	k8sservices "k8s-informer/k8s/services"
	"k8s-informer/view/output"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	xlsx "github.com/tealeg/xlsx/v3"
	"k8s.io/client-go/kubernetes"
)

func GetCluster(file *xlsx.File, client *eks.Client, clusterName string) {
	clusterInfo := awsservices.GetClusterInfo(client, clusterName)

	output.WriteResourceData(file, "CLUSTER", []eksresource.ClusterInfo{clusterInfo}, reflect.TypeOf(clusterInfo))
}

func GetNodeGroups(file *xlsx.File, client *eks.Client, clusterName string, clientset *kubernetes.Clientset) {

	initialNodegroupInfoList := awsservices.GetNodeGroupInfoList(client, clusterName)
	nodeInfoList := k8sservices.GetNodeInfo(clientset)

	nodeGroupInfoList := awsservices.PutNodesIntoNodeGroup(initialNodegroupInfoList, nodeInfoList)

	output.WriteResourceData(file, "NODE GROUP", nodeGroupInfoList, reflect.TypeOf(eksresource.NodeGroupInfo{}))
}
