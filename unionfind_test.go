package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	u := New(10)
	u.Union(1, 2)
	u.Union(3, 4)
	if u.Count() != 8 {
		t.Errorf("Union failed")
	}
}

func TestConnected(t *testing.T) {
	u := New(5)
	u.Union(3, 4)
	if !u.Connected(3, 4) {
		t.Errorf("Not connected")
	}
}
