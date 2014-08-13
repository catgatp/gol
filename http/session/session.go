/*
	Cookie based session engin implementation
	(c) 2014 Cergoo
	under terms of ISC license
*/
package session

import (
	"github.com/Cergoo/gol/err"
	"github.com/Cergoo/gol/http/cookie"
	"github.com/Cergoo/gol/http/genid"
	"math"
	"net/http"
	"strings"
)

type (
	TSession struct {
		ipProtect bool            // session ip protect
		gen       genid.HTTPGenID // id generator
		Stor      TStor           // store interface implementation
	}

	tdata struct {
		ip   string
		data interface{}
	}

	// Session store interface
	TStor interface {
		Get(string) interface{}
		Set(id string, data interface{})
		Del(string)
	}
)

const (
	sid = "sid"
)

// constructor
func NewSessionEngin(lenID uint8, ipProtect bool, stor TStor) *TSession {
	return &TSession{
		ipProtect: ipProtect,
		gen:       genid.NewHTTPGen(lenID),
		Stor:      stor,
	}
}

// Create new session
func (t *TSession) New(w http.ResponseWriter, r *http.Request, data interface{}) (id string) {
	id = t.gen.NewID()
	var sessionData *tdata
	if t.ipProtect || data != nil {
		sessionData = &tdata{data: data}
		if t.ipProtect {
			sessionData.ip = strings.SplitN(r.RemoteAddr, ":", 2)[0]
		}
	}
	t.Stor.Set(id, sessionData)
	cookie.SetCookie(w, sid, id, &cookie.Options{Path: "/", MaxAge: math.MaxInt32, HttpOnly: true})
	return
}

// Delete session
func (t *TSession) Del(w http.ResponseWriter, r *http.Request) {
	vcoockie, e := r.Cookie(sid)
	err.Panic(e)
	t.Stor.Del(vcoockie.Value)
	cookie.DelCookie(w, sid)
}

// Get session return sid and value
func (t *TSession) Get(w http.ResponseWriter, r *http.Request) (id []byte, val interface{}) {
	vcoockie, e := r.Cookie(sid)
	if e != nil {
		return
	}
	sessionData, b := t.Stor.Get(vcoockie.Value).(*tdata)
	// if session deleted or not chek ipProtect then del cookie
	if !b || sessionData == nil || (t.ipProtect && sessionData.ip != strings.SplitN(r.RemoteAddr, ":", 2)[0]) {
		cookie.DelCookie(w, sid)
		return
	}
	return []byte(vcoockie.Value), sessionData.data
}
