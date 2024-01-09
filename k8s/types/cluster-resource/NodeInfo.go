package clusterresource

type NodeInfo struct {
	name       string `xlsx:"NAME"`
	version    string `xlsx:"VERSION"`
	internalIP string `xlsx:"INTERNAL-IP"`
	osImage    string `xlsx:"OS-IMAGE"`
	nodegroup  string
}

func SetNodeInfo(
	name string,
	version string,
	internalIP string,
	osImage string,
	nodegroup string) *NodeInfo {
	nodeInfo := NodeInfo{}
	nodeInfo.name = name
	nodeInfo.version = version
	nodeInfo.internalIP = internalIP
	nodeInfo.osImage = osImage
	nodeInfo.nodegroup = nodegroup
	return &nodeInfo
}

func (n NodeInfo) GetNodeGroup() string {
	return n.nodegroup
}
