package v2

import (
	"log"
)

// 每个http方法对应一颗树
type MethodTrie map[string]*Trie

func (m MethodTrie)GetRoot(method string)*Trie  {
	flag:= false
	for _, mt := range methods {
		if method==mt {
			flag = true
			break
		}
	}
	if !flag {
		return nil
	}

	if trie := m[method];trie==nil{
		trie = &Trie{root: &TrieNode{
			value:     "/",
			children:  make(map[string]*TrieNode),
			nt:        root,
			wildChild: false,
		}}

		m[method] = trie

		return trie
	}else {
		return trie
	}



}

type  nodeType int
const (
	root nodeType = iota
	param
	catchAll
)

type Trie struct {
	root *TrieNode
}

func (trie *Trie) Root() *TrieNode {
	return trie.root
}

func (node *TrieNode) InsertChild(paths []string)  {

	for _, path := range paths {

		if node.children[path]==nil{
			n := new(TrieNode)
			n.value = path
			n.wildChild = false
			//if path[:1] == ":" || path[:1] == "*" {
			if path == "*" {
				n.nt = param
				log.Default().Println("当前节点为通配符节点")
				node.wildChild = true
				node.wildChildCount++
			}else{
				log.Default().Println("当前节点",path,"为全匹配节点")
				n.nt = catchAll
			}
			n.children = make(map[string]*TrieNode)

			node.children[path] = n

		}

		node = node.children[path]
	}

	node.wildChild = false

}

// TODO 当前该搜索，还不适用于通配符的情况
func (node *TrieNode)Search(paths []string) bool {

	for _, path := range paths {
		log.Default().Println("当前path",path)
		// *号全匹配
		if node.children["*"]!=nil {
			return true
		}

		// 没有该子节点，
		if node.children[path]==nil {
			// 且当前节点中没有统配符子节点
			//if node.wildChild == false {
			//	return false
			//}else {
			//	log.Default().Println("当前节点下的子节点中有通配符节点")
			//
			//	// 如果i==len(paths)-1,即当前节点为叶子节点且为通配符 则返回true
			//	if i==len(paths)-1 {
			//		return true
			//	}else if i<len(paths)-1 {
			//		// TODO 如果当前i!=len(paths)-1将指针指向路径中下一个不是通配符的节点然后继续匹配
			//		flag :=false
			//
			//		if flag {
			//			continue
			//		}
			//	}
			//
			//	return false
			//}

			return false
		}

		node = node.children[path]
	}

	// 如果最后的节点没有子节点，则为可达
	return len(node.children)==0

}

type TrieNode struct {
	value     string
	children  map[string]*TrieNode
	nt        nodeType
	wildChild bool // 是否有通配符子节点
	wildChildCount int
}

func (tn *TrieNode) getWildChild()[]*TrieNode  {
	tns := make([]*TrieNode,0)
	for _, node := range tn.children {
		if node.nt==param {
			log.Default().Println("提取通配符节点",node.value)
			tns = append(tns,node)
		}
	}

	return tns
}


