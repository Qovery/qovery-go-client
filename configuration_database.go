package qovery

type DatabaseConfiguration struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Version  string `json:"version"`
}
