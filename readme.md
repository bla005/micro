# Example

```golang
func main() {
	config1 := service.NewConfig()
	router1 := service.Router()
	server := service.NewServer(router1, config1)
	service1, err := service.NewService("mainService", "1.0", config1, server1, router1)
	if err != nil {
		log.Fatalln(err)
	}
	service1.UseHealthEndpoint()
	service1.UseDependency(
		service.MakeDependency("Database", "Postgres", pingme),
		service.MakeDependency("Store", "Redis", redispingme),
	)
	service1.UseEndpoint(
		service.MakeEndpoint("info", "GET", "/info", Info),                                                                        
		service.MakeEndpoint("stats", "GET", "/stats", Stats),
	}
	service1.Start()
}

func Info(w http.ResponseWriter, r *http.Request) {}
func Stats(w http.ResponseWriter, r *http.Request) {}
```
