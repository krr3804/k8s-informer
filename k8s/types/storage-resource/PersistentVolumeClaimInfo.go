package storageresource

type PersistentVolumeClaimInfo struct {
	namespace    string `xlsx:"NAMESPACE"`
	name         string `xlsx:"NAME"`
	status       string `xlsx:"STATUS"`
	volume       string `xlsx:"VOLUME"`
	capacity     string `xlsx:"CAPACITY"`
	accessMode   string `xlsx:"ACCESS MODES"`
	storageClass string `xlsx:"STORAGE CLASS"`
	volumeMode   string `xlsx:"VOLUME MODE"`
}

func SetPersistentVolumeClaimInfo(
	namespace string,
	name string,
	status string,
	volume string,
	capacity string,
	accessMode string,
	storageClass string,
	volumeMode string) *PersistentVolumeClaimInfo {
	pvcInfo := PersistentVolumeClaimInfo{}
	pvcInfo.namespace = namespace
	pvcInfo.name = name
	pvcInfo.status = status
	pvcInfo.volume = volume
	pvcInfo.capacity = capacity
	pvcInfo.accessMode = accessMode
	pvcInfo.storageClass = storageClass
	pvcInfo.volumeMode = volumeMode

	return &pvcInfo
}
