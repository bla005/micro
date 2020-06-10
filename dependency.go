package service

type Dependency struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	ping     func() error
}

func NewDependency(category, name string, ping func() error) *Dependency {
	return &Dependency{Category: category, Name: name, ping: ping}
}
