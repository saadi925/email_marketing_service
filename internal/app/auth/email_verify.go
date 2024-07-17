package auth

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/saadi925/email_marketing_api/internal/database"
	"github.com/saadi925/email_marketing_api/internal/email"
)

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(999999))
}

type emailTemplateData struct {
	VerificationCode string
}
type EmailVerify struct {
	Email       string
	Code        string
	Retry       int
	LastAttempt time.Time
}

type emailVerifyRepository interface {
	emailVerifyExists(email string) (bool, error)
	saveEmailVerify(emailVerify EmailVerify) error
	updateEmailVerify(emailVerify EmailVerify) error
}

type emailVerifyRepo struct {
	db *database.Queries
}

func newEmailVerifyRepository(db *database.Queries) emailVerifyRepository {
	return &emailVerifyRepo{db: db}
}

func (repo *emailVerifyRepo) emailVerifyExists(email string) (bool, error) {
	_, err := repo.db.GetEmailVerifyByEmail(context.Background(), email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *emailVerifyRepo) saveEmailVerify(emailVerify EmailVerify) error {
	return repo.db.CreateEmailVerify(context.Background(), database.CreateEmailVerifyParams{
		Email:       emailVerify.Email,
		Code:        emailVerify.Code,
		Retry:       int32(emailVerify.Retry),
		LastAttempt: emailVerify.LastAttempt,
	})
}

func (repo *emailVerifyRepo) updateEmailVerify(emailVerify EmailVerify) error {
	return repo.db.UpdateEmailVerify(context.Background(), database.UpdateEmailVerifyParams{
		Email:       emailVerify.Email,
		Code:        emailVerify.Code,
		Retry:       int32(emailVerify.Retry),
		LastAttempt: emailVerify.LastAttempt,
	})
}

func sendVerificationEmail(queries *database.Queries, toEmail string) error {
	emailRepo := newEmailVerifyRepository(queries)
	smtpConfig := email.NewSmtpConfig()
	verificationCode := generateVerificationCode()

	dbEmailVerify := database.EmailVerify{
		Email:       toEmail,
		Code:        verificationCode,
		Retry:       0,
		LastAttempt: time.Now(),
	}

	exists, err := emailRepo.emailVerifyExists(dbEmailVerify.Email)
	if err != nil {
		return err
	}
	emailVerify := EmailVerify{
		Email:       dbEmailVerify.Email,
		Code:        dbEmailVerify.Code,
		Retry:       int(dbEmailVerify.Retry),
		LastAttempt: dbEmailVerify.LastAttempt,
	}
	if exists {

		if err := emailRepo.updateEmailVerify(emailVerify); err != nil {
			return fmt.Errorf("error updating email verify record: %w", err)
		}
	} else {
		if err := emailRepo.saveEmailVerify(emailVerify); err != nil {
			return fmt.Errorf("error saving verification code: %w", err)
		}
	}

	err = smtpConfig.SendEmail(toEmail, "Email Verification", "../../templates/email_verficiation.html", emailTemplateData{
		VerificationCode: verificationCode,
	})
	if err != nil {
		return fmt.Errorf("error sending verification email: %w", err)
	}

	return nil
}
