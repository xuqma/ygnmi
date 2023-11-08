// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package nested is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: (devel): (ygot: v0.29.12)
using the following YANG input files:
  - ../../pathgen/testdata/yang/openconfig-simple.yang
  - ../../pathgen/testdata/yang/openconfig-withlistval.yang
  - ../../pathgen/testdata/yang/openconfig-nested.yang

Imported modules were sourced from:
*/
package nested

import (
	oc "github.com/openconfig/ygnmi/internal/uexampleoc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// OpenconfigNested_APath represents the /openconfig-nested/a YANG schema element.
type OpenconfigNested_APath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A]
}

// OpenconfigNested_APathAny represents the wildcard version of the /openconfig-nested/a YANG schema element.
type OpenconfigNested_APathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A]
}

// B (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "b"
//	Path from root:       "/a/b"
func (n *OpenconfigNested_APath) B() *OpenconfigNested_A_BPath {
	ps := &OpenconfigNested_A_BPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"b"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B](
		"OpenconfigNested_A_B",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// B (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "b"
//	Path from root:       "/a/b"
func (n *OpenconfigNested_APathAny) B() *OpenconfigNested_A_BPathAny {
	ps := &OpenconfigNested_A_BPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"b"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B](
		"OpenconfigNested_A_B",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// OpenconfigNested_A_BPath represents the /openconfig-nested/a/b YANG schema element.
type OpenconfigNested_A_BPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B]
}

// OpenconfigNested_A_BPathAny represents the wildcard version of the /openconfig-nested/a/b YANG schema element.
type OpenconfigNested_A_BPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B]
}

// C (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "c"
//	Path from root:       "/a/b/c"
func (n *OpenconfigNested_A_BPath) C() *OpenconfigNested_A_B_CPath {
	ps := &OpenconfigNested_A_B_CPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"c"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C](
		"OpenconfigNested_A_B_C",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// C (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "c"
//	Path from root:       "/a/b/c"
func (n *OpenconfigNested_A_BPathAny) C() *OpenconfigNested_A_B_CPathAny {
	ps := &OpenconfigNested_A_B_CPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"c"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C](
		"OpenconfigNested_A_B_C",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_CPath represents the /openconfig-nested/a/b/c YANG schema element.
type OpenconfigNested_A_B_CPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C]
}

// OpenconfigNested_A_B_CPathAny represents the wildcard version of the /openconfig-nested/a/b/c YANG schema element.
type OpenconfigNested_A_B_CPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C]
}

// D (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "d"
//	Path from root:       "/a/b/c/d"
func (n *OpenconfigNested_A_B_CPath) D() *OpenconfigNested_A_B_C_DPath {
	ps := &OpenconfigNested_A_B_C_DPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"d"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D](
		"OpenconfigNested_A_B_C_D",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// D (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "d"
//	Path from root:       "/a/b/c/d"
func (n *OpenconfigNested_A_B_CPathAny) D() *OpenconfigNested_A_B_C_DPathAny {
	ps := &OpenconfigNested_A_B_C_DPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"d"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D](
		"OpenconfigNested_A_B_C_D",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_DPath represents the /openconfig-nested/a/b/c/d YANG schema element.
type OpenconfigNested_A_B_C_DPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D]
}

// OpenconfigNested_A_B_C_DPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d YANG schema element.
type OpenconfigNested_A_B_C_DPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D]
}

// E (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "e"
//	Path from root:       "/a/b/c/d/e"
func (n *OpenconfigNested_A_B_C_DPath) E() *OpenconfigNested_A_B_C_D_EPath {
	ps := &OpenconfigNested_A_B_C_D_EPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"e"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E](
		"OpenconfigNested_A_B_C_D_E",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// E (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "e"
//	Path from root:       "/a/b/c/d/e"
func (n *OpenconfigNested_A_B_C_DPathAny) E() *OpenconfigNested_A_B_C_D_EPathAny {
	ps := &OpenconfigNested_A_B_C_D_EPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"e"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E](
		"OpenconfigNested_A_B_C_D_E",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_EPath represents the /openconfig-nested/a/b/c/d/e YANG schema element.
type OpenconfigNested_A_B_C_D_EPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E]
}

// OpenconfigNested_A_B_C_D_EPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e YANG schema element.
type OpenconfigNested_A_B_C_D_EPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E]
}

