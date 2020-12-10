package overider

type Patch struct {
	Port               int
	SocksPort          int
	RedirPort          int
	AllowLan           bool
	ExternalController string
	Secret             string
	Proxies            []Proxy
}

type Proxy struct {
	Type     string
	Name     string
	Server   string
	Port     int
	Cipher   string
	Password string
	Udp      bool
	UUID     string
	AlterId  int
	WsPath   string
	WsHeader map[string]string
}
