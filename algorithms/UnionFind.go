package algorithms

type UnionFind[V Node] struct {
	clusters map[V]V
	leaders  map[V][]V
}

// CONSTRUCTOR

func NewUnionFind[V Node]() *UnionFind[V] {
	uf := new(UnionFind[V])
	uf.clusters = make(map[V]V)
	uf.leaders = make(map[V][]V)

	return uf
}

// GETTERS AND SETTERS

func (uf *UnionFind[V]) GetClusters() map[V]V {
	return uf.clusters
}

func (uf *UnionFind[V]) GetLeaders() map[V][]V {
	return uf.leaders
}

func (uf *UnionFind[V]) SetClusters(clusters map[V]V) {
	uf.clusters = clusters
}

func (uf *UnionFind[V]) SetLeaders(leaders map[V][]V) {
	uf.leaders = leaders
}

// INSERT

func (uf *UnionFind[V]) Insert(vertex1, vertex2 V) {
	if _, exist := uf.clusters[vertex1]; !exist {
		uf.clusters[vertex1] = vertex1
		uf.leaders[vertex1] = []V{vertex1}
	}
	if _, exist := uf.clusters[vertex2]; !exist {
		uf.clusters[vertex2] = vertex2
		uf.leaders[vertex2] = []V{vertex2}
	}
}

// PRIVATE

func (uf *UnionFind[V]) fuse(vertex, endVertex V) {
	for _, item := range uf.leaders[vertex] {
		uf.clusters[item] = endVertex
	}

	uf.leaders[endVertex] = append(uf.leaders[endVertex], uf.leaders[vertex]...)
	delete(uf.leaders, vertex)
}

// PUBLIC

func (uf *UnionFind[V]) Find(vertex, endVertex V) bool {
	return uf.clusters[vertex] == uf.clusters[endVertex]
}

func (uf *UnionFind[V]) Union(vertex, endVertex V) {
	leader1 := uf.clusters[vertex]
	leader2 := uf.clusters[endVertex]

	if len(uf.leaders[leader2]) > len(uf.leaders[leader1]) {
		uf.fuse(leader1, leader2)
	} else {
		uf.fuse(leader2, leader1)
	}
}
