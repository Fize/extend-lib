package config

// StorageConfig database configure
type StorageConfig struct {
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,flow,omitempty"`
	User     string `json:"user,omitempty" yaml:"user,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	DB       string `json:"db,omitempty" yaml:"db,omitempty"`
}

// SMTPConfig smtp configure to send email
type SMTPConfig struct {
	Host     string `json:"host,omitempty" yaml:"host,omitempty"`
	Port     uint   `json:"port,omitempty" yaml:"port,omitempty"`
	User     string `json:"user,omitempty" yaml:"user,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

// OAuth2 client configure
type OAuth2 struct {
	OAuth2LoginURL   string `json:"oauth2LoginURL,omitempty" yaml:"oauth2LoginURL,omitempty"`
	OAuth2TokenURL   string `json:"oauth2TokenURL,omitempty" yaml:"oauth2TokenURL,omitempty"`
	OAuth2RefreshURL string `json:"oauth2RefreshURL,omitempty" yaml:"oauth2RefreshURL,omitempty"`
	OAuth2InfoURL    string `json:"oauth2InfoURL,omitempty" yaml:"oauth2InfoURL,omitempty"`
	ClientKeyID      string `json:"clientKeyID,omitempty" yaml:"clientKeyID,omitempty"`
	ClientKeySecret  string `json:"clientKeySecret,omitempty" yaml:"clientKeySecret,omitempty"`
}
