package pkg

import "pos-v2-be/internal/schema"

func ModelsToMigrate() []interface{} {

	return []interface{}{
		&schema.User{},
		&schema.PersonalInformation{},
		&schema.Log{},
		&schema.Faq{},
		&schema.Category{},
		&schema.Article{},
		&schema.ArticleCover{},
		&schema.ArticleCategory{},
		&schema.ArticleView{},
		&schema.ArticleCategory{},
		&schema.ArticleViewLog{},
		&schema.ArticleLike{},
		&schema.CommissionBatch{},
		&schema.CommissionBatchProof{},
		&schema.CommissionBatchLike{},
		&schema.CommissionBatchView{},
		&schema.Donation{},
		&schema.DonationProof{},
		&schema.DonationApprovalProof{},
	}
}
