package day8

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part2() (int, error) {
	input, err := util.Read("day8", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %w", err)
	}

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

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	// Run Kruskal-style process until only one component remains.
	ds := newDSU(n)
	components := n // start with n isolated nodes

	for _, e := range edges {
		rootA := ds.find(e.nodeAIndex)
		rootB := ds.find(e.nodeBIndex)

		// If they are already connected, skip
		if rootA == rootB {
			continue
		}

		// Merge components
		ds.union(rootA, rootB)
		components--

		// When only one component remains, this is the final required edge
		if components == 1 {
			nodeA := nodes[e.nodeAIndex]
			nodeB := nodes[e.nodeBIndex]
			return nodeA.x * nodeB.x, nil
		}
	}

	return 0, fmt.Errorf("graph never became fully connected")
}
