package nameserver

import (
	"context"

	"github.com/loft-sh/vcluster/pkg/controllers/resources/specialservices"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var Default Interface = DefaultNameserverFinder()

const (
	DefaultKubeDNSServiceName      = "kube-dns"
	DefaultKubeDNSServiceNamespace = "kube-system"
)

type SpecialServiceSyncer func(ctx context.Context,
	vClient,
	pClient client.Client,
	svcNamespace,
	svcName string,
	vSvcToSync types.NamespacedName,
	servicePortTranslator specialservices.ServicePortTranslator) error

type Interface interface {
	GetDNSServiceSuffix() *string
	SpecialServicesToSync() map[types.NamespacedName]SpecialServiceSyncer
}

type NameserverFinder struct {
	DNSServiceSuffix *string
	SpecialServices  map[types.NamespacedName]SpecialServiceSyncer
}

func (f *NameserverFinder) GetDNSServiceSuffix() *string {
	return f.DNSServiceSuffix
}

func (f *NameserverFinder) SpecialServicesToSync() map[types.NamespacedName]SpecialServiceSyncer {
	return f.SpecialServices
}

func DefaultNameserverFinder() Interface {
	return &NameserverFinder{
		SpecialServices: map[types.NamespacedName]SpecialServiceSyncer{
			{
				Name:      specialservices.DefaultKubernetesSVCName,
				Namespace: specialservices.DefaultKubernetesSVCNamespace,
			}: specialservices.SyncKubernetesService,
		},
	}
}
