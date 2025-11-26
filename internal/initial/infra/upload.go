package infra

import (
	"log"
	"os"
)

type Upload struct {
	BasePath              string
	Category              string
	Banner                string
	ArticleCover          string
	PersonalInformation   string
	CommissionBatchProof  string
	DonationProof         string
	DonationApprovalProof string
}

func NewUpload() *Upload {
	basePath := "./uploads"
	category := basePath + "/category"
	banner := basePath + "/banner"
	articleCover := basePath + "/article-cover"
	personalInformation := basePath + "/personal-information"
	commissionBatchProof := basePath + "/commission-batch-proof"
	donationProof := basePath + "/donation-proof"
	donationApprovalProof := basePath + "/donation-approval-proof"

	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", basePath, err)
	}

	if err := os.MkdirAll(articleCover, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", articleCover, err)
	}

	if err := os.MkdirAll(category, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", category, err)
	}

	if err := os.MkdirAll(banner, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", banner, err)
	}

	if err := os.MkdirAll(personalInformation, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", personalInformation, err)
	}

	if err := os.MkdirAll(commissionBatchProof, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", commissionBatchProof, err)
	}

	if err := os.MkdirAll(donationProof, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", donationProof, err)
	}

	if err := os.MkdirAll(donationApprovalProof, os.ModePerm); err != nil {
		log.Fatalf("Gagal membuat folder %s: %v", donationApprovalProof, err)
	}

	return &Upload{
		BasePath:              basePath,
		PersonalInformation:   personalInformation,
		ArticleCover:          articleCover,
		CommissionBatchProof:  commissionBatchProof,
		DonationProof:         donationProof,
		DonationApprovalProof: donationApprovalProof,
		Banner:                banner,
	}
}
