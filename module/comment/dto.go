package comment

import "travel-server/model"

type QueryListReq struct {
	model.PageInfo
	UserId   int    `form:"user_id" json:"user_id"`
	IP       string `form:"ip" json:"ip"`
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

type ExamineReq struct {
	ID            int `form:"id" json:"id"`
	ExamineStatus int `form:"examine_status" json:"examine_status"`
}
