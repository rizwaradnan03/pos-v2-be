package enums

type ArticleType string

const (
	ArticleTypePENDING   ArticleType = "PENDING"
	ArticleTypePUBLISHED ArticleType = "PUBLISHED"
	ArticleTypeDELETED   ArticleType = "DELETED"
	ArticleTypeREJECTED  ArticleType = "REJECTED"
)
