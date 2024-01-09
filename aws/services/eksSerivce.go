package services

import (
	"context"
	eksresource "k8s-informer/aws/types/eks"
	clusterresource "k8s-informer/k8s/types/cluster-resource"
	"k8s-informer/util"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func GetEksClient(clusterName string) *eks.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	client := eks.NewFromConfig(cfg, func(o *eks.Options) {
	})

	return client
}

func GetClusterInfo(client *eks.Client, clusterName string) eksresource.ClusterInfo {

	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}

	cluster, err := client.DescribeCluster(context.TODO(), input)
	if err != nil {
		panic(err.Error())
	}

	name := util.StringPointerToString(cluster.Cluster.Name)
	arn := util.StringPointerToString(cluster.Cluster.Arn)
	version := util.StringPointerToString(cluster.Cluster.Version)

	clusterInfo := eksresource.SetClusterInfo(name, arn, version)

	return *clusterInfo
}

func GetNodeGroupInfoList(client *eks.Client, clusterName string) []eksresource.NodeGroupInfo {

	nodegroups, err := client.ListNodegroups(context.TODO(),
		&eks.ListNodegroupsInput{
			ClusterName: aws.String(clusterName),
		})
	if err != nil {
		panic(err.Error())
	}

	var nodeGroupInfoList []eksresource.NodeGroupInfo

	for _, nodegroupName := range nodegroups.Nodegroups {
		ng, err := client.DescribeNodegroup(context.TODO(),
			&eks.DescribeNodegroupInput{
				ClusterName:   aws.String(clusterName),
				NodegroupName: aws.String(nodegroupName)})

		if err != nil {
			panic(err.Error())
		}
		name := util.StringPointerToString(ng.Nodegroup.NodegroupName)
		min := util.Int32PointerToString(ng.Nodegroup.ScalingConfig.MinSize)
		max := util.Int32PointerToString(ng.Nodegroup.ScalingConfig.MaxSize)
		desired := util.Int32PointerToString(ng.Nodegroup.ScalingConfig.DesiredSize)
		instanceType := strings.Join(ng.Nodegroup.InstanceTypes, ",")
		label := util.MapToString(ng.Nodegroup.Labels, "=")

		nodegroupInfo := eksresource.SetNodeGroupInfo(name, min, max, desired, instanceType, label)

		nodeGroupInfoList = append(nodeGroupInfoList, *nodegroupInfo)
	}

	return nodeGroupInfoList
}

func PutNodesIntoNodeGroup(nodegroups []eksresource.NodeGroupInfo, nodes []clusterresource.NodeInfo) []eksresource.NodeGroupInfo {

	var ngs []eksresource.NodeGroupInfo

	for _, nodegroup := range nodegroups {
		var nodeInfoList []clusterresource.NodeInfo

		for _, node := range nodes {
			if nodegroup.GetName() == node.GetNodeGroup() {
				nodeInfoList = append(nodeInfoList, node)
			}
		}

		nodegroup.SetNodes(nodeInfoList)
		ngs = append(ngs, nodegroup)
	}

	return ngs
}
