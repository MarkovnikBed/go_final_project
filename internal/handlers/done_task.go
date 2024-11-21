package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"final_project/internal/common"
)

func (h *Handler) DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if _, err := strconv.Atoi(id); err != nil || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.Response{Error: "неверно задан идентификатор задачи"})
		return
	}
	err := h.rep.DoneTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.Response{Error: "не удалось получить отметку о выполнение задачи " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(common.Response{})

}
