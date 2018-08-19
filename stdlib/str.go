// Copyright 2018 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package stdlib

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gentee/gentee/core"
)

// InitStr appends stdlib int functions to the virtual machine
func InitStr(vm *core.VirtualMachine) {
	for _, item := range []interface{}{
		AddºStr,     // binary +
		EqualºStr,   // binary ==
		GreaterºStr, // binary >
		LenºStr,     // teh length of str
		LessºStr,    // binary <
		intºStr,     // int( str )
		boolºStr,    // bool( str )
		ExpStr,      // expression in string
	} {
		vm.StdLib().NewEmbed(item)
	}
}

// ExpStr adds two strings in string expression
func ExpStr(left, right string) string {
	return left + right
}

// AddºStr adds two integer value
func AddºStr(left, right string) string {
	return left + right
}

// EqualºStr returns true if left == right
func EqualºStr(left, right string) bool {
	return left == right
}

// GreaterºStr returns true if left > right
func GreaterºStr(left, right string) bool {
	return left > right
}

// LenºStr returns teh length of the string
func LenºStr(param string) int64 {
	return int64(len(param))
}

// LessºStr returns true if left < right
func LessºStr(left, right string) bool {
	return left < right
}

// intºStr converts strings value to int64
func intºStr(val string) (ret int64, err error) {
	ret, err = strconv.ParseInt(val, 0, 64)
	if err != nil {
		err = errors.New(core.ErrorText(core.ErrStrToInt))
	}
	return
}

// intºBool converts boolean value to int false -> 0, true -> 1
func boolºStr(val string) bool {
	return len(val) != 0 && val != `0` && strings.ToLower(val) != `false`
}