// F (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "f"
//	Path from root:       "/a/b/c/d/e/f"
func (n *OpenconfigNested_A_B_C_D_EPath) F() *OpenconfigNested_A_B_C_D_E_FPath {
	ps := &OpenconfigNested_A_B_C_D_E_FPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"f"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F](
		"OpenconfigNested_A_B_C_D_E_F",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// F (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "f"
//	Path from root:       "/a/b/c/d/e/f"
func (n *OpenconfigNested_A_B_C_D_EPathAny) F() *OpenconfigNested_A_B_C_D_E_FPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_FPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"f"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F](
		"OpenconfigNested_A_B_C_D_E_F",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_FPath represents the /openconfig-nested/a/b/c/d/e/f YANG schema element.
type OpenconfigNested_A_B_C_D_E_FPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F]
}

// OpenconfigNested_A_B_C_D_E_FPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f YANG schema element.
type OpenconfigNested_A_B_C_D_E_FPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F]
}

// G (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "g"
//	Path from root:       "/a/b/c/d/e/f/g"
func (n *OpenconfigNested_A_B_C_D_E_FPath) G() *OpenconfigNested_A_B_C_D_E_F_GPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_GPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"g"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G](
		"OpenconfigNested_A_B_C_D_E_F_G",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// G (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "g"
//	Path from root:       "/a/b/c/d/e/f/g"
func (n *OpenconfigNested_A_B_C_D_E_FPathAny) G() *OpenconfigNested_A_B_C_D_E_F_GPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_GPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"g"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G](
		"OpenconfigNested_A_B_C_D_E_F_G",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_GPath represents the /openconfig-nested/a/b/c/d/e/f/g YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_GPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G]
}

// OpenconfigNested_A_B_C_D_E_F_GPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_GPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G]
}

// H (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "h"
//	Path from root:       "/a/b/c/d/e/f/g/h"
func (n *OpenconfigNested_A_B_C_D_E_F_GPath) H() *OpenconfigNested_A_B_C_D_E_F_G_HPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_HPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"h"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H](
		"OpenconfigNested_A_B_C_D_E_F_G_H",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// H (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "h"
//	Path from root:       "/a/b/c/d/e/f/g/h"
func (n *OpenconfigNested_A_B_C_D_E_F_GPathAny) H() *OpenconfigNested_A_B_C_D_E_F_G_HPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_HPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"h"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H](
		"OpenconfigNested_A_B_C_D_E_F_G_H",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_HPath represents the /openconfig-nested/a/b/c/d/e/f/g/h YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_HPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H]
}

// OpenconfigNested_A_B_C_D_E_F_G_HPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_HPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H]
}

// I (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "i"
//	Path from root:       "/a/b/c/d/e/f/g/h/i"
func (n *OpenconfigNested_A_B_C_D_E_F_G_HPath) I() *OpenconfigNested_A_B_C_D_E_F_G_H_IPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_IPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"i"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// I (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "i"
//	Path from root:       "/a/b/c/d/e/f/g/h/i"
func (n *OpenconfigNested_A_B_C_D_E_F_G_HPathAny) I() *OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"i"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_H_IPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_IPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I]
}

// J (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "j"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_IPath) J() *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"j"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// J (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "j"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny) J() *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"j"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J]
}

// K (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "k"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath) K() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"k"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// K (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "k"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny) K() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"k"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K]
}

// L (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "l"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath) L() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"l"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// L (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "l"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny) L() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"l"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L]
}

// M (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "m"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath) M() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"m"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// M (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "m"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny) M() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"m"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M]
}

// State (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "state"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath) State() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"state"},
			map[string]interface{}{},
			n,
		),
	}
	ps.SingletonQuery = ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// State (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "state"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny) State() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"state"},
			map[string]interface{}{},
			n,
		),
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
	ygnmi.SingletonQuery[string]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
	ygnmi.WildcardQuery[string]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath struct {
	*ygnmi.NodePath
	ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State]
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State]
}

// Foo (leaf):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "foo"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath) Foo() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"foo"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	ps.SingletonQuery = ygnmi.NewSingletonQuery[string](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		false,
		true,
		true,
		false,
		false,
		ps,
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State).Foo
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// Foo (leaf):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "foo"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny) Foo() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny {
	ps := &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"foo"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	ps.WildcardQuery = ygnmi.NewWildcardQuery[string](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		false,
		true,
		true,
		false,
		false,
		ps,
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State).Foo
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// OpenconfigNested_ContainerPath represents the /openconfig-nested/container YANG schema element.
type OpenconfigNested_ContainerPath struct {
	*ygnmi.NodePath
	ygnmi.ConfigQuery[*oc.OpenconfigNested_Container]
}

// OpenconfigNested_ContainerPathAny represents the wildcard version of the /openconfig-nested/container YANG schema element.
type OpenconfigNested_ContainerPathAny struct {
	*ygnmi.NodePath
	ygnmi.WildcardQuery[*oc.OpenconfigNested_Container]
}
