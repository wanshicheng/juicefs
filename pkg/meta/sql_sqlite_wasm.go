package meta

func init() {
	// errBusy = sqlite3.ErrBusy
	// dupErrorCheckers = append(dupErrorCheckers, isSQLiteDuplicateEntryErr)
	Register("sqlite3", newSQLMeta)
}