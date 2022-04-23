package params

type Param interface {
	Check() error
}
