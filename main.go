package main

import (
	"bytes"
	_ "embed"
	"html/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const tpl = `// Code generated by protoc-gen-go-grpc-fixture. DO NOT EDIT.

package {{ .GoPackageName }}

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	grpc "google.golang.org/grpc"
	. "google.golang.org/protobuf/types/known/emptypb"
)


{{range .Services}}
type Fixture{{ .GoName }}Client struct {
	baseDir string
}

{{range .Methods}}
func (m *Fixture{{ .Parent.GoName }}Client) {{ .GoName }}(ctx context.Context, in *{{ .Input.GoIdent.GoName }}, opts ...grpc.CallOption) (*{{ .Output.GoIdent.GoName }}, error) {
    f := path.Join(m.baseDir, "{{.GoName}}.json")
    rv := &{{ .Output.GoIdent.GoName }}{}
	// if file f does not exist, create it with and empty return object.
	if _, err := os.Stat(f); os.IsNotExist(err) {
		fp, err := os.Create(f)
		if err != nil {
			return nil, fmt.Errorf("failed to create %s: %w", f, err)
		}
		defer fp.Close()
		enc := json.NewEncoder(fp)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rv); err != nil {
			return nil, fmt.Errorf("failed to encode %s: %w", f, err)
		}
		return rv, nil
	}
	// if file f exists, read it and return its contents.
	fp, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", f, err)
	}
	err = json.NewDecoder(fp).Decode(rv)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", f, err)
	}
	return rv, nil
}
{{end}}


{{end}}
`

func generateSource(file *protogen.File) []byte {
	b := bytes.Buffer{}
	t := template.Must(template.New("fixture").Parse(tpl))
	err := t.Execute(&b, file)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.FilesByPath {
			if !file.Generate {
				continue
			}

			buf := generateSource(file)

			if _, err := plugin.NewGeneratedFile(
				file.GeneratedFilenamePrefix+"_grpc_fixture.pb.go",
				file.GoImportPath,
			).Write(buf); err != nil {
				return err
			}
		}
		return nil
	})
}