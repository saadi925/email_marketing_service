package contact_segments

import (
	"context"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type ContactSegmentService interface {
	AddContactToSegment(ctx context.Context, contactID, segmentID uuid.UUID) error
	RemoveContactFromSegment(ctx context.Context, contactID, segmentID uuid.UUID) error
	GetContactsBySegmentID(ctx context.Context, segmentID uuid.UUID) ([]*database.Contact, error)
}

type contactSegmentService struct {
	db *database.Queries
}

func NewContactSegmentService(db *database.Queries) ContactSegmentService {
	return &contactSegmentService{
		db: db,
	}
}

func (s *contactSegmentService) AddContactToSegment(ctx context.Context, contactID, segmentID uuid.UUID) error {
	err := s.db.AddContactToSegment(ctx, database.AddContactToSegmentParams{
		ContactID: contactID,
		SegmentID: segmentID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *contactSegmentService) RemoveContactFromSegment(ctx context.Context, contactID, segmentID uuid.UUID) error {
	err := s.db.RemoveContactFromSegment(ctx, database.RemoveContactFromSegmentParams{
		ContactID: contactID,
		SegmentID: segmentID,
	})
	if err != nil {
		return err
	}

	return nil
}
func (s *contactSegmentService) GetContactsBySegmentID(ctx context.Context, segmentID uuid.UUID) ([]*database.Contact, error) {
	contacts, err := s.db.GetContactsBySegmentID(ctx, segmentID)
	if err != nil {
		return nil, err
	}

	var contactPtrs []*database.Contact
	for _, contact := range contacts {
		c := contact // create a new instance to avoid referencing the loop variable
		contactPtrs = append(contactPtrs, &c)
	}

	return contactPtrs, nil
}
