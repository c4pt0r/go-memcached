package gomemcached

import (
    "testing"
    "log"
)

func Test_Server(t *testing.T) {
    ListenAndServe(":8088", func(request *Req, resp Resp) (error){
        if request.Op() == OpSet {
            log.Println("opSet")
        } else {
            resp.WriteResult(UnknownErr)
        }
        return nil
    })
}
