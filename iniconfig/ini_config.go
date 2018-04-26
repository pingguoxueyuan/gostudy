package iniconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func MarshalFile(filename string, data interface{}) (err error) {
	result, err := Marshal(data)
	if err != nil {
		return
	}

	return ioutil.WriteFile(filename, result, 0755)
}

func Marshal(data interface{}) (result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	var conf []string
	valueInfo := reflect.ValueOf(data)
	for i := 0; i < typeInfo.NumField(); i++ {
		sectionField := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)

		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}

		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = sectionField.Name
		}

		section := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, section)

		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}

			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			conf = append(conf, item)
		}
	}

	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	return UnMarshal(data, result)
}

func UnMarshal(data []byte, result interface{}) (err error) {

	lineArr := strings.Split(string(data), "\n")

	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}

	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	var lastFieldName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		//如果是注释，直接忽略
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			lastFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
				return
			}
			continue
		}

		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)
			return
		}
	}
	return
}

func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}

	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])

	if len(key) == 0 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}

	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFieldName)

	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFieldName)
		return
	}

	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}

	if len(keyFieldName) == 0 {
		return
	}

	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}

	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)

	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		intVal, errRet := strconv.ParseUint(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			return
		}

		fieldValue.SetFloat(floatVal)

	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
	}

	return
}

func parseSection(line string, typeInfo reflect.Type) (fieldName string, err error) {

	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error, invalid section:%s", line)
			return
		}

		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name
				break
			}
		}
	}

	return
}
