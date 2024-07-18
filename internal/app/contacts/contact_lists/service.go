package contact_lists

import (
	"context"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type ContactListService interface {
	AddContactToList(ctx context.Context, contactID, listID uuid.UUID) error
	RemoveContactFromList(ctx context.Context, contactID, listID uuid.UUID) error
	GetContactsByListID(ctx context.Context, listID uuid.UUID) ([]*database.Contact, error)
}

type contactListService struct {
	db *database.Queries
}

func NewContactListService(db *database.Queries) ContactListService {
	return &contactListService{
		db: db,
	}
}

func (s *contactListService) AddContactToList(ctx context.Context, contactID, listID uuid.UUID) error {
	return s.db.AddContactToList(ctx, database.AddContactToListParams{
		ContactID: contactID,
		ListID:    listID,
	})
}

func (s *contactListService) RemoveContactFromList(ctx context.Context, contactID, listID uuid.UUID) error {
	return s.db.RemoveContactFromList(ctx, database.RemoveContactFromListParams{
		ContactID: contactID,
		ListID:    listID,
	})
}

func (s *contactListService) GetContactsByListID(ctx context.Context, listID uuid.UUID) ([]*database.Contact, error) {
	contacts, err := s.db.GetContactsByListID(ctx, listID)
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
