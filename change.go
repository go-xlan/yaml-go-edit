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
