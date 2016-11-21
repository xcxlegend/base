package base

import (
	"sort"
)

func Mapkeys(oMap *map[interface{}]interface{}) *[]interface{} {
    var keys = &[]interface{}
    for k,_ := range oMap{
        keys = append(keys, k)
    }
    return &keys
}

func Asort(oMap *map[interface{}]interface{}) *map[interface{}]interface{} {

}
