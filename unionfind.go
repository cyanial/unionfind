package unionfind

type UnionFind struct {
	count  int
	parent []int
	size   []int
}

func New(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		count:  n,
		parent: parent,
		size:   size,
	}
}

func (u *UnionFind) Find(x int) int {
	// for u.parent[x] != x {
	// 	x = u.parent[x]
	// }
	// return x

	// 路径压缩
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFind) Union(p, q int) {
	rootP, rootQ := u.Find(p), u.Find(q)
	if rootP == rootQ {
		return
	}
	// u.parent[rootP] = rootQ
	if u.size[rootP] > u.size[rootQ] {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	} else {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	}
	u.count--
}

func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UnionFind) Count() int {
	return u.count
}
