package notifications

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/middlewares"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

func CreateNotificationHandler(service NotificationService) middlewares.AuthHandler {
	return func(w http.ResponseWriter, r *http.Request, u uuid.UUID) {
		var req struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid input")
			return
		}

		notification, err := service.createNotification(r.Context(), u, req.Message)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create notification")
			return
		}

		utils.RespondJSON(w, http.StatusCreated, notification)
	}
}

func UpdateNotificationReadStatusHandler(service NotificationService) middlewares.AuthHandler {
	return func(w http.ResponseWriter, r *http.Request, u uuid.UUID) {
		var req struct {
			ID   int  `json:"id"`
			Read bool `json:"read"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid input")
			return
		}

		err := service.updateNotificationReadStatus(r.Context(), req.ID, req.Read)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to update notification read status")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetNotificationsHandler(service NotificationService) middlewares.AuthHandler {
	return func(w http.ResponseWriter, r *http.Request, u uuid.UUID) {
		notifications, err := service.getNotificationsByUserID(r.Context(), u)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve notifications")
			return
		}

		utils.RespondJSON(w, http.StatusOK, notifications)
	}
}
