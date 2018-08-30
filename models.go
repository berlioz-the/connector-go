package berlioz

// Endpoint is TBD.
type EndpointModel struct {
	Name            string `json:"name,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	NetworkProtocol string `json:"networkProtocol,omitempty"`
	Port            uint16 `json:"port,omitempty"`
	Address         string `json:"address,omitempty"`
}

// ConsumesModel is TBD.
type ConsumesModel struct {
	Kind     string `json:"kind,omitempty"`
	ID       string `json:"id,omitempty"`
	Cluster  string `json:"cluster,omitempty"`
	Sector   string `json:"sector,omitempty"`
	Name     string `json:"name,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
}

type PeersModel map[string]EndpointModel

type EndpointsModel map[string]PeersModel

type policyModel struct {
	Values   map[string]interface{} `json:"values,omitempty"`
	Children map[string]policyModel `json:"children,omitempty"`
}

type cloudCredentialsModel struct {
	AccessKeyID     string `json:"accessKeyId,omitempty"`
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
}

type cloudConfigModel struct {
	Region      string                `json:"region,omitempty"`
	Credentials cloudCredentialsModel `json:"credentials,omitempty"`
}

type CloudResourceModel struct {
	Name     string           `json:"name,omitempty"`
	Class    string           `json:"class,omitempty"`
	SubClass string           `json:"subClass,omitempty"`
	Config   cloudConfigModel `json:"config,omitempty"`
}

type ResourceModel map[string]CloudResourceModel

type CloudResourcesModel map[string]CloudResourceModel

// type messagePeersModel struct {
// 	Service          map[string]map[string]PeersModel `json:"service,omitempty"`
// 	Cluster          map[string]map[string]PeersModel `json:"cluster,omitempty"`
// 	Database         map[string]CloudResourcesModel   `json:"database,omitempty"`
// 	Queue            map[string]CloudResourcesModel   `json:"queue,omitempty"`
// 	SecretPublicKey  map[string]CloudResourcesModel   `json:"secret_public_key,omitempty"`
// 	SecretPrivateKey map[string]CloudResourcesModel   `json:"secret_private_key,omitempty"`
// }

type agentMessageModel struct {
	Endpoints *map[string]EndpointModel `json:"endpoints,omitempty"`
	Policies  *policyModel              `json:"policies,omitempty"`
	Peers     *map[string]interface{}   `json:"peers,omitempty"`
	Consumes  *[]ConsumesModel          `json:"consumes,omitempty"`
}
