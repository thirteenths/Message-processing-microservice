package https

import (
	"encoding/json"
	"net/http"

	"github.com/thirteenths/message-processing-microservice/internal/app"
	"github.com/thirteenths/message-processing-microservice/internal/domains/request"
	_ "github.com/thirteenths/message-processing-microservice/internal/domains/response"
)

type ApiHandler struct {
	messageService app.MessageService
}

func NewApiHandler(message app.MessageService) *ApiHandler {
	return &ApiHandler{messageService: message}
}

// CreateMessage
// @Summary      CreateMessage
// @Description  Create messageService
// @Tags         messageService
// @Accept       json
// @Produce      json
// @Param        request.CreateMessage	body		request.CreateMessage	true	"Add messageService"
// @Success      201  {string}  string    "ok"
// @Failure      400  {string}  string    "bad request"
// @Failure      500  {string}  string    "internal server error"
// @Router       /message [post]
func (h *ApiHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req request.CreateMessage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := h.messageService.CreateMessage(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetMessage
// @Summary      GetMessage
// @Description  Get message
// @Tags         messageService
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.GetMessage
// @Failure      400  {string}  string    "bad request"
// @Failure      500  {string}  string    "internal server error"
// @Router       /message [get]
func (h *ApiHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	res, err := h.messageService.GetMessage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetStatistic
// @Summary      GetStatistic
// @Description  Get statistic
// @Tags         messageService
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.GetStatistic
// @Failure      400  {string}  string    "bad request"
// @Failure      500  {string}  string    "internal server error"
// @Router       /statistic [get]
func (h *ApiHandler) GetStatistic(w http.ResponseWriter, r *http.Request) {
	res, err := h.messageService.GetStatistic()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
