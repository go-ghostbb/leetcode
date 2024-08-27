package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProbability(5,
		[][]int{{2, 3}, {1, 2}, {3, 4}, {1, 3}, {1, 4}, {0, 1}, {2, 4}, {0, 4}, {0, 2}},
		[]float64{0.06, 0.26, 0.49, 0.25, 0.2, 0.64, 0.23, 0.21, 0.77},
		0,
		3))

	fmt.Println(maxProbability(500,
		[][]int{{193, 229}, {133, 212}, {224, 465}},
		[]float64{0.91, 0.78, 0.64},
		4,
		364))
}

type Node struct {
	Val  int
	Prob float64
}

type PathMap map[int][]*Node

func NewPathMap(edges [][]int, succProb []float64) PathMap {
	// 建立map
	// ex. a-b => pathMap[a][b] = succProb
	pathMap := make(PathMap)

	for index, edge := range edges {
		// 雙向, 所有a到b跟b到a都要存一份
		if _, ok := pathMap[edge[0]]; !ok {
			pathMap[edge[0]] = make([]*Node, 0)
		}
		if _, ok := pathMap[edge[1]]; !ok {
			pathMap[edge[1]] = make([]*Node, 0)
		}

		pathMap[edge[0]] = append(pathMap[edge[0]], &Node{Val: edge[1], Prob: succProb[index]})
		pathMap[edge[1]] = append(pathMap[edge[1]], &Node{Val: edge[0], Prob: succProb[index]})
	}

	// 將每個map裡面的node進行排序
	for _, nodes := range pathMap {
		sort.SliceStable(nodes, func(i, j int) bool {
			return nodes[i].Prob > nodes[j].Prob
		})
	}

	return pathMap
}

type Queue []*Node

func (q *Queue) Enqueue(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) Dequeue() *Node {
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *Queue) Len() int {
	return len(*q)
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// 使用BFS解題

	// 建立map
	pathMap := NewPathMap(edges, succProb)

	// 建立queue
	queue := make(Queue, 0)

	// 記錄每個節點的最大路徑
	result := make([]float64, n)
	result[start_node] = 1.0

	queue.Enqueue(&Node{start_node, 1.0})

	for queue.Len() > 0 {
		node := queue.Dequeue()

		for _, next := range pathMap[node.Val] {
			if node.Prob*next.Prob <= result[next.Val] {
				// 如果路徑比之前還小, 直接跳過
				continue
			}
			result[next.Val] = node.Prob * next.Prob

			queue.Enqueue(&Node{next.Val, node.Prob * next.Prob})
		}
	}

	return result[end_node]
}
