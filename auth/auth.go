package auth

import (
	"github.com/gonyyi/gointf"
	"github.com/gonyyi/mutt"
)

var (
	bID = []byte("auth")
)

func New(db gointf.Storer, encryptionKey []byte) (*Auth, error) {
	encryptionKey = append(encryptionKey, "gonisalwaysawesomeforeverandever"...)
	a := &Auth{
		db:            db,
		encryptionKey: encryptionKey[0:32],
	}

	// Make sure buckets are available
	if err := a.db.NewBucket(bID); err != nil {
		return nil, err
	}

	return a, nil
}

type Auth struct {
	db            gointf.Storer
	encryptionKey []byte
}

func (a *Auth) User(id string) (mutt.User, error) {
	var out mutt.User
	b, err := a.db.Get(bID, []byte(id))
	if err != nil {
		return out, ERR_USERID_NOT_EXIST
	}
	err = out.Load(b)
	return out, err
}

// AddUser will add new user to thee system. IF id's already exists,
// this will return an error.
// TODO: need either mutex OR channel to be usd in DB transaction
// TODO: or simply support batch transaction to the interface
func (a *Auth) AddUser(id, passwd, name, email, addedBy string) error {
	if ok := mutt.CheckParamString(id, passwd, name, email, addedBy); !ok {
		return ERR_MISSING_REQUIRED_FIELDS
	}

	if _, err := a.db.Get(bID, []byte(id)); err != nil {
		// error means, the id not exist
		tmpId := mutt.User{ID: id}
		tmpId.Name.DisplayName = name
		tmpId.Email.Email = email
		tmpB, err := tmpId.Bytes()
		if err != nil {
			return err
		}
		return a.db.Put(bID, []byte(id), tmpB)
	}
	return ERR_USERID_EXIST
}
