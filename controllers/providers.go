package controllers

import (
	v1alpha1 "github.com/RHEcosystemAppEng/dbaas-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DBaaSInventoryReconciler) getDBaaSProviders(client client.Client, scheme *runtime.Scheme) v1alpha1.DBaaSProviderList {
	return v1alpha1.DBaaSProviderList{
		Items: []v1alpha1.DBaaSProvider{
			{
				TypeMeta: metav1.TypeMeta{},
				Spec: v1alpha1.DBaaSProviderSpec{
					TypeMeta: metav1.TypeMeta{},
					ListMeta: metav1.ListMeta{},
					Provider: v1alpha1.DatabaseProvider{
						Name: "MongoDB Atlas",
					},
					InventoryKind: "AtlasAccount",
					AuthenticationFields: []v1alpha1.AuthenticationField{
						{
							Name: "Organization ID",
						},
						{
							Name: "Organization Public Key",
						},
						{
							Name: "Organization Private Key",
						},
					},
					ConnectionKind: "AtlasConnection",
				},
			},
		},
	}
}
