package permission

import (
	"log"
	"strconv"
	"strings"
)

type nodeType uint8

const (
	root nodeType = iota
	param
	catchAll
)

// 权限节点
type node struct {
	path      string   // 节点路径值
	fullPath  string   // 原路径
	perfix    string   // 公共前缀
	children  []*node  // 子节点
	wildChild bool     // 是否有子节点
	nType     nodeType // 子节点类型
	priority  uint32   // 节点优先级

}

func (n *node) AddPath(path string) {
	fullPath := path

	// 如果此树为空
	if len(n.path) == 0 && len(n.children) == 0 {
		// 设当前节点为跟节点
		n.nType = root
		// 并直接插入子节点
		n.insertChild(path, fullPath)
		return
	}
}

func (n *node) insertChild(path, fullPath string) {

	for {

		//  查找路径当中的第一个通配符
		wildcard, index, valid := findWildcard(path)

		// 不存在通配符
		if index == -1 {
			break
		}

		// The wildcard name must not contain ':' and '*'
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		// check if the wildcard has a name
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

	}

}

// 查找路径当中的第一个通配符, 如果没有找到则返回索引index=-1
func findWildcard(path string) (wildcard string, index int, valid bool) {

	// FindOne start
	for start, c := range []byte(path) {
		// A wildcard starts with ':' (param) or '*' (catch-all)
		if c != ':' && c != '*' {
			continue
		}

		// FindOne end and check for invalid characters
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			case '/':
				return path[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}
		return path[start:], start, valid
	}

	return "", -1, false
}

func NewNode() *node {

	return new(node)
}

func NewNodeWithPath(path string) *node {
	// fullPath := path

	nodePaths := strings.Split(path, "/")

	log.Default().Println("当前路径拆分节点：", nodePaths)

	n, _ := BuildNodeLink(nodePaths, nil)

	return n

}

// 将uri路径转换为node链
func BuildNodeLink(nps []string, n *node) (*node, []string) {

	if n != nil {
		log.Default().Println("当前node{ ", strconv.Itoa(int(n.nType)), n.path, n.children, " }")
	}
	log.Default().Println("当前paths: ", nps)

	if n == nil {
		// 构建链头
		n = new(node)
		n.nType = root
		n.perfix = nps[0]
		n.path = nps[0]
		n.wildChild = false

		if len(nps) > 1 {
			_, nps = BuildNodeLink(nps[1:], n)
		}

	} else {
		path := nps[0]
		nt := catchAll
		if strings.Contains(path, ":") || strings.Contains(path, "*") {
			nt = param
		}

		if len(n.children) == 0 {
			child := &node{
				path:      path,
				nType:     nt,
				perfix:    n.perfix + "/" + path,
				wildChild: false,
			}
			n.children = append(n.children, child)
			n.wildChild = true

			if len(nps) > 1 {
				_, nps = BuildNodeLink(nps[1:], child)
			} else {
				log.Default().Println("执行完成")
				return n, nps
			}

		}

	}

	return n, nps
}

// 迭代单链
func Iterate(n *node, paths []string) (*node, []string) {
	if n == nil {
		return n, paths
	}
	log.Default().Println("迭代，当前path=", n.path)
	log.Default().Println("迭代，当前paths=", paths)
	paths = append(paths, n.path)

	if n.wildChild {
		for _, c := range n.children {
			_, paths = Iterate(c, paths)
		}
	}

	return n, paths
}

func IsAuthoried(root *node, path string) bool {

	if !root.wildChild {
		return false
	}
	if path[:0] == "/" {
		path = path[1:]
	}

	pa := strings.Split(path, "/")
	return isAuthoried(root, pa, false)

}

func isAuthoried(n *node, paths []string, matchChild bool) bool {
	flag := false
	if n == nil {
		return flag
	}

	// n为根节点
	if n.path == "/" {
		for _, node := range n.children {
			if node.path == paths[0] {
				flag = true
				if len(paths) > 1 {
					flag = isAuthoried(node, paths[1:], true)
					break
				} else {
					return flag
				}
			}

		}
	} else if !matchChild {
		if n.path == paths[0] {

			log.Default().Println("当前节点匹配path=", n.path)
			flag = true
			if len(paths) > 1 {
				paths = paths[1:]

				for _, node := range n.children {
					if node.path == paths[0] {
						flag = true
						log.Default().Println("当前节点匹配path=", node.path)
						if len(paths) > 1 {
							flag = isAuthoried(node, paths[1:], true)
							break
						} else {
							return flag
						}
					}
				}

			} else {
				return flag
			}
		}
	} else if len(n.children) > 0 {

		for _, node := range n.children {
			if node.path == paths[0] {
				flag = true
				log.Default().Println("当前节点匹配path=", node.path)
				if len(paths) > 1 {
					flag = isAuthoried(node, paths[1:], true)
					break
				} else {
					return flag
				}
			}
		}

	} else {

		flag = false
	}

	return flag

}

// 将链添加到树中
func addLink(tree *node, link *node) {

	if tree.nType == root {
		for _, child := range tree.children {
			if child.path == link.path {

				for _, ch := range child.children {
					if ch.path == link.children[0].path {

					}
				}

			} else {
				tree.children = append(tree.children, link)
				return
			}
		}
	}
}
