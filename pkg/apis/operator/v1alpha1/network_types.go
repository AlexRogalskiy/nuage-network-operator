package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReleaseConfigDefinition holds the release tag for each component and registry details
type ReleaseConfigDefinition struct {
	Registry RegistryConfig `json:"registry"`
	// +kubebuilder:validation:MinLength=1
	VRSTag string `json:"vrsTag"`
	// +kubebuilder:validation:MinLength=1
	CNITag string `json:"cniTag"`
	// +kubebuilder:validation:MinLength=1
	MonitorTag string `json:"monitorTag"`
}

// MonitorConfigDefinition holds user specified config for monitor
type MonitorConfigDefinition struct {
	// +kubebuilder:validation:MinLength=1
	VSDAddress string `json:"vsdAddress"`
	// +kubebuilder:validation:Minimum=0
	VSDPort                int      `json:"vsdPort"`
	VSDMetadata            Metadata `json:"vsdMetadata"`
	VSDFlags               Flags    `json:"vsdFlags"`
	RestServerAddress      string   `json:"restServerAddress,omitempty"`
	RestServerPort         int      `json:"restServerPort,omitempty"`
	ServiceAccountName     string   `json:"ServiceAccountName,omitempty"`
	ClusterRoleName        string   `json:"ClusterRoleName,omitempty"`
	ClusterRoleBindingName string   `json:"ClusterRoleBindingName,omitempty"`
}

// VRSConfigDefinition holds user specified config for VRS
type VRSConfigDefinition struct {
	// +kubebuilder:validation:MinItems=1
	Controllers []string `json:"controllers"`
	// +kubebuilder:validation:MinLength=1
	UnderlayUplink string `json:"underlayUplink"`
	Platform       string `json:"platform,omitempty"`
}

// CNIConfigDefinition holds user specified config for CNI
type CNIConfigDefinition struct {
	VRSEndpoint             string `json:"vrsEndpoint,omitempty"`
	VRSBridge               string `json:"vrsBridge,omitempty"`
	CNIVersion              string `json:"cniVersion,omitempty"`
	LogLevel                string `json:"logLevel,omitempty"`
	MTU                     int    `json:"mtu,omitempty"`
	NuageSiteID             int    `json:"nuageSiteID,omitempty"`
	LogFileSize             int    `json:"logFileSize,omitempty"`
	MonitorInterval         int    `json:"monitorInterval,omitempty"`
	PortResolveTimer        int    `json:"portResolveTimer,omitempty"`
	VRSConnectionCheckTimer int    `json:"vrsConnectionCheckTimer,omitempty"`
	StaleEntryTimeout       int    `json:"staleEntryTimeout,omitempty"`
	ServiceAccountName      string `json:"serviceAccountName,omitempty"`
	ClusterRoleName         string `json:"clusterRoleName,omitempty"`
	ClusterRoleBindingName  string `json:"clusterRoleBindingName,omitempty"`
	NuageMonitorURL         string `json:"nuageMonitorURL,omitempty"`
	KubeConfig              string `json:"kubeConfig,omitempty"`
}

// Metadata holds the VSD metadata info
type Metadata struct {
	// +kubebuilder:validation:MinLength=1
	Enterprise string `json:"enterprise"`
	// +kubebuilder:validation:MinLength=1
	Domain string `json:"domain"`
	// +kubebuilder:validation:MinLength=1
	User string `json:"user"`
	// +kubebuilder:validation:MinLength=1
	UserCert string `json:"userCert"`
	// +kubebuilder:validation:MinLength=1
	UserKey string `json:"userKey"`
}

// Flags hold the flags for VSD behaviors
type Flags struct {
	UnderlayEnabled  bool `json:"underlayEnabled,omitempty"`
	StatsEnabled     bool `json:"statsEnabled,omitempty"`
	AutoScaleSubnets bool `json:"autoScaleSubnets,omitempty"`
}

// RegistryConfig holds the registry information
type RegistryConfig struct {
	// +kubebuilder:validation:MinLength=1
	URL string `json:"url"`
	// +kubebuilder:validation:MinLength=1
	Username string `json:"username"`
	// +kubebuilder:validation:MinLength=1
	Password string `json:"password"`
}

// NetworkSpec defines the desired state of Network
// +k8s:openapi-gen=true
type NetworkSpec struct {
	VRSConfig     VRSConfigDefinition     `json:"vrsConfig"`
	CNIConfig     CNIConfigDefinition     `json:"cniConfig"`
	MonitorConfig MonitorConfigDefinition `json:"monitorConfig"`
	ReleaseConfig ReleaseConfigDefinition `json:"releaseConfig"`
}

// NetworkStatus defines the observed state of Network
// +k8s:openapi-gen=true
type NetworkStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network is the Schema for the networks API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +genclient:nonNamespaced
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSpec   `json:"spec,omitempty"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList contains a list of Network
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// ClusterNetworkConfigDefinition contains the network configuration of cluster
type ClusterNetworkConfigDefinition struct {
	ClusterNetworkCIDR         string
	ServiceNetworkCIDR         string
	ClusterNetworkSubnetLength uint32
}

// TLSCertificates contains certificates for CNI and Monitor
type TLSCertificates struct {
	CA             *string
	Certificate    *string
	PrivateKey     *string
	CertificateDir *string
}

// RenderConfig container to hold config data that is passed to rendering logic
type RenderConfig struct {
	NetworkSpec
	K8SAPIServerURL      string
	Certificates         *TLSCertificates
	ClusterNetworkConfig *ClusterNetworkConfigDefinition
}

// CertGenConfig certificate data for input generation
type CertGenConfig struct {
	ECDSACurve *string
	ValidFrom  *string
	ValidFor   time.Duration
	RSABits    int
}

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
