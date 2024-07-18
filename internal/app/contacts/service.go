package contacts

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type ContactService interface {
	CreateContact(ctx context.Context, firstName, lastName string, subscribed, blocklisted bool, email, whatsapp, landlineNumber string) (*Contact, error)
	GetContactByID(ctx context.Context, id uuid.UUID) (*Contact, error)
	GetContacts(ctx context.Context) ([]*Contact, error)
	UpdateContact(ctx context.Context, id uuid.UUID, firstName, lastName string, subscribed, blocklisted bool, email, whatsapp, landlineNumber string) (*Contact, error)
	DeleteContact(ctx context.Context, id uuid.UUID) error
}

type contactService struct {
	db *database.Queries
}

func NewContactService(db *database.Queries) ContactService {
	return &contactService{
		db: db,
	}
}

func (s *contactService) CreateContact(ctx context.Context, firstName, lastName string, subscribed, blocklisted bool, email, whatsapp, landlineNumber string) (*Contact, error) {
	c := database.CreateContactParams{
		FirstName: sql.NullString{
			String: firstName,
			Valid:  firstName != "",
		},
		LastName: sql.NullString{
			String: lastName,
			Valid:  lastName != "",
		},
		Subscribed: sql.NullBool{
			Bool:  subscribed,
			Valid: true,
		},
		Blocklisted: sql.NullBool{
			Bool:  blocklisted,
			Valid: true,
		},
		Email: email,
		Whatsapp: sql.NullString{
			String: whatsapp,
			Valid:  whatsapp != "",
		},
		LandlineNumber: sql.NullString{
			String: landlineNumber,
			Valid:  landlineNumber != "",
		},
	}
	dbContact, err := s.db.CreateContact(ctx, c)
	if err != nil {
		return nil, err
	}

	return dbContactToModel(dbContact), nil
}

func (s *contactService) GetContactByID(ctx context.Context, id uuid.UUID) (*Contact, error) {
	dbContact, err := s.db.GetContactByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dbContactToModel(dbContact), nil
}

func (s *contactService) GetContacts(ctx context.Context) ([]*Contact, error) {
	dbContacts, err := s.db.GetContacts(ctx)
	if err != nil {
		return nil, err
	}

	var contacts []*Contact
	for _, dbContact := range dbContacts {
		contacts = append(contacts, dbContactToModel(dbContact))
	}

	return contacts, nil
}

func (s *contactService) UpdateContact(ctx context.Context, id uuid.UUID, firstName, lastName string, subscribed, blocklisted bool, email, whatsapp, landlineNumber string) (*Contact, error) {
	dbContact, err := s.db.UpdateContact(ctx, database.UpdateContactParams{
		ID: id,
		FirstName: sql.NullString{
			String: firstName,
			Valid:  firstName != "",
		},
		LastName: sql.NullString{
			String: lastName,
			Valid:  lastName != "",
		},
		Subscribed: sql.NullBool{
			Bool:  subscribed,
			Valid: true,
		},
		Blocklisted: sql.NullBool{
			Bool:  blocklisted,
			Valid: true,
		},
		Email: email,
		Whatsapp: sql.NullString{
			String: whatsapp,
			Valid:  whatsapp != "",
		},
		LandlineNumber: sql.NullString{
			String: landlineNumber,
			Valid:  landlineNumber != "",
		},
	})
	if err != nil {
		return nil, err
	}

	return dbContactToModel(dbContact), nil
}

func (s *contactService) DeleteContact(ctx context.Context, id uuid.UUID) error {
	err := s.db.DeleteContact(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func dbContactToModel(dbContact database.Contact) *Contact {
	return &Contact{
		ID:             dbContact.ID,
		FirstName:      dbContact.FirstName.String,
		LastName:       dbContact.LastName.String,
		Subscribed:     dbContact.Subscribed.Bool,
		Blocklisted:    dbContact.Blocklisted.Bool,
		Email:          dbContact.Email,
		Whatsapp:       dbContact.Whatsapp.String,
		LandlineNumber: dbContact.LandlineNumber.String,
		LastChanged:    dbContact.LastChanged.Time,
		DateAdded:      dbContact.DateAdded.Time,
	}
}
