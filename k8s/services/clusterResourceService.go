package services

import (
	"context"
	clusterresource "k8s-informer/k8s/types/cluster-resource"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNodeInfo(clientset *kubernetes.Clientset) []clusterresource.NodeInfo {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var nodeInfoList []clusterresource.NodeInfo

	for _, node := range nodes.Items {
		name := node.Name
		version := node.Status.NodeInfo.KubeletVersion
		internalIP := func(addresses []corev1.NodeAddress) string {
			var result string
			for _, address := range addresses {
				if address.Type == "InternalIP" {
					result = address.Address
				}
			}
			return result
		}(node.Status.Addresses)
		osImage := node.Status.NodeInfo.OSImage
		nodegroup := node.Labels["eks.amazonaws.com/nodegroup"]

		nodeInfo := clusterresource.SetNodeInfo(name, version, internalIP, osImage, nodegroup)
		nodeInfoList = append(nodeInfoList, *nodeInfo)
	}

	return nodeInfoList
}
