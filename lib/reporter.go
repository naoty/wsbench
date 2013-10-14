package lib

import (
    "io"
    "bytes"
    "strconv"
)

type Reporter struct {
    writer io.Writer
    CompletedRequests, FailedRequests int
}

func NewReporter(writer io.Writer) *Reporter {
    return &Reporter{
        writer: writer,
        CompletedRequests: 0,
        FailedRequests: 0,
    }
}

func (self *Reporter) Write(p []byte) (n int, err error) {
    buf := new(bytes.Buffer)

    buf.Write([]byte("CompletedRequests: "))
    buf.Write([]byte(strconv.Itoa(self.CompletedRequests)))
    buf.Write([]byte("\n"))

    buf.Write([]byte("FailedRequests: "))
    buf.Write([]byte(strconv.Itoa(self.FailedRequests)))
    buf.Write([]byte("\n"))

    buf.Write("\n")
    buf.Write(p)

    n, err = self.writer.Write(buf.Bytes())
    return
}
