package workload

type ContainerInfo struct {
	containerName string `xlsx:"CONTAINER_NAME"`
	image         string `xlsx:"IMAGE"`
	request       string `xlsx:"RESOURCE_REQUEST[cpu/memory]"`
	limit         string `xlsx:"RESOURCE_LIMIT[cpu/memory]"`
}

func SetContainerInfo(containerName string, image string, request string, limit string) *ContainerInfo {
	containerInfo := ContainerInfo{}
	containerInfo.containerName = containerName
	containerInfo.image = image
	containerInfo.request = request
	containerInfo.limit = limit

	return &containerInfo
}
