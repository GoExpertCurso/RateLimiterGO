package entity

type Request struct {
	Ip    *Ip
	Token *Token
	limit float64
}

func NewRequest(ip *Ip, token *Token) *Request {
	return &Request{
		Ip:    ip,
		Token: token,
	}
}

func (r *Request) LimitCheck() {
	if r.Token.limit >= r.Ip.limit {
		r.limit = r.Token.limit
	} else {
		r.limit = r.Ip.limit
	}
}
