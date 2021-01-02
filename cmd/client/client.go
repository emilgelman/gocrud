package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/emilg02/gocrud/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

var client proto.ArticleServiceClient

func main() {

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = proto.NewArticleServiceClient(conn)

	//Sub commands

	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)

	getIdPtr := getCommand.String("id", "", "article id")
	createIdPtr := createCommand.String("id", "", "article id")
	createTitlePtr := createCommand.String("title", "", "article title")
	createContentPtr := createCommand.String("content", "", "article content")
	deleteIdPtr := deleteCommand.String("id","","article id")

	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		getCommand.Parse(os.Args[2:])
		if *getIdPtr == "" {
			getArticles()
		} else {
			getArticle(*getIdPtr)
		}
	case "delete":
		deleteCommand.Parse(os.Args[2:])
		if *deleteIdPtr == "" {
			log.Fatalf("please specify article id")
		} else {
			deleteArticle(*deleteIdPtr)
		}
	case "create":
		createCommand.Parse(os.Args[2:])
		createArticle(*createIdPtr, *createTitlePtr, *createContentPtr)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func createArticle(id, title, content string) {
	res, err:= client.AddArticle(context.Background(), &proto.AddArticleRequest{
		Id: id,
		Title: title,
		Content: content,
	})
	if err != nil{
		log.Fatalf(err.Error())
	}
	log.Printf("%v", res)
}

func getArticles() {
	articles, err := client.GetArticles(context.Background(), new(proto.GetArticlesRequest))
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, a := range articles.Articles {
		log.Printf("%v\n", a)
	}
}

func getArticle(id string) {
	article, err := client.GetArticle(context.Background(), &proto.GetArticleRequest{Id: id})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("%v", article)
}

func deleteArticle(id string) {
	response, err := client.DeleteArticle(context.Background(), &proto.DeleteArticleRequest{Id: id})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("%v", response)
}
