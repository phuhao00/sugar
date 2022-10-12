package sugar

type Proxy struct {
	real proxy
}

func (p *Proxy) Execute() {
	p.real.Execute()
}

type proxy interface {
	Execute()
}
