package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type GetPublicRequest struct {
	OnlyMedia bool 
	SinceID   int  
	Limit     int  
}


func (h *handler) GetPublic(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	onlyMediaStr := r.URL.Query().Get("only_media")
	if onlyMediaStr == "" {
		onlyMediaStr = "false"
	}
	onlyMedia, err := strconv.ParseBool(onlyMediaStr)
	if err != nil {
		http.Error(w, "only_media must be a boolean", http.StatusBadRequest)
		return
	}

	sinceIDStr := r.URL.Query().Get("since_id")
	if sinceIDStr == "" {
		sinceIDStr = "0"
	}
	sinceID, err := strconv.Atoi(sinceIDStr)
	if err != nil {
		http.Error(w, "since_id must be a number", http.StatusBadRequest)
		return
	} else if sinceID < 0 {
		sinceID = 0
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "40"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "limit must be a number", http.StatusBadRequest)
		return
	}else if limit > 80 {
		limit = 80
	}else if limit < 1 {
		limit = 1
	}

	ctx := r.Context()

	dto, err := h.timelineUsecase.GetPublic(ctx, onlyMedia, sinceID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}