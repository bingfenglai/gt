package config

type ServerConfig struct {
	Address    string `json:"address,omitempty"`
	Port       int    `json:"port,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Url404     string `json:"url_404,omitempty"`
	Urlfavicon string `json:"url_favicon,omitempty"`
	Encrypted  bool   `json:"encrypted,omitempty"`
	EnableAuth bool
}
