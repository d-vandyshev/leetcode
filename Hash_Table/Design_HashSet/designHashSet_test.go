package main

import (
	"reflect"
	"strings"
	"testing"
)

type testcase struct {
	MethodNames  []string
	MethodParams [][]int
	Returns []string 
}

var tests = []testcase{
{	[]string{"MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"},
	[][]int{{}, {1}, {2}, {1}, {3}, {2}, {2}, {2}, {2}},
	[]string{"null", "null", "null", "true", "false", "null", "true", "null", "false"}},
}

func stringsIsEqual(a, b []string) bool {
  if len(a) != len(b) {
    return false
  }
  for i, v := range a {
    if v != b[i] {
      return false
    }
  }
  return true
}

func dynamicCallMethod(obj interface{}, fn string, args []interface{}) (res []reflect.Value){
	var inputs []reflect.Value
	for _, v := range args {
			inputs = append(inputs, reflect.ValueOf(v))
	}
	return reflect.ValueOf(obj).MethodByName(fn).Call(inputs)
}

func TestDesignHashSet(t *testing.T) {
	for _, test := range tests {
		hashSet := Constructor()
		output := []string{"null"}
		for i := 1; i < len(test.MethodNames); i++ {
			args := make([]interface{}, 0)
			for _, p := range test.MethodParams[i] {
				args = append(args, p)
			}
			res := dynamicCallMethod(&hashSet, strings.Title(test.MethodNames[i]), args)
			if len(res) == 0 {
				output = append(output, "null")
				continue
			}
			if res[0].Bool() == true {
				output = append(output, "true")
			} else {
				output = append(output, "false")
			}
		}
		if !stringsIsEqual(output, test.Returns) {
			t.Error("For methods:", test.MethodNames, "\nwith params:", test.MethodParams, "\nexpected:", test.Returns, "\n     got:", output)
		}
	}
}
