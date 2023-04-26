package tokens

import "github.com/databricks/bricks/libs/cmdio"

func init() {
	listCmd.Annotations["template"] = cmdio.Heredoc(`
	{{white "ID"}}	{{white "Expiry time"}}	{{white "Comment"}}
	{{range .}}{{.TokenId|green}}	{{white "%d" .ExpiryTime}}	{{.Comment|white}}
	{{end}}`)
}
