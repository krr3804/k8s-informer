package types

import eksresource "k8s-informer/aws/types/eks"

type ResourceType interface {
	eksresource.ClusterInfo | eksresource.NodeGroupInfo
}
