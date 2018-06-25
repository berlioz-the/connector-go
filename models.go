package berlioz


// Endpoint is TBD.
type Endpoint struct {
	Name            string `json:"name,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	NetworkProtocol string `json:"networkProtocol,omitempty"`
	Port            uint16 `json:"port,omitempty"`
	Address         string `json:"address,omitempty"`
}

type policy struct {
	Values   map[string]interface{} `json:"values,omitempty"`
	Children map[string]policy      `json:"children,omitempty"`
}

type cloudCredentials struct {
	AccessKeyID     string `json:"accessKeyId,omitempty"`
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
}

type cloudConfig struct {
	Region      string           `json:"region,omitempty"`
	Credentials cloudCredentials `json:"credentials,omitempty"`
}

type cloudResource struct {
	Name     string      `json:"name,omitempty"`
	Class    string      `json:"class,omitempty"`
	SubClass string      `json:"subClass,omitempty"`
	Config   cloudConfig `json:"config,omitempty"`
}
