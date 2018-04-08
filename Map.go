package utils

func ExistsKey(mapValue map[interface{}]interface{}, key interface{}) bool {
	_, ok := mapValue[key]
	return ok
}
