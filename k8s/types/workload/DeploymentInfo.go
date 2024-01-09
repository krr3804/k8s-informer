package workload

type DeploymentInfo struct {
	namespace  string          `xlsx:"NAMPESPACE"`
	name       string          `xlsx:"NAME"`
	ready      string          `xlsx:"READY"`
	upToDate   int             `xlsx:"UP-TO-DATE"`
	Containers []ContainerInfo `xlsx:"CHILD"`
}

func SetDeploymentInfo(namespace string, name string, ready string, upToDate int, containers []ContainerInfo) *DeploymentInfo {
	deploymentInfo := DeploymentInfo{}
	deploymentInfo.namespace = namespace
	deploymentInfo.name = name
	deploymentInfo.ready = ready
	deploymentInfo.upToDate = upToDate
	deploymentInfo.Containers = containers
	return &deploymentInfo
}
