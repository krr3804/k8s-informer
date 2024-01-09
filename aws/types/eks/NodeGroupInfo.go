package eksresource

import clusterresource "k8s-informer/k8s/types/cluster-resource"

type NodeGroupInfo struct {
	nodeGroupName string                     `xlsx:"NODE GROUP"`
	min           string                     `xlsx:"MIN"`
	max           string                     `xlsx:"MAX"`
	desired       string                     `xlsx:"DESIRED"`
	instanceType  string                     `xlsx:"INSTANCE TYPE"`
	label         string                     `xlsx:"LABEL"`
	Nodes         []clusterresource.NodeInfo `xlsx:"CHILD"`
}

func (n *NodeGroupInfo) SetNodes(nodes []clusterresource.NodeInfo) {
	n.Nodes = nodes
}

func SetNodeGroupInfo(
	nodeGroupName string,
	min string,
	max string,
	desired string,
	instanceType string,
	label string) *NodeGroupInfo {
	nodeGroupInfo := NodeGroupInfo{}
	nodeGroupInfo.nodeGroupName = nodeGroupName
	nodeGroupInfo.min = min
	nodeGroupInfo.max = max
	nodeGroupInfo.desired = desired
	nodeGroupInfo.instanceType = instanceType
	nodeGroupInfo.label = label
	return &nodeGroupInfo
}

func (n NodeGroupInfo) GetName() string {
	return n.nodeGroupName
}
