package tree

type nodes []*node

func (ns nodes) Len() int {
	return len(ns)
}

func (ns nodes) Less(i, j int) bool {
	return ns[i].less(ns[j])
}

func (ns nodes) Swap(i, j int) {
	temp := ns[i]
	ns[i] = ns[j]
	ns[j] = temp
	return
}

func search(ns nodes, x *node) int {
	for i, n := range ns {
		if x.equal(n) {
			return i
		}
	}
	return -1
}

func appendNode(ns nodes, news ...*node) nodes {
	ret := make(nodes, len(ns))
	copy(ret, ns)

	for _, n := range news {
		ret = append(ret, n)
	}

	return ret
}

// remove node from the slice of node.
func removeNode(ns nodes, n *node) nodes {
	ret := make(nodes, len(ns))
	copy(ret, ns)

	idx := search(ret, n)
	ret = append(ret[:idx], ret[idx+1:]...)

	return ret
}

type node struct {
	parent   *node
	children nodes

	item Item
	debug string
}

func newNode(i Item) *node {
	return &node{
		item: i,
		debug: i.String(),
	}
}

func (n *node) equal(comp *node) bool {
	if !n.less(comp) && !comp.less(n) {
		return true
	}
	return false
}

func (n *node) less(comp *node) bool {
	return n.item.Less(comp.item)
}

func (n *node) String() string {
	return n.item.String()
}

func isAncestor(parent, child *node) bool {
	if child.parent == nil {
		return false
	}

	if child.parent.equal(parent) {
		return true
	}

	return isAncestor(parent, child.parent)
}
