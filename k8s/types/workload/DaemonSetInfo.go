package workload

type DaemonSetInfo struct {
	namespace    string          `xlsx:"NAMESPACE"`
	name         string          `xlsx:"NAME"`
	desired      int             `xlsx:"DESIRED"`
	current      int             `xlsx:"CURRENT"`
	ready        int             `xlsx:"READY"`
	upToDate     int             `xlsx:"UP-TO-DATE"`
	available    int             `xlsx:"AVAILABLE"`
	nodeSelector string          `xlsx:"NODE SELECTOR"`
	Containers   []ContainerInfo `xlsx:"CHILD"`
}

func SetDaemonSetInfo(
	namespace string,
	name string,
	desired int,
	current int,
	ready int,
	upToDate int,
	available int,
	nodeSelector string,
	containers []ContainerInfo) *DaemonSetInfo {

	daemonSetInfo := DaemonSetInfo{}
	daemonSetInfo.namespace = namespace
	daemonSetInfo.name = name
	daemonSetInfo.desired = desired
	daemonSetInfo.current = current
	daemonSetInfo.ready = ready
	daemonSetInfo.upToDate = upToDate
	daemonSetInfo.available = available
	daemonSetInfo.nodeSelector = nodeSelector
	daemonSetInfo.Containers = containers
	return &daemonSetInfo
}
