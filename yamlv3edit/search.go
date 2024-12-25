package yamlv3edit

import (
	"strings"

	"gopkg.in/yaml.v3"
)

func LookupSubNode(node *yaml.Node, path string) *yaml.Node {
	return SearchSubNode(node, strings.Split(path, "."))
}

func SearchSubNode(node *yaml.Node, route []string) *yaml.Node {
	for _, sub := range node.Content {
		if sub.Kind == yaml.MappingNode {
			if subNode := RecurseSearch(sub, route); subNode != nil {
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
