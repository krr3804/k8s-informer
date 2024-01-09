package services

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetClientset(context string, kubeconfig string) *kubernetes.Clientset {

	// use the current context in kubeconfig
	config, err := buildConfigFromFlags(context, kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func GetContexts() ([]string, string) {
	var contexts []string
	kubeconfig := getKubeconfigPath()

	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		fmt.Printf("Error loading kubeconfig: %v\n", err)
		os.Exit(1)
	}

	for context := range config.Contexts {
		contexts = append(contexts, context)
	}

	return contexts, kubeconfig
}

func GetCurrentCluster(context string, kubeconfigPath string) string {
	var currentCluster string

	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		fmt.Printf("Error loading kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// kubeconfig 파일에서 현재 context에 대한 정보 확인
	contexts := config.Contexts
	for ctxName, ctx := range contexts {

		if ctxName == context {
			// 현재 context와 일치하는 경우 해당 클러스터 정보 가져오기
			// clusterName := ctx.Cluster
			// cluster, exists := config.Clusters[clusterName]
			// if exists {
			// 	currentCluster = cluster.Server
			// }
			// break
			currentCluster = ctx.Cluster
		}
	}
	parts := strings.Split(currentCluster, "/") // 문자열을 '/'로 분할
	clusterName := parts[len(parts)-1]          // 분할된 문자열 중 마지막 요소 선택

	return clusterName

}

func getKubeconfigPath() string {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return *kubeconfig
}

func buildConfigFromFlags(context string, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
