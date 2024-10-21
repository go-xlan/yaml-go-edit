package goyamlv3x

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestChangeYamlFieldValue(t *testing.T) {
	content := `# 这是注释
# 这是内容
info:
    title: OLD-TITLE
    version: 1.0.0
`
	// modify yaml file change the info.title from "OLD-TITLE" to "NEW-TITLE" return the new content
	newContent := ChangeYamlFieldValue([]byte(content), []string{"info", "title"}, func(node *yaml.Node) {
		node.SetString("NEW-TITLE")
	})
	t.Log(string(newContent))

	resContent := `# 这是注释
# 这是内容
info:
    title: NEW-TITLE
    version: 1.0.0
`
	require.Equal(t, resContent, string(newContent))
}
