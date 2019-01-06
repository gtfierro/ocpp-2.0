package main

type Function interface {
	Call(OCPPCallFrame) (OCPPCallResultFrame, OCPPCallErrorFrame)
}
