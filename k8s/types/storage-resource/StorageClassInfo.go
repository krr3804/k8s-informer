package storageresource

type StorageClassInfo struct {
	name                 string `xlsx:"NAME"`
	provisioner          string `xlsx:"PROVISIONER"`
	reclaimPolicy        string `xlsx:"RECLAIM POLICY"`
	volumeBindingMode    string `xlsx:"VOLUME BINDING MODE"`
	allowVolumeExpansion bool   `xlsx:"ALLOW VOLUME EXPANSION"`
}

func SetStorageClassInfo(
	name string,
	provisioner string,
	reclaimPolicy string,
	volumeBindingMode string,
	allowVolumeExpansion bool) *StorageClassInfo {
	scInfo := StorageClassInfo{}
	scInfo.name = name
	scInfo.provisioner = provisioner
	scInfo.reclaimPolicy = reclaimPolicy
	scInfo.volumeBindingMode = volumeBindingMode
	scInfo.allowVolumeExpansion = allowVolumeExpansion

	return &scInfo
}
