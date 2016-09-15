package goahocorasick

// AhoCorasick is the interface that wraps the Match method.
type AhoCorasick interface {
	Match(text string) [][]int
}

type ahocorasick struct {
	root *node
}

type node struct {
	children map[rune]*node
	depth    int
	fail     *node
	hit      bool
}

// New returns new AhoCorasick with keywords.
func New(keywords []string) AhoCorasick {
	a := ahocorasick{root: newNode(0)}
	a.createTrie(keywords)
	a.createFail()
	return &a
}

func newNode(depth int) *node {
	return &node{
		children: map[rune]*node{},
		depth:    depth}
}

func (a *ahocorasick) createTrie(keywords []string) {
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

func (a *ahocorasick) createFail() {
	for k, v := range a.root.children {
		a.walkCreateFail(v, []rune{k})
	}
}

func (a *ahocorasick) walkCreateFail(n *node, text []rune) {
	n.fail = a.backwardMatchNode(text)
	for k, v := range n.children {
		a.walkCreateFail(v, append(text, k))
	}
}

func (a *ahocorasick) backwardMatchNode(text []rune) *node {
	for t := text[1:]; len(t) > 0; t = t[1:] {
		n, ok := a.matchNode(t)
		if ok {
			return n
		}
	}
	return a.root
}

func (a *ahocorasick) matchNode(text []rune) (*node, bool) {
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

// Match keywords and returns index and length in units of rune (utf8).
func (a *ahocorasick) Match(text string) [][]int {
	var result [][]int
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
