package goyamlv3x

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"gopkg.in/yaml.v3"
)

func TestLookupSubNode(t *testing.T) {
	content := `# 这是注释
# 这是内容
info:
    title: DOC-TITLE
    version: 1.0.0
`
	data := []byte(content)

	var rootNode yaml.Node
	// 把内容转换到对象里
	done.Done(yaml.Unmarshal(data, &rootNode))
	// 搜到内容且设置内容
	subNode := LookupSubNode(&rootNode, "info.title")

	require.NotNil(t, subNode)
	t.Log(subNode.Value)
	require.Equal(t, "DOC-TITLE", subNode.Value)
}

func TestSearchSubNode(t *testing.T) {
	content := `# 这是注释
# 这是内容
info:
    title: DOC-TITLE
    version: 1.0.0
`
	data := []byte(content)

	var rootNode yaml.Node
	// 把内容转换到对象里
	done.Done(yaml.Unmarshal(data, &rootNode))
	// 搜到内容且设置内容
	subNode := SearchSubNode(&rootNode, []string{"info", "version"})

	require.NotNil(t, subNode)
	t.Log(subNode.Value)
	require.Equal(t, "1.0.0", subNode.Value)
}
