package securityresource

type ServiceAccountInfo struct {
	namespace string `xlsx:"NAMESPACE"`
	name      string `xlsx:"NAME"`
	secrets   int    `xlsx:"SECRETS"`
}

func SetServiceAccountInfo(namespace string, name string, secrets int) *ServiceAccountInfo {
	saInfo := ServiceAccountInfo{}
	saInfo.namespace = namespace
	saInfo.name = name
	saInfo.secrets = secrets
	return &saInfo
}
