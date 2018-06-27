package berlioz

// Endpoint is TBD.
type EndpointModel struct {
	Name            string `json:"name,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	NetworkProtocol string `json:"networkProtocol,omitempty"`
	Port            uint16 `json:"port,omitempty"`
	Address         string `json:"address,omitempty"`
}

type PeersModel map[string]EndpointModel

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

type cloudResourceModel struct {
	Name     string           `json:"name,omitempty"`
	Class    string           `json:"class,omitempty"`
	SubClass string           `json:"subClass,omitempty"`
	Config   cloudConfigModel `json:"config,omitempty"`
}

type messagePeersModel struct {
	Service  map[string]map[string]PeersModel         `json:"service,omitempty"`
	Cluster  map[string]map[string]PeersModel         `json:"cluster,omitempty"`
	Database map[string]map[string]cloudResourceModel `json:"database,omitempty"`
	Queue    map[string]map[string]cloudResourceModel `json:"queue,omitempty"`
}

type agentMessageModel struct {
	Endpoints *map[string]EndpointModel `json:"endpoints,omitempty"`
	Policies  *policyModel              `json:"policies,omitempty"`
	Peers     *messagePeersModel        `json:"peers,omitempty"`
}
