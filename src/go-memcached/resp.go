package gomemcached
import (
    "io"
    "fmt"
)
type Resp struct {
    w io.Writer
}

func NewResp(conn io.Writer) Resp {
    return Resp{conn}
}

func (r Resp) WriteResult(result StorageResult) error {
    s := result + CRLF
    _, err := r.w.Write([]byte(s))
    return err
}

func (r Resp) WriteRetrievalResult(key string, flags uint16, value []byte) error {
    line := fmt.Sprintf("VALUE %s %d %d\r\n", key, flags, len(value))
    if _, err := r.w.Write([]byte(line)); err != nil {
        return err
    }

    r.w.Write(value)
    if _, err := r.w.Write([]byte("\r\n")); err != nil {
        return err
    }
    return nil
}


func (r Resp) BegineMultiGet() {
    return
}

func (r Resp) EndMultiGet() {
    r.w.Write([]byte("END\r\n"))
}
