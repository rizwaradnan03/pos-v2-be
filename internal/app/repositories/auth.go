package repositories

import (
	"pos-v2-be/internal/app/dtos"
	repositoryinterface "pos-v2-be/internal/app/interfaces/repository_interface"
	"pos-v2-be/internal/app/responses"
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/pkg/auth"
	"pos-v2-be/internal/pkg/number"
	"pos-v2-be/internal/pkg/str"
	"pos-v2-be/internal/pkg/validate"
	"pos-v2-be/internal/schema"
	"time"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repositoryinterface.AuthRepositoryInterface {
	return &authRepository{db: db}
}

func (r *authRepository) SignIn(input dtos.AuthSignInDto) (*responses.AuthSignInResponse, error) {
	var user schema.User

	err := r.db.Where("email = ?", input.Email).Where("type = ?", enums.AccountTypeCREDENTIAL).First(&user).Error
	if err != nil {
		return nil, err
	}

	comparedPassword := str.ComparePassword(input.Password, *user.Password)
	if !comparedPassword {
		return nil, err
	}

	tokenPayload := &dtos.TokenDto{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	tokenizedPayload, err := auth.JwtEncode(*tokenPayload)
	if err != nil {
		return nil, err
	}

	if user.RefreshToken != nil {
		dec, err := auth.JwtDecode(*user.RefreshToken)
		if err != nil {
			return nil, err
		}

		// to update the refresh token
		if *dec.Exp < number.TimeToNumber(time.Now()) {
			changeTime := number.HourToUnix(int(168))

			refreshTokenPayload := &dtos.TokenDto{
				ID:    user.ID,
				Email: user.Email,
				Role:  user.Role,
				Exp:   &changeTime,
			}

			encodedRefreshToken, err := auth.JwtEncode(*refreshTokenPayload)
			if err != nil {
				return nil, err
			}

			err = r.db.Model(&schema.User{}).Where("id = ?", user.ID).Updates(&schema.User{
				RefreshToken: &encodedRefreshToken,
			}).Error

			if err != nil {
				return nil, err
			}
		}
	}

	returnedPayload := &responses.AuthSignInResponse{
		AccessToken: tokenizedPayload,
	}

	return returnedPayload, nil
}

func (r *authRepository) IsEmailExist(email string) (bool, error) {
	err := r.db.Where("email = ?", email).First(&schema.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *authRepository) SignUp(input dtos.AuthSignInUpDto) (*schema.User, error) {
	err := dtos.ValidateAuthSignUp(input)
	if err != nil {
		return nil, err
	}

	vd, err := r.IsEmailExist(input.Email)
	if err != nil {
		return nil, err
	}

	if vd {
		return nil, validate.NewError("Email telah digunakan!")
	}

	hashedPassword, err := str.HashPassword(input.Password)
	if err != nil {
		return nil, validate.NewError("Gagal melakukan hash password!")
	}

	createValue := &schema.User{
		Email:      input.Email,
		Password:   &hashedPassword,
		MemberType: enums.AccountMemberTypeTypeEXTERNAL,
	}

	err = r.db.Create(createValue).Error
	if vd != false {
		return nil, validate.NewError("Gagal Melakukan SignUp!")
	}

	return createValue, nil
}

func (r *authRepository) FindOneByEmail(email string) (*schema.User, error) {
	User := &schema.User{}

	err := r.db.Where("email = ?", email).First(User).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return User, nil
}
