package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// DBaaSProviderRegistrationSpec defines a database provider for DBaaS operator
type DBaaSProviderRegistrationSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Provider is the name of the database provider
	Provider DatabaseProvider `json:"provider"`

	// The name of the inventory CRD defined by the provider
	InventoryKind string `json:"inventoryKind"`

	// The authentication fields the user needs to receive and provide
	AuthenticationFields []AuthenticationField `json:"authenticationFields"`

	// The name of the connection CRD defined by the provider
	ConnectionKind string `json:"connectionKind"`
}

type DatabaseProvider struct {
	Name string `json:"name"`
}

type AuthenticationField struct {
	Name string `json:"name"`
	Masked bool `json:"masked"`
}

// DBaaSProviderRegistrationList contains a list of DBaaSProviderRegistrations
type DBaaSProviderRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBaaSProviderRegistrationSpec   `json:"spec,omitempty"`
}

// DBaaSInventorySpec defines the desired state of DBaaSInventory
type DBaaSInventorySpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// The secret storing the vendor-specific connection credentials to
	// use with the API endpoint. The secret may be placed in a separate
	// namespace to control access.
	CredentialsRef *NamespacedName `json:"credentialsRef"`
}

// DBaaSInventoryStatus defines the observed state of DBaaSInventory
type DBaaSInventoryStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// E.g., MongoDB, Postgres
	Type string `json:"type"`

	// A list of instances returned from querying the DB provider
	Instances []Instance `json:"instances,omitempty"`
}

type Instance struct {
	// The ID for this instance in the database service
	InstanceID string `json:"instanceID"`

	// The name of this instance in the database service
	Name string `json:"name,omitempty"`

	// Any other provider-specific information related to this instance
	InstanceInfo map[string]string `json:"extraInfo,omitempty"`
}

type NamespacedName struct {
	// The namespace where object of known type is store
	Namespace string `json:"namespace,omitempty"`

	// The name for object of known type
	Name string `json:"name,omitempty"`
}