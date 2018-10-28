// Copyright 2018 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package stdlib

import (
	"github.com/gentee/gentee/core"
)

type embedInfo struct {
	Func    interface{}
	InTypes string
	OutType string
}

// InitArray appends stdlib array functions to the virtual machine
func InitArray(vm *core.VirtualMachine) {
	for _, item := range []embedInfo{
		{AssignAddºArrStr, `arr.str,str`, `arr.str`},     // arr += str
		{LenºArr, `arr*`, `int`},                         // the length of array
		{AssignAddºArrInt, `arr.int,int`, `arr.int`},     // arr += int
		{AssignAddºArrArr, `arr.arr*,arr*`, `arr.arr*`},  // arr.arr += arr
		{AssignAddºArrBool, `arr.bool,bool`, `arr.bool`}, // arr += bool
		{AssignAddºArrMap, `arr.map*,map*`, `arr.map*`},  // arr.map += map
		{AssignºArrArr, `arr*,arr*`, `arr*`},             // arr = arr
	} {
		vm.StdLib().NewEmbedExt(item.Func, item.InTypes, item.OutType)
	}
}

// LenºArr returns the length of the array
func LenºArr(arr *core.Array) int64 {
	return int64(len(arr.Data))
}

// AssignºArrArr copies one array to another one
func AssignºArrArr(ptr *interface{}, value *core.Array) *core.Array {
	core.CopyVar(ptr, value)
	return (*ptr).(*core.Array)
}

// AssignAddºArrArr appends one array to another one
func AssignAddºArrArr(ptr *interface{}, value *core.Array) *core.Array {
	(*ptr).(*core.Array).Data = append((*ptr).(*core.Array).Data, value)
	return (*ptr).(*core.Array)
}

// AssignAddºArrMap appends a map to array
func AssignAddºArrMap(ptr *interface{}, value *core.Map) *core.Array {
	(*ptr).(*core.Array).Data = append((*ptr).(*core.Array).Data, value)
	return (*ptr).(*core.Array)
}

// AssignAddºArrStr appends one string to array
func AssignAddºArrStr(ptr *interface{}, value string) *core.Array {
	(*ptr).(*core.Array).Data = append((*ptr).(*core.Array).Data, value)
	return (*ptr).(*core.Array)
}

// AssignAddºArrInt appends one integer to array
func AssignAddºArrInt(ptr *interface{}, value int64) *core.Array {
	(*ptr).(*core.Array).Data = append((*ptr).(*core.Array).Data, value)
	return (*ptr).(*core.Array)
}

// AssignAddºArrBool appends one boolean value to array
func AssignAddºArrBool(ptr *interface{}, value bool) *core.Array {
	(*ptr).(*core.Array).Data = append((*ptr).(*core.Array).Data, value)
	return (*ptr).(*core.Array)
}