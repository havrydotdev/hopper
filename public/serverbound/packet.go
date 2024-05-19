package sbound

import (
	"io"
)

type Packet interface {
	io.ReaderFrom
}
