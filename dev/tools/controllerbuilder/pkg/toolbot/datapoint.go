package toolbot

type DataPoint struct {
	Type   string
	Input  map[string]string
	Output string
}

func (p *DataPoint) SetInput(k, v string) {
	if p.Input == nil {
		p.Input = make(map[string]string)
	}
	p.Input[k] = v
}
