package storage

type Message struct {
	person []PersonService
}

func (m *Message) AddMessage(p PersonService) {
	m.person = append(m.person, p)
}

func (m *Message) Notify(p ...*Person) {
	for _, v := range p {
		for _, k := range m.person {
			k.Add(v)
		}
	}
}
