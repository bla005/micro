package service

type Dependency struct {
	category string
	name     string
	ping     func() error
}

func NewDependency(category, name string, ping func() error) *Dependency {
	return &Dependency{category: category, name: name, ping: ping}
}
