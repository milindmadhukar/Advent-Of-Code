package day7

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day7 struct {
	data            []string
	startTime       time.Time
	nodes           map[string]*Node
	root            *Node
	cummulativeSums map[*Node]int
}

type Node struct {
	name     string
	weight   int
	children []*Node
}

func (d *day7) Part1() any {
	parents := make(map[*Node]*Node)

	for _, node := range d.nodes {
		for _, child := range node.children {
			parents[child] = node
		}
	}

	for name, node := range d.nodes {
		if parents[node] == nil {
			d.root = node
			return name
		}
	}
	return ""
}

func (d *day7) Traverse(node *Node) int {
	if len(node.children) == 0 {
		d.cummulativeSums[node] = node.weight
		return node.weight
	}

	sum := 0
	for _, child := range node.children {
		sum += d.Traverse(child)
	}

	sum += node.weight

	d.cummulativeSums[node] = sum

	return d.cummulativeSums[node]
}

func (d *day7) Part2() any {
	d.Traverse(d.root)

	currentNode := d.root

	var siblings []*Node

	for {
		childrenWeights := utils.Map(func(n *Node) int { return d.cummulativeSums[n] }, currentNode.children)
		wrongWeightFound := false

		for _, child := range currentNode.children {
			if utils.CountOf(childrenWeights, d.cummulativeSums[child]) == 1 {
				wrongWeightFound = true
				siblings = currentNode.children
				currentNode = child
				break
			}
		}

		if !wrongWeightFound {
			break
		}
	}

	var difference int

	for _, sibling := range siblings {
		if sibling != currentNode {
			difference = d.cummulativeSums[currentNode] - d.cummulativeSums[sibling]
			break
		}
	}

	return currentNode.weight - difference
}

func (d *day7) createOrUpdateNode(name string, weight int) *Node {
	if n, ok := d.nodes[name]; ok {
		if weight != -1 {
			n.weight = weight
		}
		return n
	}
	n := &Node{name: name, weight: weight}
	d.nodes[name] = n
	return n
}

func Solve() *day7 {

	data, err := utils.GetInputDataFromAOC(2017, 6)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day7/example.txt")

	startTime := time.Now()

	splitData := utils.GetSplitData(data, " -> ")

	d := day7{
		data:            data,
		startTime:       startTime,
		nodes:           make(map[string]*Node),
		cummulativeSums: make(map[*Node]int),
	}

	nameHash := func(name string) string {
		return name
	}

	for _, line := range splitData {
		name_weight := strings.Split(line[0], " ")
		name := name_weight[0]
		weightStr := strings.Trim(name_weight[1], "()")
		weight, _ := strconv.Atoi(weightStr)
		n := d.createOrUpdateNode(name, weight)
		if len(line) == 2 {
			childrenNames := strings.Split(line[1], ", ")
			for _, childrenName := range childrenNames {
				n.children = append(n.children, d.createOrUpdateNode(childrenName, -1))
			}
		}
	}

	// Just to visualize graph
	g := graph.New(nameHash, graph.Directed())

	for name, n := range d.nodes {
		g.AddVertex(name,
			graph.VertexAttribute(
				"wt",
				strconv.Itoa(n.weight),
			),
		)
	}

	for name, n := range d.nodes {
		for _, child := range n.children {
			g.AddEdge(name, child.name)
		}
	}

	file, _ := os.Create("day7/graph.gv")
	_ = draw.DOT(g, file)

	return &d
}

func (d day7) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
