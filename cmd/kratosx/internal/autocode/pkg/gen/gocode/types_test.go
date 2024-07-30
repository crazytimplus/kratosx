package gocode

import (
	"encoding/json"
	"fmt"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg/gen/types"
	"os"
	"testing"
)

func TestGenTypes(t *testing.T) {
	initTable := func() *types.Table {
		var table types.Table
		content, err := os.ReadFile("internal/autocode/pkg/gen/auto.json")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(content, &table); err != nil {
			panic(err)
		}
		return &table
	}

	builder := NewTypesBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.ScanTypes()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}