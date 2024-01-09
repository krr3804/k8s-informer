package workload

type StatefulSetInfo struct {
	namespace  string          `xlsx:"NAMESPACE"`
	name       string          `xlsx:"NAME"`
	ready      string          `xlsx:"READY"`
	Containers []ContainerInfo `xlsx:"CHILD"`
}

func SetStatefulSetInfo(namespace string, name string, ready string, container []ContainerInfo) *StatefulSetInfo {
	statefulSetInfo := StatefulSetInfo{}
	statefulSetInfo.namespace = namespace
	statefulSetInfo.name = name
	statefulSetInfo.ready = ready
	statefulSetInfo.Containers = container

	return &statefulSetInfo
}
