package gomemcached
import (
    "io"
    "log"
    "bufio"
    "strings"
    "strconv"
)

type Req struct {
    op  Op
    keys [][]byte
    values [][]byte
    isMulti bool
    flags uint16
    exp   int64
    cas int64
    noreply bool
}

func (r *Req) Op() Op {
    return r.op
}

func (r *Req) Key() []byte {
    if r.isMulti == false && len(r.keys) == 1 {
        return r.keys[0]
    }
    return nil
}

func (r *Req) Keys() [][]byte {
    return r.keys
}

func (r *Req) Value() []byte {
    if r.isMulti == false && len(r.values) == 1 {
        return r.values[0]
    }
    return nil
}

func (r *Req) Values() [][]byte {
    return r.values
}

func (r *Req) Flags() uint16 {
    return r.flags;
}

func (r *Req) Exp() int64 {
    return r.exp
}

func (r *Req) NoReply() bool {
    return r.noreply
}


func NewReqFromReader(reader io.Reader) *Req {
    r := bufio.NewReader(reader)
    line, _, err := r.ReadLine()
    if err != nil {
        return nil
    }
    if parts := strings.Split(string(line), " "); len(parts) >0 {
        cmd := strings.ToUpper(parts[0])
        switch cmd {
            case OpGets, OpGet: {
                req := new(Req)
                if cmd == OpGet {
                    req.op = OpGet
                    req.isMulti = false
                } else {
                    req.op = OpGets
                    req.isMulti = true
                }
                for _, key := range parts[1:] {
                    req.keys = append(req.keys, []byte(key))
                }
                return req
            }

            case OpSet: {
                req := new(Req)
                req.keys = append(req.keys, []byte(parts[1]))
                // parse flags
                flags, err := strconv.ParseUint(parts[2], 10, 16)
                if err != nil {
                    log.Println(err)
                    return nil
                }
                req.flags = uint16(flags)
                // parse exp
                exp, err := strconv.ParseInt(parts[3], 10, 64)
                if err != nil {
                    log.Println(err)
                    return nil
                }
                req.exp = exp
                // parse bytes
                blen, err := strconv.ParseInt(parts[4], 10, 64)
                if err != nil {
                    log.Println(err)
                    return nil
                }
                // read value bytes
                buf := make([]byte, blen)
                cur := int64(0)
                for {
                    n , err := r.Read(buf[cur:])
                    if err != nil {
                        log.Println(err)
                        return nil
                    }
                    cur += int64(n)
                    if cur >= blen {
                        break
                    }
                }
                req.values = append(req.values, buf)
                // check \r\n valid
                buf = make([]byte,2)
                r.Read(buf)
                if string(buf) != CRLF {
                    log.Println("request not valid")
                    return nil
                }
                return req
            }
        }
    }
    return nil
}
