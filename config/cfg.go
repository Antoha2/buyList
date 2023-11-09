package cfg

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	HostAddr   string
	ContextKey ContextKey
}

func GetConfig() *Config {

	return &Config{
		
		HTTP: HTTPConfig{
			HostAddr:   os.Getenv("HOST_ADDR"),
			ContextKey: "History",
		}
	}

}
