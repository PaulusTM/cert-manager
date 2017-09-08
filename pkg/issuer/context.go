package issuer

import (
	"k8s.io/client-go/kubernetes"

	"github.com/jetstack-experimental/cert-manager/pkg/client"
	"github.com/jetstack-experimental/cert-manager/pkg/kube"
)

type Context struct {
	Client   kubernetes.Interface
	CMClient client.Interface

	SharedInformerFactory kube.SharedInformerFactory

	Namespace string
}
