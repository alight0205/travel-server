package comment

import "travel-server/model"

type QueryListReq struct {
	model.PageInfo
	IP       string `form:"ip" json:"ip"`
	Account  string `form:"account" json:"account"`
	Email    string `form:"email" json:"email"`
	Content  string `form:"content" json:"content"`
	Province string `form:"province" json:"province"`
	City     string `form:"city" json:"city"`
	Examine  int    `form:"examine" json:"examine"`
}

type QueryByArticleReq struct {
	ArticleID int `form:"article_id" json:"article_id"`
}

type CreateReq struct {
	IP        string `form:"ip" json:"ip"`
	Nickname  string `form:"nickname" json:"nickname"`
	Avatar    string `form:"avatar" json:"avatar"`
	Content   string `form:"content" json:"content"`
	Province  string `form:"province" json:"province"`
	City      string `form:"city" json:"city"`
	ArticleID int    `form:"article_id" json:"article_id"`
	CommentID int    `form:"comment_id" json:"comment_id"`
}

type DeleteReq struct {
	ID int `form:"id" json:"id"`
}

type UpdateReq struct {
	ID        int    `form:"id" json:"id"`
	IP        string `form:"ip" json:"ip"`
	Nickname  string `form:"nickname" json:"nickname"`
	Avatar    string `form:"avatar" json:"avatar"`
	Account   string `form:"account" json:"account"`
	Email     string `form:"email" json:"email"`
	Content   string `form:"content" json:"content"`
	Province  string `form:"province" json:"province"`
	City      string `form:"city" json:"city"`
	Examine   int    `form:"examine" json:"examine"`
	Like      int    `form:"like" json:"like"`
	ArticleID int    `form:"article_id" json:"article_id"`
	CommentID int    `form:"comment_id" json:"comment_id"`
}
