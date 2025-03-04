package edgetts

type Option struct {
	Voice          string
	Rate           string
	Volume         string
	Pitch          string
	ConnectTimeout int
	ReceiveTimeout int
}

func DefaultOption() Option {
	return gDefaultOption
}
