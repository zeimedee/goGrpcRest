package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		log.Println("logging....")
		res, err := handler(ctx, req)

		return res, err
	}

}
