package eksresource

type ClusterInfo struct {
	name    string `xlsx:"NAME"`
	arn     string `xlsx:"ARN"`
	version string `xlsx:"VERSION"`
}

func SetClusterInfo(
	name string,
	arn string,
	version string) *ClusterInfo {
	clusterInfo := ClusterInfo{}
	clusterInfo.name = name
	clusterInfo.arn = arn
	clusterInfo.version = version

	return &clusterInfo
}
