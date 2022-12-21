package dtoArticle

type ArticleRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"user_id" `
}

type ArticleResponse struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Created string `json:"created"`
}
