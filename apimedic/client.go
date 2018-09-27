package apimedic

//Mode is the mode that apimedic is running in (Sandbox or Live).
type Mode int

const (
	//Sandbox uses apimedic's Sandbox
	Sandbox Mode = 0
	//Live uses apimedic's Live mode
	Live Mode = 1
)

type service struct {
	authUrl   string
	healthUrl string
}

var services = map[Mode]service{
	Sandbox: service{
		authUrl:   "https://sandbox-authservice.priaid.ch",
		healthUrl: "https://sandbox-healthservice.priaid.ch",
	},
	Live: service{
		authUrl:   "https://authservice.priaid.ch",
		healthUrl: "https://healthservice.priaid.ch",
	},
}

//Client is a client for accessing apimedic's apis
type Client struct {
	Mode Mode
}

//NewClient returns a new client
func NewClient(m Mode) *Client {
	return &Client{
		Mode: m,
	}
}

func (c *Client) LogIn(uid, password string) (string, error) {
	return "", nil
}
