package service

// Dependency structure
type Dependency struct {
	// Type of the depedency (Database, Store, Message Queue etc.)
	Type string
	// Name of the dependency (Postgres, MySQL, Cassandra, Redis)
	Name string
	// Ping is used for health checks
	Ping func() error
}

// func (e *Dependency) ping() error {
// 	return e.Ping()
// }

func (s *service) RegisterDependency(d *Dependency) {
	s.dependencies = append(s.dependencies, d)
}
