package configurationresource

type SecretInfo struct {
	namespace  string `xlsx:"NAMESPACE"`
	name       string `xlsx:"NAME"`
	secretType string `xlsx:"TYPE"`
	data       int    `xlsx:"DATA"`
}

func SetSecretInfo(namespace string, name string, secretType string, data int) *SecretInfo {
	secretInfo := SecretInfo{}
	secretInfo.namespace = namespace
	secretInfo.name = name
	secretInfo.secretType = secretType
	secretInfo.data = data

	return &secretInfo

}
