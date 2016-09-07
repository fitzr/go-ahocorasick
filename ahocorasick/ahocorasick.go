package ahocorasick

type AhoCorasick struct {
	root *node
}

type node struct {
	children map[rune]*node
	depth    int
	backward *node
	hit      bool
}

func New(keywords []string) *AhoCorasick {
	a := AhoCorasick{root: newNode(0)}
	a.createTrie(keywords)
	a.createBackward()
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

func (a *AhoCorasick) createBackward() {
	for k, v := range a.root.children {
		a.walkCreateBackward(v, []rune{k})
	}
}

func (a *AhoCorasick) walkCreateBackward(n *node, text []rune) {
	n.backward = a.backwardMatchNode(text)
	for k, v := range n.children {
		a.walkCreateBackward(v, append(text, k))
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
	runes := []rune(text)
	result := make([][]int, 0)
	n := a.root
	i := 0

	check := func(node *node) {
		if node.hit {
			result = append(result, []int{i - node.depth, node.depth})
		}
	}

	for i < len(runes) {
		check(n)
		child, ok := n.children[runes[i]]
		if ok {
			for n != a.root {
				n = n.backward
				check(n)
			}
			n = child
			i++
		} else if n != a.root {
			n = n.backward
		} else {
			i++
		}
	}

	return result
}
