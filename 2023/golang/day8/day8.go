package day8

import (
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day8 struct {
	data                  []string
	root                  *node
	nodes                 map[string]*node
	directionInstructions string
	startTime             time.Time
}

type node struct {
	name  string
	left  *node
	right *node
}

func getNodeNameLeftRightFromName(name string, data []string) (string, string, string) {
	for _, line := range data {
		node_left_right := strings.Split(line, " = ")
		nodeName := node_left_right[0]
		if nodeName == name {
			nodeValue := strings.Split(node_left_right[1][1:len(node_left_right[1])-1], ", ")
			return nodeName, nodeValue[0], nodeValue[1]
		}
	}
	return "", "", ""
}

func traverse(root string, data []string, nodes map[string]*node) *node {
	if currentNode, ok := nodes[root]; ok {
		return currentNode
	}

	nodeName, nodeLeft, nodeRight := getNodeNameLeftRightFromName(root, data)
	newNode := &node{
		name: nodeName,
	}
	nodes[nodeName] = newNode

	newNode.left = traverse(nodeLeft, data, nodes)
	newNode.right = traverse(nodeRight, data, nodes)

	return traverse(root, data, nodes)
}

func (d day8) Part1() any {
	steps := 0

	for currentNode := d.root; currentNode.name != "ZZZ"; {
		currentInstruction := d.directionInstructions[steps%len(d.directionInstructions)]
		if currentInstruction == 'L' {
			currentNode = currentNode.left
		} else if currentInstruction == 'R' {
			currentNode = currentNode.right
		}
		steps++
	}

	return steps
}

func (d day8) Part2() any {
	var concernedNodes []*node
	for _, node := range d.nodes {
		if strings.HasSuffix(node.name, "A") {
			concernedNodes = append(concernedNodes, node)
		}
	}

	var minimumStepsPerNode []int

	// Calculate minimum steps per node to reach one that ends with Z
	for _, givenNode := range concernedNodes {
		nodeSteps := 0
		for currentNode := givenNode; !strings.HasSuffix(currentNode.name, "Z"); {
			currentInstruction := d.directionInstructions[nodeSteps%len(d.directionInstructions)]
			if currentInstruction == 'L' {
				currentNode = currentNode.left
			} else if currentInstruction == 'R' {
				currentNode = currentNode.right
			}
			nodeSteps++
		}
		minimumStepsPerNode = append(minimumStepsPerNode, nodeSteps)
	}

	return utils.Lcm(minimumStepsPerNode)
}

func Solve() day8 {
	data, err := utils.GetInputDataFromAOC(2023, 8)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day8/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	directionInstructions := data[0]
	data = data[2:]

	var nodes = make(map[string]*node)

	// Traversing through all nodes as graph is disjoint with multiple components
	for _, line := range data {
		node_left_right := strings.Split(line, " = ")
		nodeName := node_left_right[0]
		traverse(nodeName, data, nodes)
	}

	root := nodes["AAA"]

  // NOTE: Creating the graph visualization using dominikbraun/graph

	// g := graph.New(graph.StringHash, graph.Directed())
	// for _, node := range nodes {
	// 	g.AddVertex(node.name)
	// }
	// for _, node := range nodes {
	// 	if node.left == node.right {
	// 		g.AddEdge(node.name, node.left.name, graph.EdgeAttribute("label", "L-R"), graph.EdgeAttribute("color", "red"))
	// 	}
	// 	if node.left != nil {
	// 		g.AddEdge(node.name, node.left.name, graph.EdgeAttribute("label", "L"))
	// 	}
	// 	if node.right != nil {
	// 		g.AddEdge(node.name, node.right.name, graph.EdgeAttribute("label", "R"))
	// 	}
	// }
	//
	// file, _ := os.Create("day8/graph.gv")
	//
	// _ = draw.DOT(g, file)

	return day8{
		data:                  data,
		directionInstructions: directionInstructions,
		root:                  root,
		nodes:                 nodes,
		startTime:             startTime,
	}
}

func (d day8) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
