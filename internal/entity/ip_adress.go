package entity

type Ip struct {
	Address string
	limit   float64
}

func NewIp(address string, limit float64) *Ip {
	return &Ip{
		Address: address,
		limit:   limit,
	}
}
