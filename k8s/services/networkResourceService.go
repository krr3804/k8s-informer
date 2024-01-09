package services

import (
	"context"
	networkresource "k8s-informer/k8s/types/network-resource"
	"k8s-informer/util"
	"strconv"
	"strings"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

/*************
* Service 정보
**************/
func GetServiceInfo(clientset *kubernetes.Clientset) []networkresource.ServiceInfo {

	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var serviceInfoList []networkresource.ServiceInfo

	for _, service := range services.Items {
		namespace := service.Namespace
		name := service.Name
		serviceType := string(service.Spec.Type)
		clusterIP := service.Spec.ClusterIP
		externalIP := func(ips []string) string {
			if len(ips) == 0 {
				return "None"
			}
			return strings.Join(ips, ",")
		}(service.Spec.ExternalIPs)

		ports := func(ports []corev1.ServicePort) string {
			var results []string
			for _, port := range ports {
				result := strconv.Itoa(int(port.Port))
				if port.NodePort > 0 {
					result = result + ":" + strconv.Itoa(int(port.NodePort))
				}
				result = result + "/" + string(port.Protocol)
				results = append(results, result)
			}

			return strings.Join(results, ",")
		}(service.Spec.Ports)
		selectors := util.MapToString(service.Spec.Selector, "=")
		serviceInfo := networkresource.SetServiceInfo(namespace, name, serviceType, clusterIP, externalIP, ports, selectors)
		serviceInfoList = append(serviceInfoList, *serviceInfo)
	}

	return serviceInfoList
}

/*************
* Ingress 정보
**************/
func GetIngressInfo(clientset *kubernetes.Clientset) []networkresource.IngressInfo {
	ingresses, err := clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var ingressInfoList []networkresource.IngressInfo

	for _, ing := range ingresses.Items {
		namespace := ing.Namespace
		name := ing.Name
		class := util.StringPointerToString(ing.Spec.IngressClassName)
		hosts, ports := func(rules []networkingv1.IngressRule) (string, string) {
			var result_hosts []string
			var result_ports []string
			for _, rule := range rules {
				result_hosts = append(result_hosts, rule.Host)
				var p []string
				for _, path := range rule.HTTP.Paths {
					p = append(p, strconv.Itoa(int(path.Backend.Service.Port.Number)))
				}
				result_ports = append(result_ports, strings.Join(p, ","))
			}
			return strings.Join(result_hosts, ","), strings.Join(result_ports, ",")
		}(ing.Spec.Rules)
		address := func(ingresses []networkingv1.IngressLoadBalancerIngress) string {
			var results []string
			for _, ingress := range ingresses {
				results = append(results, ingress.Hostname)
			}
			return strings.Join(results, ",")
		}(ing.Status.LoadBalancer.Ingress)

		ingressInfo := networkresource.SetIngressInfo(namespace, name, class, hosts, address, ports)
		ingressInfoList = append(ingressInfoList, *ingressInfo)
	}

	return ingressInfoList
}
