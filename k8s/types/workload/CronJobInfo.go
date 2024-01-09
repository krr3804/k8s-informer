package workload

type CronJobInfo struct {
	namespace         string          `xlsx:"NAMESPACE"`
	name              string          `xlsx:"NAME"`
	schedule          string          `xlsx:"SCHEDULE"`
	suspend           bool            `xlsx:"SUSPEND"`
	lastSuccessful    string          `xlsx:"LAST SUCCESSFUL TIME"`
	concurrencyPolicy string          `xlsx:"CONCURRENCY POLICY"`
	Containers        []ContainerInfo `xlsx:"CHILD"`
}

func SetCronJobInfo(namespace string, name string, schedule string, suspend bool, lastSuccessful string, concurrencyPolicy string, containers []ContainerInfo) *CronJobInfo {
	cronjobInfo := CronJobInfo{}
	cronjobInfo.namespace = namespace
	cronjobInfo.name = name
	cronjobInfo.schedule = schedule
	cronjobInfo.suspend = suspend
	cronjobInfo.lastSuccessful = lastSuccessful
	cronjobInfo.concurrencyPolicy = concurrencyPolicy
	cronjobInfo.Containers = containers

	return &cronjobInfo
}
