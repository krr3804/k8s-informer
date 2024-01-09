package workload

type JobInfo struct {
	namespace   string          `xlsx:"NAMESPACE"`
	name        string          `xlsx:"NAME"`
	completions int             `xlsx:"COMPLETIONS"`
	parallelism int             `xlsx:"PARALLELISM"`
	duration    int             `xlsx:"DURATION(sec)"`
	ttl         string          `xlsx:"TTL(sec)"`
	Containers  []ContainerInfo `xlsx:"CHILD"`
}

func SetJobInfo(namespace string, name string, completions int, parallelism int, duration int, ttl string, containers []ContainerInfo) *JobInfo {
	jobInfo := JobInfo{}
	jobInfo.namespace = namespace
	jobInfo.name = name
	jobInfo.completions = completions
	jobInfo.parallelism = parallelism
	jobInfo.duration = duration
	jobInfo.ttl = ttl
	jobInfo.Containers = containers

	return &jobInfo
}
