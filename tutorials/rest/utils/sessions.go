package utils

import(
  "time"
  "sync"
  "net/http"
  "../models"
  "github.com/satori/go.uuid"
)

const(
  cookieName = "go_session"
  cookieExpires = 24 * 2 * time.Hour //dos días
)

var Sessions = struct{
  m map[string] *models.User
  sync.RWMutex
}{m: make(map[string] *models.User)}

func SetSession(user *models.User, w http.ResponseWriter){
  Sessions.Lock()
  defer Sessions.Unlock()

  uuid := uuid.NewV4().String()
  Sessions.m[uuid] = user

  cookie := &http.Cookie{
    Name: cookieName,
    Value: uuid,
    Path: "/",
    Expires: time.Now().Add( cookieExpires ),
  }
  http.SetCookie(w,cookie)
}

func GetUser(r *http.Request) *models.User{
  Sessions.Lock()
  defer Sessions.Unlock()

  uuid := getValCookie(r)
  if user, ok := Sessions.m[uuid]; ok{
    return user
  }
  return &models.User{}
}

func DeleteSession(w http.ResponseWriter){
  cookie := &http.Cookie{
    Name: cookieName,
    Value: "",
    Path: "/",
    MaxAge: -1,
  }
  http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request) string{
  if cookie, err := r.Cookie(cookieName); err == nil{
    return cookie.Value //uuid
  }
  return ""
}

func IsAuthenticated(r *http.Request) bool{
  return getValCookie(r) != ""
}