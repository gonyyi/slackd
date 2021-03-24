package slackd

// Logger is an interface slackd will take
type Logger interface {
	User(id string, ip string, category string, detail string)
	Store(key string, action string, detail string)
	Request(url string, id string, category string, detail string)
	System(category string, detail string)
}

// Storer is an interface for DB connection
type Storer interface {
	Open(dst string, id string, pwd string) error
	Close() error
	Create(key string, data []byte) error
	Read(key string) ([]byte, error)
	Update(key string, data []byte) error
	Delete(key string) error
}

// Requester is an interface for HTTP request
type Requester interface {
	Get(url string, id string, passwd string)
	Put(url string, id string, passwd string, data []byte)
	Post(url string, id string, passwd string, data []byte)
}
