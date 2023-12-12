package Domain

type D5Transform struct {
	Source      int64
	Destination int64
	Size        int64
}

type D5SeedRange struct {
	Start int64
	Size  int64
}

type D5Map struct {
	rules []D5Transform
}

func (m *D5Map) Clear() {
	m.rules = []D5Transform{}
}
func (m *D5Map) AddRule(destination, source, size int64) {
	m.rules = append(m.rules, D5Transform{Destination: destination, Source: source, Size: size})
}

func (m *D5Map) Transform(source int64) int64 {
	for _, trans := range m.rules {
		if source >= trans.Source && source < (trans.Source+trans.Size) {
			return trans.Destination - trans.Source + source
		}
	}

	return source
}
