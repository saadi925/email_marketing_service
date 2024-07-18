package segments

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type Segment struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Criteria    string    `json:"criteria"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SegmentService interface {
	CreateSegment(ctx context.Context, name, description, criteria string) (*Segment, error)
	GetSegmentByID(ctx context.Context, id uuid.UUID) (*Segment, error)
	GetSegments(ctx context.Context) ([]*Segment, error)
	UpdateSegment(ctx context.Context, id uuid.UUID, name, description, criteria string) (*Segment, error)
	DeleteSegment(ctx context.Context, id uuid.UUID) error
}

type segmentService struct {
	db *database.Queries
}

func NewSegmentService(db *database.Queries) SegmentService {
	return &segmentService{
		db: db,
	}
}

func (s *segmentService) CreateSegment(ctx context.Context, name, description, criteria string) (*Segment, error) {
	dbSegment, err := s.db.CreateSegment(ctx, database.CreateSegmentParams{
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		Criteria: json.RawMessage(criteria),
	})
	if err != nil {
		return nil, err
	}

	return dbSegmentToModel(dbSegment), nil
}

func (s *segmentService) GetSegmentByID(ctx context.Context, id uuid.UUID) (*Segment, error) {
	dbSegment, err := s.db.GetSegmentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dbSegmentToModel(dbSegment), nil
}

func (s *segmentService) GetSegments(ctx context.Context) ([]*Segment, error) {
	dbSegments, err := s.db.GetSegments(ctx)
	if err != nil {
		return nil, err
	}

	var segments []*Segment
	for _, dbSegment := range dbSegments {
		segments = append(segments, dbSegmentToModel(dbSegment))
	}

	return segments, nil
}

func (s *segmentService) UpdateSegment(ctx context.Context, id uuid.UUID, name, description, criteria string) (*Segment, error) {
	dbSegment, err := s.db.UpdateSegment(ctx, database.UpdateSegmentParams{
		ID:   id,
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		Criteria: json.RawMessage(criteria),
	})
	if err != nil {
		return nil, err
	}

	return dbSegmentToModel(dbSegment), nil
}

func (s *segmentService) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	err := s.db.DeleteSegment(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func dbSegmentToModel(dbSegment database.Segment) *Segment {
	return &Segment{
		ID:          dbSegment.ID,
		Name:        dbSegment.Name,
		Description: dbSegment.Description.String,
		Criteria:    string(dbSegment.Criteria),
		CreatedAt:   dbSegment.CreatedAt.Time,
		UpdatedAt:   dbSegment.UpdatedAt.Time,
	}
}
