package config

// StorageConfig database configure
type StorageConfig struct {
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,flow,omitempty" envconfig:"STORAGE_ENDPOINT"`
	User     string `json:"user,omitempty" yaml:"user,omitempty" envconfig:"STORAGE_USER"`
	Password string `json:"password,omitempty" yaml:"password,omitempty" envconfig:"STORAGE_PASSWORD"`
	DB       string `json:"db,omitempty" yaml:"db,omitempty" envconfig:"STORAGE_DB"`
}

// SMTPConfig smtp configure to send email
type SMTPConfig struct {
	Host     string `json:"host,omitempty" yaml:"host,omitempty" envconfig:"SMTP_HOST"`
	Port     uint   `json:"port,omitempty" yaml:"port,omitempty" envconfig:"SMTP_PORT"`
	User     string `json:"user,omitempty" yaml:"user,omitempty" envconfig:"SMTP_USER"`
	Password string `json:"password,omitempty" yaml:"password,omitempty" envconfig:"SMTP_USER"`
}

// OAuth2 client configure
type OAuth2 struct {
	OAuth2LoginURL   string `json:"oauth2LoginURL,omitempty" yaml:"oauth2LoginURL,omitempty" envconfig:"OAUTH_LOGIN_URL"`
	OAuth2TokenURL   string `json:"oauth2TokenURL,omitempty" yaml:"oauth2TokenURL,omitempty" envconfig:"OAUTH_TOKEN_URL"`
	OAuth2RefreshURL string `json:"oauth2RefreshURL,omitempty" yaml:"oauth2RefreshURL,omitempty" envconfig:"OAUTH_REFRESH_URL"`
	OAuth2InfoURL    string `json:"oauth2InfoURL,omitempty" yaml:"oauth2InfoURL,omitempty" envconfig:"OAUTH_INFO_URL"`
	ClientKeyID      string `json:"clientKeyID,omitempty" yaml:"clientKeyID,omitempty" envconfig:"CLIENT_KEY_ID"`
	ClientKeySecret  string `json:"clientKeySecret,omitempty" yaml:"clientKeySecret,omitempty" envconfig:"CLIENT_KEY_SECRET"`
}
