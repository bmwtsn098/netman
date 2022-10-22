package fuzz

import "strconv"
import "github.com/ikilobyte/netman/server"
import "github.com/ikilobyte/netman/util"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }

    switch num {

    case 0:
        var test server.BaseConnect
        test.Write(bytes)
        return 0

    case 1:
        content := string(bytes)
        server.WithTls("mayhem", content)
        return 0

    case 2:
        var test util.DataPacker
        test.UnPack(bytes)
        return 0

    default:
        var test util.DataPacker
        num := uint32(num)
        test.Pack(num, bytes)
        return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}