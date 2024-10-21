package goyamlv3x

import (
	"github.com/yyle88/done"
	"gopkg.in/yaml.v3"
)

func ChangeYamlFieldValue(data []byte, route []string, change func(*yaml.Node)) []byte {
	var rootNode yaml.Node
	// 把内容转换到对象里
	done.Done(yaml.Unmarshal(data, &rootNode))
	// 搜到内容且设置内容
	if subNode := SearchSubNode(rootNode, route); subNode != nil {
		change(subNode)
	}
	// 把对象再转换回内容
	return done.VAE(yaml.Marshal(&rootNode)).Nice()
}

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
