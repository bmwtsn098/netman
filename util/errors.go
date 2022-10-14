package util

import "errors"

var HeadBytesLengthFail = errors.New("head bytes fail")
var RouterNotFound = errors.New("router Not Found")
var BodyLenExceedLimit = errors.New("body length exceed limit")
var TLSHandshakeUnFinish = errors.New("tls handshake un finish")
var WebsocketOpcodeFail = errors.New("websocket opcode fail")
var WebsocketRsvFail = errors.New("websocket RSV must be 0")
