package shared

import (
	"time"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/yosssi/boltstore/shared/protobuf"
)

// Session converts the byte slice to the session struct value.
func Session(data []byte) (protobuf.Session, error) {
	session := protobuf.Session{}
	err := proto.Unmarshal(data, &session)
	return session, err
}

// Expired checks if the session is expired.
func Expired(session protobuf.Session) bool {
	return *session.ExpiresAt > 0 && *session.ExpiresAt <= time.Now().Unix()
}