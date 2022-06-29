package core

import "sync"

type Database struct {
	messages []Message
	m        sync.Mutex
}

// NewDatabase : constructor
func NewDatabase() *Database {
	db := Database{}
	db.Init()
	return &db
}

// Init : Database initialization
func (db *Database) Init() {
	db.messages = make([]Message, 0)
}

// Add : add Message in daatabase
func (db *Database) Add(msg Message) {
	db.m.Lock()
	defer db.m.Unlock()

	db.messages = append(db.messages, msg)
}

// Get : list database content
func (db *Database) Get(filterBy string) []Message {
	db.m.Lock()
	defer db.m.Unlock()

	results := make([]Message, 0)

	for _, msg := range db.messages {
		//apply filter
		if filterBy == "" || (msg.Tag != nil && filterBy == *msg.Tag) {
			results = append(results, msg)
		}
	}
	return results
}

// Len : number of database elements
func (db *Database) Len() int {
	db.m.Lock()
	defer db.m.Unlock()

	return len(db.messages)
}
