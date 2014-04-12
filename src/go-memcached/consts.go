package gomemcached

type StorageResult string

const (
    Sep = " "
    CRLF = "\r\n"
    OK = "OK"
    UnknownErr = "ERROR"
    CliErr = "CLIENT_ERROR"
    SrvErr = "SERVER_ERROR"
    End = "END"
    Value = "VALUE"

    KeyStored = "STORED"
    KeyNotStored = "NOT_STORED"
    KeyExists = "EXISTS"
    KeyNotFound = "NOT_FOUND"
    KeyDeleted = "DELETED"
    KeyTouched = "TOUCHED"
)

type Op string
const (
    OpGet = "GET"
    OpGets = "GETS"
    OpSet = "SET"
    OpAdd = "ADD"
    OpReplace = "REPLACE"
    OpDel = "DEL"
    OpIncr = "INCR"
    OpDecr = "DECR"
    OpPrepend = "PREPEND"
    OpAppend = "APPEND"
    OpCAS = "CAS"
    OpTouch = "TOUCH"
)
