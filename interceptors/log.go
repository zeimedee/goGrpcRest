package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()

		log.Println(info.FullMethod, time.Since(start), req)
		res, err := handler(ctx, req)

		return res, err
	}

}
