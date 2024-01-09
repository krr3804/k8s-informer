package configurationresource

type ConfigMapInfo struct {
	namespace string `xlsx:"NAMESPACE"`
	name      string `xlsx:"NAME"`
	data      int    `xlsx:"DATA"`
}

func SetConfigMapInfo(namespace string, name string, data int) *ConfigMapInfo {
	cmInfo := ConfigMapInfo{}
	cmInfo.namespace = namespace
	cmInfo.name = name
	cmInfo.data = data

	return &cmInfo

}
