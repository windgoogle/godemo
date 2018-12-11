package main

/**
@autor huangfeng
*/

type Reader interface {
	Read(p []byte) (n int, err error)
}
