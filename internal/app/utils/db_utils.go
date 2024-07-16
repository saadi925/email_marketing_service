package utils

import (
	"reflect"
)

// dbModelToModel converts a database model to an application model.
func DbModelToModel(dbModel interface{}, appModel interface{}) {
	dbValue := reflect.ValueOf(dbModel)
	appValue := reflect.ValueOf(appModel).Elem() // Dereference to set fields

	for i := 0; i < dbValue.NumField(); i++ {
		dbFieldName := dbValue.Type().Field(i).Name
		appField := appValue.FieldByName(dbFieldName)

		if appField.IsValid() {
			dbFieldValue := dbValue.Field(i)
			appField.Set(dbFieldValue)
		}
	}
}
