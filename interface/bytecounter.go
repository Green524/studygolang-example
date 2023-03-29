package count

import (
	"io"
	"time"
)

type ByteCounter int

func (r *ByteCounter) Write(p []byte) (int, error) {
	*r += ByteCounter(len(p))
	return len(p), nil
}

//func (r *ByteCounter) Read(b []byte) (n int, err error) {
//	return 0, nil
//}

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}
