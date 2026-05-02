package core_test

import (
	"fmt"
	"github.com/oceakun/myredis/core"
	"testing"
)

func TestSimpleString(t *testing.T) {
	cases := map[string]string{
		"+OK\r\n": "OK",
	}

	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestError(t *testing.T) {
	cases := map[string]string{
		"-ERROR ERROR\r\n": "ERROR ERROR",
	}

	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestInt64(t *testing.T) {
	cases := map[string]int64{
		":2\r\n": 2,
	}

	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestBulkString(t *testing.T) {
	cases := map[string]string{
		"$6\r\nYellow\r\n": "Yellow",
		"$0\r\n\r\n":       "",
	}

	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if v != value {
			t.Fail()
		}
	}
}

func TestArrayDecode(t *testing.T) {
	cases := map[string][]interface{}{
		"*0\r\n":                               {},
		"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n": {"hello", "world"},
		"*2\r\n:4\r\n:5\r\n":                   {int64(4), int64(5)},
		"*3\r\n:4\r\n:5\r\n$5\r\nworld\r\n":    {int64(4), int64(5), "world"},
		"*2\r\n*3\r\n:4\r\n:5\r\n:6\r\n*2\r\n$5\r\nworld\r\n$8\r\ndevourer\r\n": {
			[]interface{}{int64(4), int64(5), int64(6)},
			[]interface{}{"world", "devourer"},
		},
	}

	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		array := value.([]interface{})
		if len(array) != len(v) {
			t.Fail()
		}
		for i := range array {
			if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]) {
				t.Fail()
			}
		}
	}
}
