package config

type Listen struct {
	Addr           string
	Backlog        int
	AllowedClients []string
	Owner          string
	Group          string
	Mode           string
	ACLUsers       []string
	ACLGroups      []string
}
