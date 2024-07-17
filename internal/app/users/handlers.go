package users

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func protectedService(w http.ResponseWriter, _ *http.Request, userID uuid.UUID, db *database.Queries) {
	utils.RespondJSON(w, 200, "User ID: "+userID.String())

}
