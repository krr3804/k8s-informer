package networkresource

type IngressInfo struct {
	namespace string `xlsx:"NAMESPACE"`
	name      string `xlsx:"NAME"`
	class     string `xlsx:"CLASS"`
	hosts     string `xlsx:"HOSTS"`
	address   string `xlsx:"ADDRESS"`
	ports     string `xlsx:"PORTS"`
}

func SetIngressInfo(
	namespace string,
	name string,
	class string,
	hosts string,
	address string,
	ports string) *IngressInfo {
	ingressInfo := IngressInfo{}
	ingressInfo.namespace = namespace
	ingressInfo.name = name
	ingressInfo.class = class
	ingressInfo.hosts = hosts
	ingressInfo.address = address
	ingressInfo.ports = ports

	return &ingressInfo
}
