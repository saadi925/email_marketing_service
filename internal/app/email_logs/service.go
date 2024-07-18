package email_logs

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type EmailLogService interface {
	CreateEmailLog(ctx context.Context, emailID uuid.UUID, status, message string) (*EmailLog, error)
	GetEmailLogsByEmailID(ctx context.Context, emailID uuid.UUID) ([]*EmailLog, error)
}

type emailLogService struct {
	db *database.Queries
}

func NewEmailLogService(db *database.Queries) EmailLogService {
	return &emailLogService{
		db: db,
	}
}

func (s *emailLogService) CreateEmailLog(ctx context.Context, emailID uuid.UUID, status, message string) (*EmailLog, error) {
	dbEmailLog, err := s.db.CreateEmailLog(ctx, database.CreateEmailLogParams{
		EmailID: emailID,
		Status:  status,
		Message: sql.NullString{
			String: message,
			Valid:  message != "",
		},
	})
	if err != nil {
		return nil, err
	}

	return dbEmailLogToModel(dbEmailLog), nil
}

func (s *emailLogService) GetEmailLogsByEmailID(ctx context.Context, emailID uuid.UUID) ([]*EmailLog, error) {
	dbEmailLogs, err := s.db.GetEmailLogsByEmailID(ctx, emailID)
	if err != nil {
		return nil, err
	}

	var emailLogs []*EmailLog
	for _, dbEmailLog := range dbEmailLogs {
		emailLogs = append(emailLogs, dbEmailLogToModel(dbEmailLog))
	}

	return emailLogs, nil
}

func dbEmailLogToModel(dbEmailLog database.EmailLog) *EmailLog {
	message := ""
	if dbEmailLog.Message.Valid {
		message = dbEmailLog.Message.String
	}
	return &EmailLog{
		ID:        dbEmailLog.ID,
		EmailID:   dbEmailLog.EmailID,
		Status:    dbEmailLog.Status,
		Message:   message,
		CreatedAt: dbEmailLog.CreatedAt.Time,
	}
}
