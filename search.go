package goyamlv3x

import (
	"gopkg.in/yaml.v3"
)

func SearchSubNode(rootNode yaml.Node, route []string) *yaml.Node {
	for _, node := range rootNode.Content {
		if node.Kind == yaml.MappingNode {
			if subNode := RecurseSearch(node, route); subNode != nil {
				return subNode
			}
		}
	}
	return nil
}

func RecurseSearch(node *yaml.Node, route []string) *yaml.Node {
	if len(route) == 0 {
		return node
	}
	for index := 0; index < len(node.Content)-1; index += 2 {
		if node.Content[index].Value == route[0] {
			subNode := node.Content[index+1]

			return RecurseSearch(subNode, route[1:])
		}
	}
	return nil
}
