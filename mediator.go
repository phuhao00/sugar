package sugar

//Mediator  a FSM implement

type Mediator struct {
	transitions map[Member][]Member
}

var (
	gMediator *Mediator
)

func GetMediator() *Mediator {
	if gMediator == nil {
		gMediator = &Mediator{make(map[Member][]Member)}
	}
	return gMediator
}

func (m *Mediator) Changed(i Member) {
	for _, member := range m.transitions[i] {
		member.Process(i.GetTransData())
	}
}

func (m *Mediator) RegisterMember(from, to Member) {
	if m.transitions[from] == nil {
		m.transitions[from] = make([]Member, 0)
	}
	m.transitions[from] = append(m.transitions[from], to)
}

type Member interface {
	Process(a any)
	GetTransData() any
}
