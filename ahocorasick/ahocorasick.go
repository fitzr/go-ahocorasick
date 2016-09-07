package ahocorasick

type AhoCorasick struct {
	root *node
}

type node struct {
	children map[rune]*node
	depth    int
	next     *node
	hit      bool
}

func New(keywords []string) *AhoCorasick {
	a := AhoCorasick{root: newNode(0)}
	a.createTrie(keywords)
	a.createNext()
	return &a
}

func newNode(depth int) *node {
	return &node{
		children: map[rune]*node{},
		depth:    depth}
}

func (a *AhoCorasick) createTrie(keywords []string) {
	for _, keyword := range keywords {
		n := a.root
		for _, r := range keyword {
			v, ok := n.children[r]
			if !ok {
				v = newNode(n.depth + 1)
				n.children[r] = v
			}
			n = v
		}
		n.hit = true
	}
}

func (a *AhoCorasick) createNext() {
	a.root.next = a.root
	a.walkCreateNext(a.root.children, make([]rune, 0))
}

func (a *AhoCorasick) walkCreateNext(nodes map[rune]*node, text []rune) {
	for k, n := range nodes {
		text = append(text, k)
		n.next = a.backwardMatchNode(text)
		a.walkCreateNext(n.children, text)
	}
}

func (a *AhoCorasick) backwardMatchNode(text []rune) *node {
	for t := text[1:]; len(t) > 0; t = t[1:] {
		n, ok := a.matchNode(t)
		if ok {
			return n
		}
	}
	return a.root
}

func (a *AhoCorasick) matchNode(text []rune) (*node, bool) {
	n := a.root
	for _, r := range text {
		v, ok := n.children[r]
		if ok {
			n = v
		} else {
			return nil, false
		}
	}
	return n, true
}

func (a *AhoCorasick) Match(text string) [][]int {
	result := make([][]int, 0)
	n := a.root
	pos := 0
	for _, r := range text {
		v, ok := n.children[r]
		if ok {
			n = v
			if v.hit {
				result = append(result, []int{pos - v.depth, v.depth})
			}
		} else {
			n = n.next
		}

		pos++
	}
	return result
}
