package model

const (
	TABLE_NAME_SQLITE_SEQUENCE = "sqlite_sequence"
)

func GetNextSequence(tableName string) (int, error) {
	db := GetDB().Table(TABLE_NAME_SQLITE_SEQUENCE)

	var currentSeq int
	if result := db.Select("seq").Where("name = ?", tableName).Find(&currentSeq); result.Error != nil {
		return 0, result.Error
	}

	// update cache

	return currentSeq, nil
}
