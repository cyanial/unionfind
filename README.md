# Union-Find

## Operatrions:

### Initialization:

```golang
uf := unionfind.New(6)
```

parent: [0, 1, 2, 3, 4, 5]
index:   0  1  2  3  4  5

### Find

find parent of element x

```golang
// var p int
p = uf.Find(x)
```

### Union

Merge two node into one set

```golang
// var p, q int
uf.Union(p, q)
```


