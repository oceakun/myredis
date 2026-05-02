package core

//contains everything around RESP implementation - encoding and decoding of data transferred over the net
//throughout this file, int datatype is mentioned between the interface{} and error types in the return type tuple, that represents the lenght of the serialized input until which it has been traversed by DecodeOne() func

import "errors"

func readSimpleString(data []byte) (interface{}, int, error) {
	//traverse the data and look for these two in the same order : '+' and '\r\n', return everything in between as the actual value of the string
	//regex -> int = length of the value + 4

	//start traversing from the 2nd character of data
	pos := 1

	//traverse until you find '\r'
	for ; data[pos] != '\r'; pos++ {

	}

	//pos+2 represents the position of the next character after '\n'
	return string(data[1:pos]), pos + 2, nil
}

func readError(data []byte) (interface{}, int, error) {
	return readSimpleString(data)
}

func readInt64(data []byte) (int64, int, error) {
	pos := 1
	var value int64 = 0

	for ; data[pos] != '\r'; pos++ {
		value = value*10 + int64(data[pos]-'0')
	}

	return value, pos + 2, nil
}

func readLength(data []byte) (int, int) {
	pos, len := 0, 0

	for ; data[pos] != '\r'; pos++ {

		len = len*10 + int(data[pos]-'0')
	}

	return len, pos + 2
}

// func readLength(data []byte) (int, int){
// 	pos, len:= 0, 0

// 	for pos = range data{
// 		b := data[pos]
// 		if !(b>='0' && b<='9')
// 		{
// 			return len, pos+2
// 		}

// 		len=len*10+ int(b-'0')
// 	}

// 	return 0, 0
// }

func bulkString(data []byte) (string, int, error) {
	pos := 1
	len, delta := readLength(data[pos:])
	pos += delta

	return string(data[pos : pos+len]), pos + len + 2, nil
}

func readArray(data []byte) (interface{}, int, error) {
	pos := 1

	count, delta := readLength(data[pos:])
	pos += delta

	var elems []interface{} = make([]interface{}, count)
	for i := range elems {
		elem, delta, err := DecodeOne(data[pos:])
		if err != nil {
			return nil, 0, err
		}
		elems[i] = elem
		pos += delta
	}
	return elems, pos, nil
}

func DecodeOne(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data present")
	}
	switch data[0] {
	case '+':
		return readSimpleString(data)
	case '-':
		return readError(data)
	case '*':
		return readArray(data)
	case ':':
		return readInt64(data)
	case '$':
		return bulkString(data)
	}

	return nil, 0, nil

}

func Decode(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	// DecodeOne function decodes the first RESP value out of the serialized data, no matter the total number of values in it originally
	value, _, err := DecodeOne(data)
	return value, err
}
