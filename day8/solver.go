package day8

import (
	"log"
	"regexp"
	"slices"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

type BinaryTree struct {
	head *Node
}

func partTwo(filepath string) {
	lines := util.GetLines("day8/input.txt")
	var directions string
	nodes := make(map[string]*Node)
	re := regexp.MustCompile(`\w\w\w`)
	for i, l := range lines {
		if i == 0 {
			directions = l
		}

		if i > 1 {
			regexdString := re.FindAllString(l, -1)
			parent, left, right := regexdString[0], regexdString[1], regexdString[2]
			if _, ok := nodes[parent]; !ok {
				nodes[parent] = NewNode(parent)
			}

			if _, ok := nodes[left]; !ok {
				nodes[left] = NewNode(left)
			}

			if _, ok := nodes[right]; !ok {
				nodes[right] = NewNode(right)
			}
		}
	}

	nodesEndingInA := make([]*Node, 0)

	for i, l := range lines {
		if i <= 1 {
			continue
		}
		regexdString := re.FindAllString(l, -1)
		parent, left, right := regexdString[0], regexdString[1], regexdString[2]
		if parent[2] == 'A' {
			nodesEndingInA = append(nodesEndingInA, nodes[parent])
		}
		nodes[parent].AddLeft(nodes[left])
		nodes[parent].AddRight(nodes[right])
	}
	correspondingNodes := []*Node{}
	for _, n := range nodesEndingInA {
		correspondingNodes = append(correspondingNodes, nodes[n.value[0:2]+"Z"])
	}

	reachedZIn := []int{}
	steps := 0
	// allEndInZ := false
	for len(nodesEndingInA) != 0 {
		for _, d := range directions {
			steps += 1
			for _, node := range nodesEndingInA {
				if d == 'L' {
					*node = *node.left
				}
				if d == 'R' {
					*node = *node.right
				}
			}

			for i, node := range nodesEndingInA {
				if node.value[2] == 'Z' {
					nodesEndingInA = slices.Delete(nodesEndingInA, i, i+1)
					reachedZIn = append(reachedZIn, steps)
					break
				}
			}
		}
	}

	log.Printf("Day eight, part two answer: %v", lcmSlice(reachedZIn))
}

func lcmSlice(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = LCM(result, number)
	}
	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func partOne(filepath string) {
	lines := util.GetLines(filepath)
	var directions string
	nodes := make(map[string]*Node)
	re := regexp.MustCompile(`\w\w\w`)
	for i, l := range lines {
		if i == 0 {
			directions = l
		}

		if i > 1 {
			regexdString := re.FindAllString(l, -1)
			parent, left, right := regexdString[0], regexdString[1], regexdString[2]
			if _, ok := nodes[parent]; !ok {
				nodes[parent] = NewNode(parent)
			}

			if _, ok := nodes[left]; !ok {
				nodes[left] = NewNode(left)
			}

			if _, ok := nodes[right]; !ok {
				nodes[right] = NewNode(right)
			}
		}
	}

	for i, l := range lines {
		if i <= 1 {
			continue
		}
		regexdString := re.FindAllString(l, -1)
		parent, left, right := regexdString[0], regexdString[1], regexdString[2]
		nodes[parent].AddLeft(nodes[left])
		nodes[parent].AddRight(nodes[right])
	}

	root := nodes["AAA"]
	steps := 0
	currentNode := root
	for currentNode.value != "ZZZ" {
		for _, d := range directions {
			if d == 'L' {
				currentNode = currentNode.left
				steps += 1
			}
			if d == 'R' {
				currentNode = currentNode.right
				steps += 1
			}
		}
	}

	log.Printf("Day eight, part one answer: %v", steps)
}

func Solve(filepath string) {
	partOne(filepath)
	partTwo(filepath)
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func (n *Node) AddLeft(left *Node) {
	n.left = left
}

func (n *Node) AddRight(right *Node) {
	n.right = right
}
