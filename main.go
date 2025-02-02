package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"runtime"

	"github.com/lfun125/tesseract-go/ocr"
	"github.com/lfun125/tesseract-go/service"
	"google.golang.org/grpc"
)

func main() {
	logger := slog.Default()
	ocrService := service.NewOCRService()

	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			newUnaryRecoveryInterceptor(logger),
		),
	}
	s := grpc.NewServer(grpcOpts...)
	ocr.RegisterOCRServiceServer(s, ocrService)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}
	logger.Info("server started at :50051")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func newUnaryRecoveryInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		defer func() {
			if err := recover(); err != nil {
				stackTrace := captureStackTrace(20)
				logger.With("stackTrace", stackTrace).Error(fmt.Sprintf("Panic recovered: %v", err))
			}
		}()
		data, err := handler(ctx, req)
		if err != nil {
			logger.With("err", err).Error("unary handler error")
		}
		return data, err
	}
}

func captureStackTrace(depth int) []string {
	var stack []string
	for i := 0; i < depth; i++ {
		_, file, line, ok := runtime.Caller(i + 1) // +1 to skip current frame
		if !ok {
			break
		}
		stack = append(stack, fmt.Sprintf("%s:%d", file, line))
	}
	return stack
}
