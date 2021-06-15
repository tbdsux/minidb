package minidb

import (
	"os"
)

// writes the file
// it just wraps the os.WriteFile
func write(file string, data []byte) error {
	return os.WriteFile(file, data, 0755)
}

// writeToDB write the db.store to the defined json db file.
func (db *MiniDB) writeToDB() {
	data := marshalStore(db.content)

	write(db.db, data)
}

// writeToDB writes the cols.store to the defined json db file.
func (cols *MiniCollections) writeToDB() {
	data := marshalStore(cols.content)

	write(cols.db, data)
}

// writeToDB writes the st.store to the defined json db file.
func (st *MiniStore) writeToDB() {
	data := marshalStore(st.content)

	write(st.db, data)
}
