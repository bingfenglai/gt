package permission

var (
	strColon = []byte(":")
	strStar  = []byte("*")
	strSlash = []byte("/")
)

type methodTree struct {
	method string
	root   *node
}

// 方法树集
type methodTrees []methodTree

func (m methodTrees) getRoot(method string) *node {

	for _, methodTree := range m {
		if methodTree.method == method {
			return methodTree.root
		}
	}

	return nil
}

func NewMethodTrees() *methodTrees {

	return new(methodTrees)
}


