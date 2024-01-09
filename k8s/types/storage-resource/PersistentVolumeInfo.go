package storageresource

type PersistentVolumeInfo struct {
	name          string `xlsx:"NAME"`
	capacity      string `xlsx:"CAPACITY"`
	accessMode    string `xlsx:"ACCESS MODES"`
	reclaimPolicy string `xlsx:"RECLAIM POLICY"`
	status        string `xlsx:"STATUS"`
	claim         string `xlsx:"CLAIM"`
	storageClass  string `xlsx:"STORAGE CLASS"`
	volumeMode    string `xlsx:"VOLUME MODE"`
}

func SetPersistentVolumeInfo(
	name string,
	capacity string,
	accessMode string,
	reclaimPolicy string,
	status string,
	claim string,
	storageClass string,
	volumeMode string) *PersistentVolumeInfo {
	pvInfo := PersistentVolumeInfo{}
	pvInfo.name = name
	pvInfo.capacity = capacity
	pvInfo.accessMode = accessMode
	pvInfo.reclaimPolicy = reclaimPolicy
	pvInfo.status = status
	pvInfo.claim = claim
	pvInfo.storageClass = storageClass
	pvInfo.volumeMode = volumeMode

	return &pvInfo
}
