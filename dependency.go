package service

// Dependency structure
type Dependency struct {
	// Type of the depedency (Database, Store, Message Queue etc.)
	_type string
	// Name of the dependency (Postgres, MySQL, Cassandra, Redis)
	name string
	// Ping is used for health checks
	Ping func() error
}

// func (e *Dependency) ping() error {
// 	return e.Ping()
// }

func (d *Dependency) Type() string {
	return d._type
}
func (d *Dependency) Name() string {
	return d.name
}

func MakeDependency(_type, name string, ping func() error) *Dependency {
	return &Dependency{_type: _type, name: name, Ping: ping}
}
