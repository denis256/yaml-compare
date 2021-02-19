package main

import (
	"gopkg.in/yaml.v2"

	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.Printf("Loading files")

	values, err := ioutil.ReadFile("values.yaml")
	if err != nil {
		panic(err)
	}

	customValues, err := ioutil.ReadFile("custom-values.yaml")
	if err != nil {
		panic(err)
	}

	var valuesMap map[interface{}]interface{}
	var customValuesMap map[interface{}]interface{}

	yaml.Unmarshal(values, &valuesMap)
	yaml.Unmarshal(customValues, &customValuesMap)

	log.Printf("Loaded data:")

	log.Printf("%+v\n", valuesMap)
	log.Printf("%+v\n", customValuesMap)

	var customFlatMap = make(map[string]string)
	var valuesFlatMap = make(map[string]string)

	flatMap(&valuesFlatMap, "", &valuesMap)
	flatMap(&customFlatMap, "", &customValuesMap)

	fmt.Printf("***Values FlatMap:\n")
	for k, v := range valuesFlatMap {
		fmt.Printf("%+v => %+v\n", k, v)
	}

	fmt.Printf("***Custom FlatMap:\n")
	for k, v := range customFlatMap {
		fmt.Printf("%+v => %+v\n", k, v)
	}

	fmt.Printf("***Keys from custom yaml which not overriding values.yaml:\n")
	for k, _ := range customFlatMap {
		_, ok := valuesFlatMap[k]
		if !ok {
			fmt.Printf("%+v\n", k)
		}
	}

}
func flatMap(storage *map[string]string, prefix string, m *map[interface{}]interface{}) {
	for k, v := range *m {
		_, ok := v.(string)
		if ok {
			(*storage)[fmt.Sprintf("%+v.%+v", prefix, k)] = fmt.Sprintf("%+v", v)
		} else {
			var x = v.(map[interface{}]interface{})
			flatMap(storage, fmt.Sprintf("%+v.%+v", prefix, k), &x)
		}

	}
}
