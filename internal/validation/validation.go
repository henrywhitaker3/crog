package validation

import (
	"fmt"
	"strconv"

	"github.com/fatih/structs"
	"github.com/oleiade/reflections"
)

func Validate(item any) error {
	fields := structs.Fields(item)

	for _, field := range fields {
		if isRequired(field) {
			if !isSet(field) {
				return fmt.Errorf("field %s is required", field.Name())
			}
			continue
		}

		if hasDefault(field) && !isSet(field) {
			switch field.Kind().String() {
			case "int":
				val, err := strconv.Atoi(getDefault(field))
				if err != nil {
					return err
				}
				if err := reflections.SetField(item, field.Name(), val); err != nil {
					return err
				}
				fmt.Printf("setting %s to %d\n", field.Name(), val)
			case "bool":
				val, err := strconv.ParseBool(getDefault(field))
				if err != nil {
					return err
				}
				if err := reflections.SetField(item, field.Name(), val); err != nil {
					return err
				}
			default:
				if err := reflections.SetField(item, field.Name(), getDefault(field)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func isRequired(field *structs.Field) bool {
	return field.Tag("required") == "true"
}

func isSet(field *structs.Field) bool {
	return !field.IsZero()
}

func hasDefault(field *structs.Field) bool {
	return field.Tag("default") != ""
}

func getDefault(field *structs.Field) string {
	return field.Tag("default")
}
