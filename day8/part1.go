package day8

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/izzanzahrial/aoc2025/util"
)

// A node in 3D space.
type node struct {
	rawKey  [][]byte // original CSV-split bytes from input
	x, y, z int
}

// An edge connecting two nodes with a Euclidean distance.
type edge struct {
	nodeAIndex int
	nodeBIndex int
	distance   float64
}

// Disjoint-set Union (DSU) structure for maintaining connected components.
// parent[i]  → parent of node i (or itself if root)
// size[i]    → size of the tree whose root is i
type dsu struct {
	parent []int
	size   []int
}

// Create a DSU with n independent sets (each element is its own root).
func newDSU(n int) *dsu {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i // each node is its own parent
		size[i] = 1   // initial component size is 1
	}
	return &dsu{parent, size}
}

// Find the root of x to be used in union operations.
func (d *dsu) find(x int) int {
	if d.parent[x] != x {
		// Path compression: rewrite parent[x] to point directly to the root.
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

// Union merges the sets of nodes a and b.
// Uses union-by-size to optimize tree depth.
func (d *dsu) union(a, b int) {
	rootA := d.find(a)
	rootB := d.find(b)

	// If the roots are the same, nothing to do
	// since they are located in the same circuit
	if rootA == rootB {
		return
	}

	// Always attach smaller circuit under larger circuit.
	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	d.parent[rootB] = rootA
	d.size[rootA] += d.size[rootB]
}

func distance(a, b node) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Part1() (int, error) {
	input, err := util.Read("day8", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %w", err)
	}

	// Parse lines into nodes.
	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	n := len(lines)
	nodes := make([]node, n)

	for i, line := range lines {
		parts := bytes.Split(line, []byte(","))
		x, _ := strconv.Atoi(string(parts[0]))
		y, _ := strconv.Atoi(string(parts[1]))
		z, _ := strconv.Atoi(string(parts[2]))

		nodes[i] = node{
			rawKey: parts,
			x:      x,
			y:      y,
			z:      z,
		}
	}

	// Construct all pairwise edges (complete graph).
	edges := make([]edge, 0, n*n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{
				nodeAIndex: i,
				nodeBIndex: j,
				distance:   distance(nodes[i], nodes[j]),
			})
		}
	}

	// Sort edges by ascending distance.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	// Use DSU to union the first 1000 nearest edges.
	// This grows clusters based on proximity.
	limit := 1000
	if limit > len(edges) {
		limit = len(edges)
	}

	ds := newDSU(n)

	for i := 0; i < limit; i++ {
		e := edges[i]
		ds.union(e.nodeAIndex, e.nodeBIndex)
	}

	// Count the size of each connected component.
	componentSize := make(map[int]int)

	for i := 0; i < n; i++ {
		root := ds.find(i)
		componentSize[root]++
	}

	// Extract component sizes into a slice.
	sizes := make([]int, 0, len(componentSize))
	for _, sz := range componentSize {
		sizes = append(sizes, sz)
	}

	// We need at least 3 distinct components.
	if len(sizes) < 3 {
		return 0, fmt.Errorf("not enough components after 1000 unions")
	}

	// Pick three largest components.
	sort.Ints(sizes)
	a := sizes[len(sizes)-1]
	b := sizes[len(sizes)-2]
	c := sizes[len(sizes)-3]

	return a * b * c, nil
}
