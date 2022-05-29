package domain

var ps = &ParamSet{
	data: make(map[string]string),
}

type ParamSet struct {
	data map[string]string
}

func GetParamSet() *ParamSet {
	return ps
}

func (p *ParamSet) Set(key string, value string) error {
	p.data[key] = value
	return nil
}

func (p *ParamSet) Get(key string) string {
	return p.data[key]
}
