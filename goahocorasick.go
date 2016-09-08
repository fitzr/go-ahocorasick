package goahocorasick

type AhoCorasick struct {
	root *node
}

type node struct {
	children map[rune]*node
	depth    int
	fail     *node
	hit      bool
}

func New(keywords []string) *AhoCorasick {
	a := AhoCorasick{root: newNode(0)}
	a.createTrie(keywords)
	a.createFail()
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

func (a *AhoCorasick) createFail() {
	for k, v := range a.root.children {
		a.walkCreateFail(v, []rune{k})
	}
}

func (a *AhoCorasick) walkCreateFail(n *node, text []rune) {
	n.fail = a.backwardMatchNode(text)
	for k, v := range n.children {
		a.walkCreateFail(v, append(text, k))
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
	i := 0

	for _, r := range text {
	L:
		if n.hit {
			result = append(result, []int{i - n.depth, n.depth})
		}
		child, ok := n.children[r]
		if ok {
			for n != a.root {
				n = n.fail
				if n.hit {
					result = append(result, []int{i - n.depth, n.depth})
				}
			}
			n = child
			i++
		} else if n == a.root {
			i++
		} else {
			n = n.fail
			goto L
		}
	}

	return result
}
