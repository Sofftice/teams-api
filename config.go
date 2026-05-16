package teamsapi

type Config struct {
	client_id string
}

func DefaultConfig() Config {
	return Config{
		client_id: "5e3ce6c0-2b1f-4285-8d4b-75ee78787346",
	}
}