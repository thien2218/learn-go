package algorithms

type UnionFind[V Node] struct {
	nodes   map[V]V
	leaders map[V][]V
}

// CONSTRUCTOR

func NewUnionFind[V Node]() *UnionFind[V] {
	uf := new(UnionFind[V])
	uf.nodes = make(map[V]V)
	uf.leaders = make(map[V][]V)

	return uf
}

// GETTERS AND SETTERS

func (uf *UnionFind[V]) GetClusters() map[V]V {
	return uf.nodes
}

func (uf *UnionFind[V]) GetLeaders() map[V][]V {
	return uf.leaders
}

func (uf *UnionFind[V]) SetClusters(nodes map[V]V) {
	uf.nodes = nodes
}

func (uf *UnionFind[V]) SetLeaders(leaders map[V][]V) {
	uf.leaders = leaders
}

// INSERT

func (uf *UnionFind[V]) Insert(vertex1, vertex2 V) {
	if _, exist := uf.nodes[vertex1]; !exist {
		uf.nodes[vertex1] = vertex1
		uf.leaders[vertex1] = []V{vertex1}
	}
	if _, exist := uf.nodes[vertex2]; !exist {
		uf.nodes[vertex2] = vertex2
		uf.leaders[vertex2] = []V{vertex2}
	}
}

// PRIVATE

func (uf *UnionFind[V]) fuse(leader1, leader2 V) {
	for _, item := range uf.leaders[leader1] {
		uf.nodes[item] = leader2
	}

	uf.leaders[leader2] = append(uf.leaders[leader2], uf.leaders[leader1]...)
	delete(uf.leaders, leader1)
}

// PUBLIC

func (uf *UnionFind[V]) Find(vertex, endVertex V) bool {
	return uf.nodes[vertex] == uf.nodes[endVertex]
}

func (uf *UnionFind[V]) Union(vertex, endVertex V) {
	leader1 := uf.nodes[vertex]
	leader2 := uf.nodes[endVertex]

	if len(uf.leaders[leader2]) > len(uf.leaders[leader1]) {
		uf.fuse(leader1, leader2)
	} else {
		uf.fuse(leader2, leader1)
	}
}
