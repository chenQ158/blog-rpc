package service

import "github.com/jinzhu/gorm"

func GeneralQuery(db *gorm.DB, nativeSql string, args ...interface{}) []interface{} {
	rows, _ := db.Raw(nativeSql, args...).Rows()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var resList []interface{}
	for rows.Next() {
		rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}
			}
		}
		resList = append(resList, record)
	}
	return resList
}
