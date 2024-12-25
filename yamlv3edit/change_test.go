package yamlv3edit_test

import (
	"testing"

	"github.com/go-xlan/yaml-go-edit/yamlv3edit"
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
	newContent := yamlv3edit.ChangeYamlFieldValue([]byte(content), []string{"info", "title"}, func(node *yaml.Node) {
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

func TestModifyYamlFieldValue(t *testing.T) {
	content := `# 这是注释
# 这是内容
info:
    title: DOC-TITLE
    version: 1.0.0
`
	// modify yaml file change the info.title from "OLD-TITLE" to "NEW-TITLE" return the new content
	newContent := yamlv3edit.ModifyYamlFieldValue([]byte(content), "info.version", "1.0.1")
	t.Log(string(newContent))

	resContent := `# 这是注释
# 这是内容
info:
    title: DOC-TITLE
    version: 1.0.1
`
	require.Equal(t, resContent, string(newContent))
}
