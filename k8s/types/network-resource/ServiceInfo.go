package networkresource

type ServiceInfo struct {
	namespace   string `xlsx:"NAMESPACE"`
	name        string `xlsx:"NAME"`
	serviceType string `xlsx:"TYPE"`
	clusterIP   string `xlsx:"CLUSTER-IP"`
	externalIP  string `xlsx:"EXTERNAL-IP"`
	ports       string `xlsx:"PORT(S)"`
	selectors   string `xlsx:"SELECTOR"`
}

func SetServiceInfo(
	namespace string,
	name string,
	serviceType string,
	clusterIP string,
	externalIP string,
	ports string,
	selectors string) *ServiceInfo {
	serviceInfo := ServiceInfo{}
	serviceInfo.namespace = namespace
	serviceInfo.name = name
	serviceInfo.serviceType = serviceType
	serviceInfo.clusterIP = clusterIP
	serviceInfo.externalIP = externalIP
	serviceInfo.ports = ports
	serviceInfo.selectors = selectors

	return &serviceInfo
}
