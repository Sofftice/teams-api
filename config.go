package teamsapi

type Config struct {
	ClientId string
}

func DefaultConfig() Config {
	return Config{
		ClientId: "5e3ce6c0-2b1f-4285-8d4b-75ee78787346",
	}
}
