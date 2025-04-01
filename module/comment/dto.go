package comment

import "travel-server/model"

type _QueryListReq struct {
	model.PageInfo
	Creator       int    `form:"creator" json:"creator"`
	IP            string `form:"ip" json:"ip"`
	Content       string `form:"content" json:"content"`
	ArticleID     int    `form:"article_id" json:"article_id"`
	ExamineStatus int    `form:"examine_status" json:"examine_status"`
}

type _DeleteReq struct {
	ID int `form:"id" json:"id"`
}

type _ExamineReq struct {
	ID            int `form:"id" json:"id"`
	ExamineStatus int `form:"examine_status" json:"examine_status"`
}

type QueryListReq struct {
	model.PageInfo
	Creator int `form:"creator" json:"creator"`
}

type QueryListByArticleReq struct {
	ArticleID int `form:"article_id" json:"article_id"`
}

type CreateReq struct {
	Content   string `form:"content" json:"content" binding:"required"`
	ArticleID int    `form:"article_id" json:"article_id" binding:"required"`
	CommentID int    `form:"comment_id" json:"comment_id"`
}

type RemoveReq struct {
	ID int `form:"id" json:"id"`
}
