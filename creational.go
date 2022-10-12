package sugar

type Builder interface {
	Build(m any) Builder
	CompletedCheck() bool
}
