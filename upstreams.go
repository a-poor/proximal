package proximate

type ServerGroup interface {
	AddServer()
	NextServer()
}

type RoundRobinServerGroup struct {
	Servers []string
}
