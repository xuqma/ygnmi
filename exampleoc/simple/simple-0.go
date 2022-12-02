// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package simple is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by ygnmi version: (devel): (ygot: v0.25.4)
using the following YANG input files:
  - ../pathgen/testdata/yang/openconfig-simple.yang
  - ../pathgen/testdata/yang/openconfig-withlistval.yang
  - ../pathgen/testdata/yang/openconfig-nested.yang

Imported modules were sourced from:
*/
package simple

import (
	"reflect"

	oc "github.com/openconfig/ygnmi/exampleoc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// ParentPath represents the /openconfig-simple/parent YANG schema element.
type ParentPath struct {
	*ygnmi.NodePath
}

// ParentPathAny represents the wildcard version of the /openconfig-simple/parent YANG schema element.
type ParentPathAny struct {
	*ygnmi.NodePath
}

// Child (container):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "child"
//	Path from root:       "/parent/child"
func (n *ParentPath) Child() *Parent_ChildPath {
	return &Parent_ChildPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"child"},
			map[string]interface{}{},
			n,
		),
	}
}

// Child (container):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "child"
//	Path from root:       "/parent/child"
func (n *ParentPathAny) Child() *Parent_ChildPathAny {
	return &Parent_ChildPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"child"},
			map[string]interface{}{},
			n,
		),
	}
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// State returns a Query that can be used in gNMI operations.
func (n *ParentPath) State() ygnmi.SingletonQuery[*oc.Parent] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Parent](
		"Parent",
		true,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *ParentPathAny) State() ygnmi.WildcardQuery[*oc.Parent] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Parent](
		"Parent",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *ParentPath) Config() ygnmi.ConfigQuery[*oc.Parent] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Parent](
		"Parent",
		false,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *ParentPathAny) Config() ygnmi.WildcardQuery[*oc.Parent] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Parent](
		"Parent",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Parent_Child_FivePath represents the /openconfig-simple/parent/child/state/five YANG schema element.
type Parent_Child_FivePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_FivePathAny represents the wildcard version of the /openconfig-simple/parent/child/state/five YANG schema element.
type Parent_Child_FivePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
func (n *Parent_ChildPath) State() ygnmi.SingletonQuery[*oc.Parent_Child] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Parent_Child](
		"Parent_Child",
		true,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Parent_ChildPathAny) State() ygnmi.WildcardQuery[*oc.Parent_Child] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Parent_Child](
		"Parent_Child",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Parent_ChildPath) Config() ygnmi.ConfigQuery[*oc.Parent_Child] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Parent_Child](
		"Parent_Child",
		false,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Parent_ChildPathAny) Config() ygnmi.WildcardQuery[*oc.Parent_Child] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Parent_Child](
		"Parent_Child",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/five"
//	Path from root:       "/parent/child/state/five"
func (n *Parent_Child_FivePath) State() ygnmi.SingletonQuery[float32] {
	return ygnmi.NewLeafSingletonQuery[float32](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "five"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (float32, bool) {
			ret := gs.(*oc.Parent_Child).Five
			return ygot.BinaryToFloat32(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/five"
//	Path from root:       "/parent/child/state/five"
func (n *Parent_Child_FivePathAny) State() ygnmi.WildcardQuery[float32] {
	return ygnmi.NewLeafWildcardQuery[float32](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "five"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (float32, bool) {
			ret := gs.(*oc.Parent_Child).Five
			return ygot.BinaryToFloat32(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/five"
//	Path from root:       "/parent/child/config/five"
func (n *Parent_Child_FivePath) Config() ygnmi.ConfigQuery[float32] {
	return ygnmi.NewLeafConfigQuery[float32](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "five"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (float32, bool) {
			ret := gs.(*oc.Parent_Child).Five
			return ygot.BinaryToFloat32(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/five"
//	Path from root:       "/parent/child/config/five"
func (n *Parent_Child_FivePathAny) Config() ygnmi.WildcardQuery[float32] {
	return ygnmi.NewLeafWildcardQuery[float32](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "five"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (float32, bool) {
			ret := gs.(*oc.Parent_Child).Five
			return ygot.BinaryToFloat32(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/four"
//	Path from root:       "/parent/child/state/four"
func (n *Parent_Child_FourPath) State() ygnmi.SingletonQuery[oc.Binary] {
	return ygnmi.NewLeafSingletonQuery[oc.Binary](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "four"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.Binary, bool) {
			ret := gs.(*oc.Parent_Child).Four
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/four"
//	Path from root:       "/parent/child/state/four"
func (n *Parent_Child_FourPathAny) State() ygnmi.WildcardQuery[oc.Binary] {
	return ygnmi.NewLeafWildcardQuery[oc.Binary](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "four"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.Binary, bool) {
			ret := gs.(*oc.Parent_Child).Four
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/four"
//	Path from root:       "/parent/child/config/four"
func (n *Parent_Child_FourPath) Config() ygnmi.ConfigQuery[oc.Binary] {
	return ygnmi.NewLeafConfigQuery[oc.Binary](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "four"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.Binary, bool) {
			ret := gs.(*oc.Parent_Child).Four
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/four"
//	Path from root:       "/parent/child/config/four"
func (n *Parent_Child_FourPathAny) Config() ygnmi.WildcardQuery[oc.Binary] {
	return ygnmi.NewLeafWildcardQuery[oc.Binary](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "four"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.Binary, bool) {
			ret := gs.(*oc.Parent_Child).Four
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/one"
//	Path from root:       "/parent/child/state/one"
func (n *Parent_Child_OnePath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Parent_Child",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "one"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Parent_Child).One
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/one"
//	Path from root:       "/parent/child/state/one"
func (n *Parent_Child_OnePathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Parent_Child",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "one"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Parent_Child).One
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/one"
//	Path from root:       "/parent/child/config/one"
func (n *Parent_Child_OnePath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewLeafConfigQuery[string](
		"Parent_Child",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"config", "one"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Parent_Child).One
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/one"
//	Path from root:       "/parent/child/config/one"
func (n *Parent_Child_OnePathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Parent_Child",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"config", "one"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Parent_Child).One
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/six"
//	Path from root:       "/parent/child/state/six"
func (n *Parent_Child_SixPath) State() ygnmi.SingletonQuery[[]float32] {
	return ygnmi.NewLeafSingletonQuery[[]float32](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "six"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]float32, bool) {
			ret := gs.(*oc.Parent_Child).Six
			return binarySliceToFloatSlice(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/six"
//	Path from root:       "/parent/child/state/six"
func (n *Parent_Child_SixPathAny) State() ygnmi.WildcardQuery[[]float32] {
	return ygnmi.NewLeafWildcardQuery[[]float32](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "six"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]float32, bool) {
			ret := gs.(*oc.Parent_Child).Six
			return binarySliceToFloatSlice(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/six"
//	Path from root:       "/parent/child/config/six"
func (n *Parent_Child_SixPath) Config() ygnmi.ConfigQuery[[]float32] {
	return ygnmi.NewLeafConfigQuery[[]float32](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "six"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]float32, bool) {
			ret := gs.(*oc.Parent_Child).Six
			return binarySliceToFloatSlice(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/six"
//	Path from root:       "/parent/child/config/six"
func (n *Parent_Child_SixPathAny) Config() ygnmi.WildcardQuery[[]float32] {
	return ygnmi.NewLeafWildcardQuery[[]float32](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "six"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]float32, bool) {
			ret := gs.(*oc.Parent_Child).Six
			return binarySliceToFloatSlice(ret), !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/three"
//	Path from root:       "/parent/child/state/three"
func (n *Parent_Child_ThreePath) State() ygnmi.SingletonQuery[oc.E_Child_Three] {
	return ygnmi.NewLeafSingletonQuery[oc.E_Child_Three](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "three"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Child_Three, bool) {
			ret := gs.(*oc.Parent_Child).Three
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/three"
//	Path from root:       "/parent/child/state/three"
func (n *Parent_Child_ThreePathAny) State() ygnmi.WildcardQuery[oc.E_Child_Three] {
	return ygnmi.NewLeafWildcardQuery[oc.E_Child_Three](
		"Parent_Child",
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "three"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Child_Three, bool) {
			ret := gs.(*oc.Parent_Child).Three
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/three"
//	Path from root:       "/parent/child/config/three"
func (n *Parent_Child_ThreePath) Config() ygnmi.ConfigQuery[oc.E_Child_Three] {
	return ygnmi.NewLeafConfigQuery[oc.E_Child_Three](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "three"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Child_Three, bool) {
			ret := gs.(*oc.Parent_Child).Three
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/three"
//	Path from root:       "/parent/child/config/three"
func (n *Parent_Child_ThreePathAny) Config() ygnmi.WildcardQuery[oc.E_Child_Three] {
	return ygnmi.NewLeafWildcardQuery[oc.E_Child_Three](
		"Parent_Child",
		false,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "three"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Child_Three, bool) {
			ret := gs.(*oc.Parent_Child).Three
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/two"
//	Path from root:       "/parent/child/state/two"
func (n *Parent_Child_TwoPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Parent_Child",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "two"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Parent_Child).Two
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/two"
//	Path from root:       "/parent/child/state/two"
func (n *Parent_Child_TwoPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Parent_Child",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "two"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Parent_Child).Two
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Parent_Child) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Parent_Child_FourPath represents the /openconfig-simple/parent/child/state/four YANG schema element.
type Parent_Child_FourPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_FourPathAny represents the wildcard version of the /openconfig-simple/parent/child/state/four YANG schema element.
type Parent_Child_FourPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_OnePath represents the /openconfig-simple/parent/child/state/one YANG schema element.
type Parent_Child_OnePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_OnePathAny represents the wildcard version of the /openconfig-simple/parent/child/state/one YANG schema element.
type Parent_Child_OnePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_SixPath represents the /openconfig-simple/parent/child/state/six YANG schema element.
type Parent_Child_SixPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_SixPathAny represents the wildcard version of the /openconfig-simple/parent/child/state/six YANG schema element.
type Parent_Child_SixPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_ThreePath represents the /openconfig-simple/parent/child/state/three YANG schema element.
type Parent_Child_ThreePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_ThreePathAny represents the wildcard version of the /openconfig-simple/parent/child/state/three YANG schema element.
type Parent_Child_ThreePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_TwoPath represents the /openconfig-simple/parent/child/state/two YANG schema element.
type Parent_Child_TwoPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_Child_TwoPathAny represents the wildcard version of the /openconfig-simple/parent/child/state/two YANG schema element.
type Parent_Child_TwoPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Parent_ChildPath represents the /openconfig-simple/parent/child YANG schema element.
type Parent_ChildPath struct {
	*ygnmi.NodePath
}

// Parent_ChildPathAny represents the wildcard version of the /openconfig-simple/parent/child YANG schema element.
type Parent_ChildPathAny struct {
	*ygnmi.NodePath
}

// Five (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/five"
//	Path from root:       "/parent/child/*/five"
func (n *Parent_ChildPath) Five() *Parent_Child_FivePath {
	return &Parent_Child_FivePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "five"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Five (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/five"
//	Path from root:       "/parent/child/*/five"
func (n *Parent_ChildPathAny) Five() *Parent_Child_FivePathAny {
	return &Parent_Child_FivePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "five"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Four (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/four"
//	Path from root:       "/parent/child/*/four"
func (n *Parent_ChildPath) Four() *Parent_Child_FourPath {
	return &Parent_Child_FourPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "four"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Four (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/four"
//	Path from root:       "/parent/child/*/four"
func (n *Parent_ChildPathAny) Four() *Parent_Child_FourPathAny {
	return &Parent_Child_FourPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "four"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// One (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/one"
//	Path from root:       "/parent/child/*/one"
func (n *Parent_ChildPath) One() *Parent_Child_OnePath {
	return &Parent_Child_OnePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "one"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// One (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/one"
//	Path from root:       "/parent/child/*/one"
func (n *Parent_ChildPathAny) One() *Parent_Child_OnePathAny {
	return &Parent_Child_OnePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "one"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Six (leaf-list):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/six"
//	Path from root:       "/parent/child/*/six"
func (n *Parent_ChildPath) Six() *Parent_Child_SixPath {
	return &Parent_Child_SixPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "six"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Six (leaf-list):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/six"
//	Path from root:       "/parent/child/*/six"
func (n *Parent_ChildPathAny) Six() *Parent_Child_SixPathAny {
	return &Parent_Child_SixPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "six"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Three (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/three"
//	Path from root:       "/parent/child/*/three"
func (n *Parent_ChildPath) Three() *Parent_Child_ThreePath {
	return &Parent_Child_ThreePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "three"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Three (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/three"
//	Path from root:       "/parent/child/*/three"
func (n *Parent_ChildPathAny) Three() *Parent_Child_ThreePathAny {
	return &Parent_Child_ThreePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "three"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Two (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/two"
//	Path from root:       "/parent/child/state/two"
func (n *Parent_ChildPath) Two() *Parent_Child_TwoPath {
	return &Parent_Child_TwoPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "two"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Two (leaf):
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/two"
//	Path from root:       "/parent/child/state/two"
func (n *Parent_ChildPathAny) Two() *Parent_Child_TwoPathAny {
	return &Parent_Child_TwoPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "two"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// RemoteContainer_ALeafPath represents the /openconfig-simple/remote-container/state/a-leaf YANG schema element.
type RemoteContainer_ALeafPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// RemoteContainer_ALeafPathAny represents the wildcard version of the /openconfig-simple/remote-container/state/a-leaf YANG schema element.
type RemoteContainer_ALeafPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
func (n *RemoteContainerPath) State() ygnmi.SingletonQuery[*oc.RemoteContainer] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.RemoteContainer](
		"RemoteContainer",
		true,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *RemoteContainerPathAny) State() ygnmi.WildcardQuery[*oc.RemoteContainer] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.RemoteContainer](
		"RemoteContainer",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *RemoteContainerPath) Config() ygnmi.ConfigQuery[*oc.RemoteContainer] {
	return ygnmi.NewNonLeafConfigQuery[*oc.RemoteContainer](
		"RemoteContainer",
		false,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *RemoteContainerPathAny) Config() ygnmi.WildcardQuery[*oc.RemoteContainer] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.RemoteContainer](
		"RemoteContainer",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/a-leaf"
//	Path from root:       "/remote-container/state/a-leaf"
func (n *RemoteContainer_ALeafPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"RemoteContainer",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "a-leaf"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.RemoteContainer).ALeaf
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.RemoteContainer) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "state/a-leaf"
//	Path from root:       "/remote-container/state/a-leaf"
func (n *RemoteContainer_ALeafPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"RemoteContainer",
		true,
		true,
		ygnmi.NewNodePath(
			[]string{"state", "a-leaf"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.RemoteContainer).ALeaf
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.RemoteContainer) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/a-leaf"
//	Path from root:       "/remote-container/config/a-leaf"
func (n *RemoteContainer_ALeafPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewLeafConfigQuery[string](
		"RemoteContainer",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"config", "a-leaf"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.RemoteContainer).ALeaf
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.RemoteContainer) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "config/a-leaf"
//	Path from root:       "/remote-container/config/a-leaf"
func (n *RemoteContainer_ALeafPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"RemoteContainer",
		false,
		true,
		ygnmi.NewNodePath(
			[]string{"config", "a-leaf"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.RemoteContainer).ALeaf
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.RemoteContainer) },
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// RemoteContainerPath represents the /openconfig-simple/remote-container YANG schema element.
type RemoteContainerPath struct {
	*ygnmi.NodePath
}

// RemoteContainerPathAny represents the wildcard version of the /openconfig-simple/remote-container YANG schema element.
type RemoteContainerPathAny struct {
	*ygnmi.NodePath
}

// ALeaf (leaf):
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/a-leaf"
//	Path from root:       "/remote-container/*/a-leaf"
func (n *RemoteContainerPath) ALeaf() *RemoteContainer_ALeafPath {
	return &RemoteContainer_ALeafPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "a-leaf"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// ALeaf (leaf):
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "*/a-leaf"
//	Path from root:       "/remote-container/*/a-leaf"
func (n *RemoteContainerPathAny) ALeaf() *RemoteContainer_ALeafPathAny {
	return &RemoteContainer_ALeafPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "a-leaf"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}
