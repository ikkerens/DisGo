package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"

	"github.com/slf4go/logger"
	"os"
	"sort"
)

type internalField struct {
	Parent  *internalType
	Name    string
	TypeStr string
}

type internalType struct {
	Name     string
	Exported string
	Fields   []internalField
}

func determineType(field ast.Expr) (typeStr string, err error) {
	switch v := field.(type) {
	case *ast.Ident:
		typeStr = v.Name
	case *ast.StarExpr:
		subType, err := determineType(v.X)
		if err != nil {
			return "", err
		}
		typeStr = "*" + subType
	case *ast.SelectorExpr:
		subType, err := determineType(v.X)
		if err != nil {
			return "", err
		}
		typeStr = subType + "." + v.Sel.Name
	case *ast.ArrayType:
		subType, err := determineType(v.Elt)
		if err != nil {
			return "", err
		}
		typeStr = "[]" + subType
	default:
		err = errors.New(fmt.Sprintf("Unknown field type (%T): %+v", v, v))
	}

	return
}

func main() {
	file, err := parser.ParseFile(token.NewFileSet(), "model.go", nil, 0)
	if err != nil {
		logger.ErrorE(err)
		return
	}

	var types = make([]internalType, 0)
	for name, object := range file.Scope.Objects {
		if object.Kind == ast.Typ {
			typ := object.Decl.(*ast.TypeSpec)

			if len(typ.Name.Name) < 8 || typ.Name.Name[:8] != "internal" {
				continue
			}

			structDef, ok := typ.Type.(*ast.StructType)

			if ok {
				typeDef := internalType{name, name[8:], make([]internalField, 0)}
				logger.Infof("Creating API struct for %s with name %s.", typeDef.Name, typeDef.Exported)
				for _, field := range structDef.Fields.List {
					typeStr, err := determineType(field.Type)
					if err != nil {
						logger.ErrorE(err)
						continue
					}

					iField := internalField{&typeDef, field.Names[0].Name, typeStr}
					typeDef.Fields = append(typeDef.Fields, iField)
					logger.Infof("Adding func %s() with return type %s to %s.", iField.Name, iField.TypeStr, iField.Parent.Exported)
				}

				types = append(types, typeDef)
			}
		}
	}

	sort.SliceStable(types, func(i, j int) bool {
		return types[i].Name < types[j].Name
	})

	logger.Infof("Generating GO file")
	var tpl = template.Must(template.New("apimodel").Parse(`package disgo

		// Warning: This file has been automatically generated by generate/apimodel/main.go
		// Do NOT make changes here, instead adapt model.go and run go generate

		import (
			"encoding/json"
			"time"
		)

		{{range .}}
			type {{.Exported}} struct {
				discordObject *{{.Name}}
			}

			func (s *{{.Exported}}) MarshalJSON() ([]byte, error) {
				return json.Marshal(s.discordObject)
			}

			func (s *{{.Exported}}) UnmarshalJSON(b []byte) error {
				s.discordObject = &{{.Name}}{}
				return json.Unmarshal(b, &s.discordObject)
			}

			{{range .Fields}}
				func (s *{{.Parent.Exported}}) {{.Name}}() {{.TypeStr}} {
					return s.discordObject.{{.Name}}
				}
			{{end}}
		{{end}}
	`))

	var result []byte
	var buf bytes.Buffer
	err = tpl.Execute(&buf, types)

	if err == nil {
		var formatted []byte
		logger.Infof("Formatting GO file")
		formatted, err = format.Source(buf.Bytes())
		result = formatted
	}

	if err != nil {
		logger.ErrorE(err)
		os.Exit(1)
	}

	ioutil.WriteFile("apimodel.go", result, 0644)
}
