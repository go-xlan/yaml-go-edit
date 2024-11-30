package goyamlv3up

import (
	"github.com/yyle88/done"
	"gopkg.in/yaml.v3"
)

func ChangeYamlFieldValue(data []byte, route []string, change func(*yaml.Node)) []byte {
	var rootNode yaml.Node
	// 把内容转换到对象里
	done.Done(yaml.Unmarshal(data, &rootNode))
	// 搜到内容且设置内容
	if subNode := SearchSubNode(&rootNode, route); subNode != nil {
		change(subNode)
	}
	// 把对象再转换回内容
	return done.VAE(yaml.Marshal(&rootNode)).Nice()
}

func ModifyYamlFieldValue(data []byte, path string, value string) []byte {
	var rootNode yaml.Node
	// 把内容转换到对象里
	done.Done(yaml.Unmarshal(data, &rootNode))
	// 搜到内容且设置内容
	if subNode := LookupSubNode(&rootNode, path); subNode != nil {
		subNode.SetString(value)
	}
	// 把对象再转换回内容
	return done.VAE(yaml.Marshal(&rootNode)).Nice()
}
