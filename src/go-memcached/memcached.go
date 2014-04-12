package gomemcached
import (
    "net"
)

var Verbose bool

type ReqHandler func(request *Req, resp Resp) error

func parseRequest(conn net.Conn) (*Req, error) {
    return nil, nil
}

func connHandler(conn net.Conn, handler ReqHandler) {
    defer conn.Close()
    for {
        req := NewReqFromReader(conn)
        if handler != nil {
            resp := NewResp(conn)
            if req == nil {
                resp.WriteResult(CliErr)
                continue
            }
            if err := handler(req, resp); err != nil {
                resp.WriteResult(SrvErr)
            }
        } else {
            conn.Write([]byte(UnknownErr))
        }
    }
}

func ListenAndServe(addr string, handler ReqHandler) error {
    ln, err := net.Listen("tcp", addr)
    defer ln.Close()
    if err != nil {
        return err
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go connHandler(conn, handler)
    }

    return nil
}
