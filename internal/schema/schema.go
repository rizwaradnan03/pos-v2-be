package schema

import (
	"pos-v2-be/internal/enums"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Faq struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Ask    string    `json:"ask"`
	Answer string    `json:"answer"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Banner struct {
	ID     uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	File   string       `json:"file"`
	Status enums.Banner `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID           uuid.UUID                   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email        string                      `gorm:"unique" json:"email"`
	Password     *string                     `gorm:"size:255" json:"-"`
	Role         enums.RoleType              `gorm:"type:varchar(20);default:'AUTHOR'" json:"role"`
	Type         enums.AccountType           `gorm:"type:varchar(20);default:'CREDENTIAL'" json:"type"`
	MemberType   enums.AccountMemberTypeType `gorm:"type:varchar(20);default:'EXTERNAL'" json:"account_type"`
	RefreshToken *string                     `json:"refresh_token"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Article             []*Article           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article,omitempty"`
	PersonalInformation *PersonalInformation `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"personal_information,omitempty"`
	CommissionBatch     []*CommissionBatch   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch,omitempty"`
}

type PersonalInformation struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FullName string    `json:"full_name"`
	Image    *string   `json:"image"`
	// UserID      uuid.UUID `json:"author_id"`
	UserID      uuid.UUID `gorm:"column:user_id" json:"author_id"`
	PhoneNumber string    `json:"phone_number"`
	BankName    string    `json:"bank_name"`
	BankAccount string    `json:"bank_account"`
	Web         *string   `json:"web"`
	Twitter     *string   `json:"twitter"`
	Instagram   *string   `json:"instagram"`
	Youtube     *string   `json:"youtube"`
	FaceBook    *string   `json:"facebook"`

	User *User `json:"user,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	File string    `json:"file"`
	Name string    `gorm:"unique" json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ArticleCategory []*ArticleCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article_category,omitempty"`
}

type Article struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID  uuid.UUID `json:"author_id"`
	Title   string    `json:"title"`
	Slug    string    `gorm:"unique" json:"slug"`
	Content string    `gorm:"text" json:"content"`

	Status enums.ArticleType `gorm:"type:varchar(20);default:'PENDING'" json:"status"`

	User *User `json:"user,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ArticleLike     []*ArticleLike     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article_like,omitempty"`
	ArticleCategory []*ArticleCategory `gorm:"foreignKey:ArticleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article_category,omitempty"`
	ArticleView     []*ArticleView     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article_view,omitempty"`
	ArticleCover    []*ArticleCover    `gorm:"foreignKey:ArticleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article_cover,omitempty"`
	CommissionBatch []*CommissionBatch `gorm:"foreignKey:ArticleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch,omitempty"`
	Donation        []*Donation        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"donation,omitempty"`
}

type Donation struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	From      string             `json:"from"`
	Message   string             `json:"message"`
	ArticleId uuid.UUID          `json:"article_id"`
	Status    enums.DonationType `gorm:"type:varchar(20);default:'PENDING'" json:"status"`

	ActualAmount   float64 `json:"actual_amount"`
	ReceivedAmount float64 `json:"received_amount"`

	Article *Article `json:"article,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DonationProof         []*DonationProof         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"donation_proof,omitempty"`
	DonationApprovalProof []*DonationApprovalProof `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"donation_approval_proof,omitempty"`
}

type DonationProof struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	File       string    `json:"file"`
	DonationId uuid.UUID `json:"donation_id"`

	Donation *Donation `json:"donation,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DonationApprovalProof struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	File       string    `json:"file"`
	DonationId uuid.UUID `json:"donation_id"`

	Donation *Donation `json:"donation,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleCover struct {
	ID        uuid.UUID               `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ArticleID uuid.UUID               `json:"article_id"`
	File      string                  `json:"file"`
	Order     enums.ArticleCoverOrder `gorm:"type:varchar(20);default:'SECONDARY'" json:"order"`
	Type      enums.ArticleCoverType  `gorm:"type:varchar(20);default:'IMAGE'" json:"type"`

	Article *Article `json:"article,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleCategory struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ArticleID  uuid.UUID `json:"article_id"`
	CategoryID uuid.UUID `json:"category_id"`

	Article  *Article  `gorm:"foreignKey:ArticleID;references:ID" json:"article,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleLike struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ArticleID uuid.UUID `json:"article_id"`
	Ip        string    `json:"ip"`

	Article *Article `json:"article,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CommissionBatchLike *CommissionBatchLike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch_like,omitempty"`
}

type ArticleView struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ArticleID uuid.UUID `json:"article_id"`
	Ip        string    `json:"ip"`
	LastSeen  time.Time `json:"last_seen"`

	Article *Article `json:"article,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ArticleViewLog []*ArticleViewLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"article_view_log,omitempty"`
}

type ArticleViewLog struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ArticleViewID uuid.UUID `json:"article_view_id"`

	ArticleView *ArticleView `json:"article_view,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CommissionBatchView *CommissionBatchView `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch_view,omitempty"`
}

type CommissionBatch struct {
	ID        uuid.UUID                 `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID                 `json:"author_id"`
	ArticleID uuid.UUID                 `json:"article_id"`
	Amount    float64                   `json:"amount"`
	Status    enums.CommissionBatchType `gorm:"type:varchar(20);default:'PENDING'" json:"status"`

	User    *User    `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Article *Article `gorm:"foreignKey:ArticleID;references:ID" json:"article,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CommissionBatchView  []*CommissionBatchView  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch_view,omitempty"`
	CommissionBatchLike  []*CommissionBatchLike  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch_like,omitempty"`
	CommissionBatchProof []*CommissionBatchProof `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch_proof,omitempty"`
}

type CommissionBatchProof struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	File              string    `json:"file"`
	CommissionBatchId uuid.UUID `json:"commission_batch_id"`

	CommissionBatch *CommissionBatch `json:"commission_batch,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommissionBatchView struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CommissionBatchID uuid.UUID `json:"commission_batch_id"`
	ArticleViewLogID  uuid.UUID `json:"article_view_log_id"`

	CommissionBatch *CommissionBatch `json:"commission_batch,omitempty"`
	ArticleViewLog  *ArticleViewLog  `json:"article_view_log,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommissionBatchLike struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CommissionBatchID uuid.UUID `json:"commission_batch_id"`
	ArticleLikeID     uuid.UUID `json:"article_like_id"`

	CommissionBatch *CommissionBatch `json:"commission_batch,omitempty"`
	ArticleLike     *ArticleLike     `json:"article_like,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
