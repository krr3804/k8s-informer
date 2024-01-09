package workload

type PodInfo struct {
	namespace string `xlsx:"NAMESPACE"`
	name      string `xlsx:"NAME"`
	ready     string `xlsx:"READY"`
	status    string `xlsx:"STATUS"`
	ip        string `xlsx:"IP"`
	node      string `xlsx:"NODE"`
}

func SetPodInfo(namespace string, name string, ready string, status string, ip string, node string) *PodInfo {
	podInfo := PodInfo{}
	podInfo.namespace = namespace
	podInfo.name = name
	podInfo.ready = ready
	podInfo.status = status
	podInfo.ip = ip
	podInfo.node = node

	return &podInfo
}
