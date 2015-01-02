package spec

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

func ShouldSerializeJSON(actual interface{}, expected ...interface{}) string {
	ser, err := json.Marshal(actual)
	if err != nil {
		return err.Error()
	}
	exp := expected[0].(string)
	return ShouldEqual(string(ser), exp)
}

func ShouldParseJSON(actual interface{}, expected ...interface{}) string {
	exp := expected[0]
	tpe := reflect.TypeOf(exp)
	if tpe.Kind() == reflect.Ptr {
		tpe = tpe.Elem()
	}
	parsed := reflect.New(tpe)
	err := json.Unmarshal([]byte(actual.(string)), parsed.Interface())
	if err != nil {
		return err.Error()
	}
	return ShouldResemble(parsed.Interface(), exp)
}

func ShouldParseYAML(actual interface{}, expected ...interface{}) string {
	exp := expected[0]
	tpe := reflect.TypeOf(exp)
	if tpe.Kind() == reflect.Ptr {
		tpe = tpe.Elem()
	}
	parsed := reflect.New(tpe)
	err := yaml.Unmarshal([]byte(actual.(string)), parsed.Interface())
	if err != nil {
		return err.Error()
	}
	return ShouldResemble(parsed.Interface(), exp)
}

func ShouldSerializeYAML(actual interface{}, expected ...interface{}) string {
	ser, err := yaml.Marshal(actual)
	if err != nil {
		return err.Error()
	}
	exp := expected[0].(string)
	return ShouldEqual(string(ser), exp)
}

func TestSerialization(t *testing.T) {
	Convey("Swagger should serialize", t, func() {

		Convey("a string or array property", func() {
			Convey("when string", func() {
				obj := StringOrArray{Single: "hello"}

				Convey("for json returns quoted string", func() {
					So(obj, ShouldSerializeJSON, "\"hello\"")
				})
			})

			Convey("when slice", func() {
				obj := StringOrArray{Multi: []string{"hello", "world", "and", "stuff"}}
				Convey("for json returns an array of strings", func() {
					So(obj, ShouldSerializeJSON, "[\"hello\",\"world\",\"and\",\"stuff\"]")
				})
			})

			Convey("when empty", func() {
				obj := StringOrArray{}
				Convey("for json returns an empty array", func() {
					So(obj, ShouldSerializeJSON, "null")
				})
			})
		})

		Convey("a schema or array property", func() {
			Convey("when string", func() {
				obj := SchemaOrArray{Single: &Schema{schemaProps: schemaProps{Type: &StringOrArray{Single: "string"}}}}

				Convey("for json returns quoted string", func() {
					So(obj, ShouldSerializeJSON, "{\"type\":\"string\"}")
				})
			})

			Convey("when slice", func() {
				obj := SchemaOrArray{
					Multi: []Schema{
						Schema{schemaProps: schemaProps{Type: &StringOrArray{Single: "string"}}},
						Schema{schemaProps: schemaProps{Type: &StringOrArray{Single: "string"}}},
					},
				}
				Convey("for json returns an array of strings", func() {
					So(obj, ShouldSerializeJSON, "[{\"type\":\"string\"},{\"type\":\"string\"}]")
				})
			})

			Convey("when empty", func() {
				obj := SchemaOrArray{}
				Convey("for json returns an empty array", func() {
					So(obj, ShouldSerializeJSON, "null")
				})
			})
		})
	})

	Convey("Swagger should deserialize", t, func() {

		Convey("a string or array property", func() {
			Convey("when string", func() {
				obj := StringOrArray{Single: "hello"}

				Convey("for json returns quoted string", func() {
					So("\"hello\"", ShouldParseJSON, &obj)
				})
			})

			Convey("when slice", func() {
				obj := StringOrArray{Multi: []string{"hello", "world", "and", "stuff"}}
				Convey("for json returns an array of strings", func() {
					So("[\"hello\",\"world\",\"and\",\"stuff\"]", ShouldParseJSON, &obj)
				})
				Convey("for json returns an array of strings with nil", func() {
					So("[\"hello\",\"world\",null,\"stuff\"]", ShouldParseJSON, &StringOrArray{Multi: []string{"hello", "world", "", "stuff"}})
				})
			})

			Convey("when empty", func() {
				obj := StringOrArray{}
				Convey("for json returns an empty array", func() {
					So("null", ShouldParseJSON, &obj)
				})
			})
		})

		Convey("a schema or array property", func() {
			Convey("when string", func() {
				obj := SchemaOrArray{Single: &Schema{schemaProps: schemaProps{Type: &StringOrArray{Single: "string"}}}}

				Convey("for json returns quoted string", func() {
					So("{\"type\":\"string\"}", ShouldParseJSON, &obj)
				})
			})

			Convey("when slice", func() {
				obj := SchemaOrArray{
					Multi: []Schema{
						Schema{schemaProps: schemaProps{Type: &StringOrArray{Single: "string"}}},
						Schema{schemaProps: schemaProps{Type: &StringOrArray{Single: "string"}}},
					},
				}
				Convey("for json returns an array of strings", func() {
					So("[{\"type\":\"string\"},{\"type\":\"string\"}]", ShouldParseJSON, &obj)
				})
			})

			Convey("when empty", func() {
				Convey("for json returns an empty array", func() {
					So("null", ShouldParseJSON, &SchemaOrArray{})
				})
			})
		})
	})
}