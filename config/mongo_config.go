package mongo_config

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"sync"
)

var session *mgo.Session
var once sync.Once

func GetMongoSession() *mgo.Session {

	once.Do(func() {
		var err error
		session, err = mgo.Dial("mongodb://root:example@localhost")
		if err != nil {
			log.Fatal(err)
		}

		if err = session.Copy().DB("minishop").C("accounts").EnsureIndex(mgo.Index{
			Key:    []string{"email"},
			Unique: true,
		}); err != nil {
			log.Fatal(err)
		}
	})

	return session.Clone()
}

func CloseMongoSession(s *mgo.Session) {
	s.Close()
}
