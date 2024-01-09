package services

import (
	"context"
	"fmt"
	storageresource "k8s-informer/k8s/types/storage-resource"
	"k8s-informer/util"
	"strings"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func GetStorageClassInfo(clientset *kubernetes.Clientset) []storageresource.StorageClassInfo {

	storageclasses, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var scInfoList []storageresource.StorageClassInfo

	for _, sc := range storageclasses.Items {
		name := sc.Name
		provisioner := sc.Provisioner
		reclaimPolicy := func(rp *corev1.PersistentVolumeReclaimPolicy) string {
			if rp == nil {
				return "None"
			}
			return fmt.Sprint(*rp)
		}(sc.ReclaimPolicy)
		volumeBindingMode := func(vbm *v1.VolumeBindingMode) string {
			if vbm == nil {
				return "None"
			}
			return fmt.Sprint(*vbm)
		}(sc.VolumeBindingMode)
		allowVolumeExpansion := util.BoolPointerToBool(sc.AllowVolumeExpansion, false)

		scInfo := storageresource.SetStorageClassInfo(name, provisioner, reclaimPolicy, volumeBindingMode, allowVolumeExpansion)
		scInfoList = append(scInfoList, *scInfo)
	}

	return scInfoList
}

func GetPersistentVolumeClaimInfo(clientset *kubernetes.Clientset) []storageresource.PersistentVolumeClaimInfo {
	persistentvolumeclaims, err := clientset.CoreV1().PersistentVolumeClaims("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var pvcInfoList []storageresource.PersistentVolumeClaimInfo

	for _, pvc := range persistentvolumeclaims.Items {
		namespace := pvc.Namespace
		name := pvc.Name
		status := fmt.Sprint(pvc.Status.Phase)
		volume := pvc.Spec.VolumeName
		capacity := pvc.Status.Capacity.Storage().String()
		accessMode := persistentVolumeAccessModesToString(pvc.Status.AccessModes)
		storageClass := util.StringPointerToString(pvc.Spec.StorageClassName)
		volumeMode := volumeModeToString(pvc.Spec.VolumeMode)
		pvcInfo := storageresource.SetPersistentVolumeClaimInfo(namespace, name, status, volume, capacity, accessMode, storageClass, volumeMode)
		pvcInfoList = append(pvcInfoList, *pvcInfo)
	}

	return pvcInfoList
}

func GetPersistentVolumeInfo(clientset *kubernetes.Clientset) []storageresource.PersistentVolumeInfo {

	persistentvolumes, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var pvInfoList []storageresource.PersistentVolumeInfo

	for _, pv := range persistentvolumes.Items {
		name := pv.Name
		capacity := pv.Spec.Capacity.Storage().String()
		accessMode := persistentVolumeAccessModesToString(pv.Spec.AccessModes)
		reclaimPolicy := fmt.Sprint(pv.Spec.PersistentVolumeReclaimPolicy)
		status := fmt.Sprint(pv.Status.Phase)
		cliam := pv.Spec.ClaimRef.Name
		storageClass := pv.Spec.StorageClassName
		volumeMode := volumeModeToString(pv.Spec.VolumeMode)
		pvInfo := storageresource.SetPersistentVolumeInfo(name, capacity, accessMode, reclaimPolicy, status, cliam, storageClass, volumeMode)
		pvInfoList = append(pvInfoList, *pvInfo)
	}

	return pvInfoList
}

func persistentVolumeAccessModesToString(ams []corev1.PersistentVolumeAccessMode) string {
	var result []string
	for _, am := range ams {
		result = append(result, fmt.Sprint(am))
	}
	return strings.Join(result, ",")
}

func volumeModeToString(vm *corev1.PersistentVolumeMode) string {
	if vm == nil {
		return "None"
	}

	return fmt.Sprint(*vm)
}
