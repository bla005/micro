package service

type Dependency struct {
	Category string       `json:"category"`
	Name     string       `json:"name"`
	Ping     func() error `json:"-"`
}

func NewDependency(category, name string, ping func() error) *Dependency {
	return &Dependency{Category: category, Name: name, Ping: ping}
}
