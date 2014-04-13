package main

import (
    "./go-memcached"
    "log"
)

func main() {
    gomemcached.ListenAndServe(":8088", func(req *gomemcached.Req, resp gomemcached.Resp)(error) {
        if req.Op() == gomemcached.OpSet {
            log.Println(req.Key())
            log.Println(req.Value())
            resp.WriteResult(gomemcached.KeyStored)
        }
        return nil
    })
}
