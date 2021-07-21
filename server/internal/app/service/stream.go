package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/raspiantoro/realtime-web-apps/server/internal/pkg/stream"
	streamPB "github.com/raspiantoro/realtime-web-apps/server/pkg/stream/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//StreamService define struct for stream service
type StreamService struct {
	Context context.Context
	streamPB.UnimplementedStreamServer
}

//NewStreamService create new StreamService that implement streamPB.StreamServer
func NewStreamService(ctx context.Context) streamPB.StreamServer {
	return &StreamService{
		Context: ctx,
	}
}

//GetDataStream service for handling GetDataStream request
func (s *StreamService) GetDataStream(req *streamPB.StreamRequest, streamResponse streamPB.Stream_GetDataStreamServer) error {
	ch := make(chan stream.Message, 1)
	key := time.Now().UnixNano()
	stream.AppendChan(key, ch)

	for {
		select {
		case <-s.Context.Done():
			delete(stream.GetChan(), key)
			close(ch)
			return nil
		case <-streamResponse.Context().Done():
			log.Printf("request %d has been canceled", key)
			delete(stream.GetChan(), key)
			close(ch)
			return nil
		case val := <-ch:
			resp := &streamPB.StreamResponse{
				Count: uint64(val.Count),
			}

			streamResponse.Send(resp)
		}
	}
}

func (s *StreamService) BidiStream(stream streamPB.Stream_BidiStreamServer) (err error) {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		fmt.Println("Receive request with ReqID: ", req.ReqID)

		resp := &streamPB.BiResponse{
			ReqID: req.ReqID,
			Reply: "Pong",
		}

		err = stream.Send(resp)
		if err != nil {
			log.Println("Error while sending response: ", err.Error())
			return status.Error(codes.Internal, "oops...something bad happen")
		}
	}

}
