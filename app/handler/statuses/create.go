package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/auth"
	"yatter-backend-go/app/domain/object"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
	Account *object.Account
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	// ボディのデコード時にエラーが発生した場合は、400 Bad Requestを返す
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// リクエストのコンテキストを取得
	ctx := r.Context()

	accountPtr := auth.AccountOf(ctx) // 認証情報を取得する
	// accountPtr: &{1 john $2a$10$of72ISxyPb7k1hZU39etrO2N8Kc9BfEHki/a2oA.LCFnhs4LpweWi <nil> <nil> <nil> <nil> 2024-08-07 13:31:39 +0900 JST}
	if accountPtr == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	req.Account = accountPtr

	fmt.Println("req: \n", req.Account)

	// AddStatusメソッドが実行されて、エラーが発生した場合は、500 Internal Server Errorを返す
	dto, err := h.statusUsecase.AddStatus(ctx, req.Status, req.Account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	// panic(fmt.Sprintf("Must Implement Status Creation And Check Acount Info %v", account_info))


	// レスポンスヘッダーにContent-Typeを設定
	w.Header().Set("Content-Type", "application/json")
	// レスポンスボディにエンコードされたJSONを書き込む
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}