// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
"context"

"article/rpc/internal/logic"
"article/rpc/internal/svc"
"article/rpc/rpc"
)

type ArticleServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedArticleServer
}

func NewArticleServer(svcCtx *svc.ServiceContext) *ArticleServer {
	return &ArticleServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServer) Articles(ctx context.Context, in *rpc.ArtsRequest) (*rpc.ArtsResponse, error) {
	l := logic.NewArticlesLogic(ctx, s.svcCtx)
	return l.Articles(in)
}

func (s *ArticleServer) Article(ctx context.Context, in *rpc.ArtRequest) (*rpc.ArtResponse, error) {
	l := logic.NewArticleLogic(ctx, s.svcCtx)
	return l.Article(in)
}

func (s *ArticleServer) AddMyArticle(ctx context.Context, in *rpc.AddArtRequest) (*rpc.AddArtResponse, error) {
	l := logic.NewAddMyArticleLogic(ctx, s.svcCtx)
	return l.AddMyArticle(in)
}

func (s *ArticleServer) MyArticle(ctx context.Context, in *rpc.MyArtRequest) (*rpc.MyArtResponse, error) {
	l := logic.NewMyArticleLogic(ctx, s.svcCtx)
	return l.MyArticle(in)
}

func (s *ArticleServer) UpdateArticle(ctx context.Context, in *rpc.UpdateArtRequest) (*rpc.UpdateArtResponse, error) {
	l := logic.NewUpdateArticleLogic(ctx, s.svcCtx)
	return l.UpdateArticle(in)
}

func (s *ArticleServer) DelArticle(ctx context.Context, in *rpc.DelArtRequest) (*rpc.DelArtResponse, error) {
	l := logic.NewDelArticleLogic(ctx, s.svcCtx)
	return l.DelArticle(in)
}