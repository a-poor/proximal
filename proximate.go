package proximate

type Proxy struct {
	idx       int
	Upstreams []string
}

func (p *Proxy) AddUpstream(up string) {
	if p.Upstreams == nil {
		p.Upstreams = make([]string, 1)
	}
	p.Upstreams = append(p.Upstreams, up)
}

func (p *Proxy) incUpstream() {
	p.idx = (p.idx + 1) % len(p.Upstreams)
}

func (p *Proxy) getUpstream() string {
	return p.Upstreams[p.idx]
}

func (p *Proxy) GetNextUpstream() string {
	up := p.getUpstream()
	p.incUpstream()
	return up
}

func (p *Proxy) Run() error {
	return nil
}
