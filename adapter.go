package sugar

type Adapter struct {
	Adapt
}

func (a *Adapter) Do() {
	a.SpecifyDo()
}

type Adapt interface {
	SpecifyDo()
}

type AdaptImpl struct {
}

func (a AdaptImpl) SpecifyDo() {

}

type Target interface {
	Do()
}

func NewAdapter() Target {
	return &Adapter{AdaptImpl{}}
}
