package transport

import (
	"context"
	"fmt"
	"github.com/emilg02/gocrud/pkg/db"
	"github.com/emilg02/gocrud/pkg/domain"
	"github.com/emilg02/gocrud/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)


const port = 8081

type grpcServer struct {
	srv *grpc.Server
	db *db.Db
}

func (server grpcServer) AddArticle(ctx context.Context, request *proto.AddArticleRequest) (*proto.AddArticleResponse, error) {
	article:= &domain.Article{
		Id: request.Id,
		Title: request.Title,
		Content: request.Content,
	}
	(*server.db).Create(request.Id, *article)
	return &proto.AddArticleResponse{Response: fmt.Sprintf("Successfuly added article with id %s", request.Id)}, nil
}

func (server grpcServer) GetArticles(ctx context.Context, request *proto.GetArticlesRequest) (*proto.GetArticlesResponse, error) {
	response := new(proto.GetArticlesResponse)
	for _, b := range (*server.db).GetAll() {
	response.Articles = append(response.Articles,
			&proto.Article{
				Id:      b.Id,
				Title:   b.Title,
				Content: b.Content,
			})
	}
	return response, nil
}

func (server grpcServer) GetArticle(ctx context.Context, request *proto.GetArticleRequest) (*proto.GetArticleResponse, error) {
	res, err:= (*server.db).Get(request.Id)
	if err != nil {
		return nil, err
	}
	response := &proto.GetArticleResponse{
		Article: &proto.Article{
			Id: res.Id,
			Title: res.Title,
			Content: res.Content,
		},
	}
	return response, nil
}

func (server grpcServer) DeleteArticle(ctx context.Context, request *proto.DeleteArticleRequest) (*proto.DeleteArticleResponse, error) {
	if err:= (*server.db).Delete(request.Id); err != nil {
		return nil, err
	}
	return &proto.DeleteArticleResponse{Response: "success"}, nil
}



func NewGRPCServer(db *db.Db) *grpcServer {
	return &grpcServer{
		srv:  grpc.NewServer(grpc.EmptyServerOption{}),
		db: db,
	}
}

func (server grpcServer) Serve() {
	proto.RegisterArticleServiceServer(server.srv, server)
	listener, err := net.Listen("tcp", "localhost: " +strconv.Itoa(port))
	if err != nil {
		log.Fatalf("unable to listen on port 81: %v", err)
	}
	log.Printf("gRPC listening on port %v", port)
	if err:= server.srv.Serve(listener); err != nil {
		log.Fatalf("Unable to start gRPC server: %v", err)
	}
}
