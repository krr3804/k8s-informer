package services

import (
	"context"
	"k8s-informer/k8s/types"
	"k8s-informer/k8s/types/workload"
	"k8s-informer/util"
	"strconv"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

/*********
* Pod 정보
**********/
func GetPodsInfo(clientset *kubernetes.Clientset) []workload.PodInfo {

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var podInfoList []workload.PodInfo

	for _, pod := range pods.Items {
		totalContainers := len(pod.Spec.Containers)

		var readyContainers int64

		for _, container := range pod.Status.ContainerStatuses {
			if container.Ready {
				readyContainers++
			}
		}

		ready := strconv.Itoa(int(readyContainers)) + "/" + strconv.Itoa(totalContainers)

		podInfo := workload.SetPodInfo(pod.Namespace, pod.Name, ready, string(pod.Status.Phase), pod.Status.PodIP, pod.Spec.NodeName)
		podInfoList = append(podInfoList, *podInfo)
	}

	return podInfoList
}

/*******************
* Delployment 정보
********************/
func GetDeploymentsInfo(clientset *kubernetes.Clientset) []workload.DeploymentInfo {
	deploys, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var deployInfoList []workload.DeploymentInfo

	for _, deploy := range deploys.Items {

		namespace := deploy.Namespace
		name := deploy.Name
		upToDate := int(deploy.Status.UpdatedReplicas)
		ready := strconv.Itoa(int(deploy.Status.ReadyReplicas)) + "/" + strconv.Itoa(int(deploy.Status.Replicas))
		containers := getContainersInfo[workload.DeploymentInfo](deploy.Spec.Template.Spec.Containers)

		deployInfo := workload.SetDeploymentInfo(namespace, name, ready, upToDate, containers)
		deployInfoList = append(deployInfoList, *deployInfo)
	}

	return deployInfoList
}

/*******************
* StatefulSet 정보
********************/
func GetStatefulSetsInfo(clientset *kubernetes.Clientset) []workload.StatefulSetInfo {
	statefulsets, err := clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var stsInfoList []workload.StatefulSetInfo

	for _, sts := range statefulsets.Items {
		namespace := sts.Namespace
		name := sts.Name
		ready := strconv.Itoa(int(sts.Status.ReadyReplicas)) + "/" + strconv.Itoa(int(sts.Status.Replicas))
		containers := getContainersInfo[workload.StatefulSetInfo](sts.Spec.Template.Spec.Containers)

		stsInfo := workload.SetStatefulSetInfo(namespace, name, ready, containers)
		stsInfoList = append(stsInfoList, *stsInfo)
	}

	return stsInfoList
}

/*******************
* DaemonSet 정보
********************/
func GetDaemonSetsInfo(clientset *kubernetes.Clientset) []workload.DaemonSetInfo {
	daemonsets, err := clientset.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var dsInfoList []workload.DaemonSetInfo

	for _, ds := range daemonsets.Items {
		namespace := ds.Namespace
		name := ds.Name
		desired := int(ds.Status.DesiredNumberScheduled)
		current := int(ds.Status.CurrentNumberScheduled)
		ready := int(ds.Status.NumberReady)
		upToDate := int(ds.Status.UpdatedNumberScheduled)
		available := int(ds.Status.NumberAvailable)
		nodeSelector := util.MapToString(ds.Spec.Template.Spec.NodeSelector, ":")
		containers := getContainersInfo[workload.DaemonSetInfo](ds.Spec.Template.Spec.Containers)

		dsInfo := workload.SetDaemonSetInfo(namespace, name, desired, current, ready, upToDate, available, nodeSelector, containers)
		dsInfoList = append(dsInfoList, *dsInfo)
	}

	return dsInfoList
}

/*******************
* Job 정보
********************/
func GetJobsInfo(clientset *kubernetes.Clientset) []workload.JobInfo {
	jobs, err := clientset.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var jobInfoList []workload.JobInfo

	for _, job := range jobs.Items {
		namespace := job.Namespace
		name := job.Name
		completions := int(*job.Spec.Completions)
		parallelism := int(*job.Spec.Parallelism)
		duration := int(job.Status.CompletionTime.Time.Sub(job.Status.StartTime.Time).Seconds())
		ttl := util.Int32PointerToString(job.Spec.TTLSecondsAfterFinished)
		containers := getContainersInfo[workload.JobInfo](job.Spec.Template.Spec.Containers)

		jobInfo := workload.SetJobInfo(namespace, name, completions, parallelism, duration, ttl, containers)
		jobInfoList = append(jobInfoList, *jobInfo)
	}

	return jobInfoList
}

/*******************
* CronJob 정보
********************/
func GetCronJobsInfo(clientset *kubernetes.Clientset) []workload.CronJobInfo {
	cronjobs, err := clientset.BatchV1().CronJobs("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	var cronjobInfoList []workload.CronJobInfo

	for _, cronjob := range cronjobs.Items {
		namespace := cronjob.Namespace
		name := cronjob.Name
		schedule := cronjob.Spec.Schedule
		suspend := *cronjob.Spec.Suspend
		lastSuccessful := util.MetaV1TimeToString(cronjob.Status.LastSuccessfulTime)
		concurrencyPolicy := string(cronjob.Spec.ConcurrencyPolicy)
		containers := getContainersInfo[workload.CronJobInfo](cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers)

		cronjobInfo := workload.SetCronJobInfo(namespace, name, schedule, suspend, lastSuccessful, concurrencyPolicy, containers)
		cronjobInfoList = append(cronjobInfoList, *cronjobInfo)
	}

	return cronjobInfoList
}

func getContainersInfo[T types.ResourceType](containers []v1.Container) []workload.ContainerInfo {

	var containerInfoList []workload.ContainerInfo

	for _, container := range containers {
		containerName := container.Name
		image := container.Image
		request := container.Resources.Requests.Cpu().String() + "/" + container.Resources.Requests.Memory().String()
		limit := container.Resources.Limits.Cpu().String() + "/" + container.Resources.Limits.Memory().String()
		containerInfo := workload.SetContainerInfo(containerName, image, request, limit)
		containerInfoList = append(containerInfoList, *containerInfo)
	}

	return containerInfoList
}
