package handlersplit

import (
	"encoding/json"
	"net/http"

	"github.com/harness/gitness/app/api/controller/split"
	"github.com/harness/gitness/app/api/render"
	"github.com/harness/gitness/app/api/request"
)

func HandleCreateWorkspace(splitCtrl *split.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, _ := request.AuthSessionFrom(ctx)

		in := new(split.CreateWorkspaceInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequestf(ctx, w, "Invalid request body: %s.", err)
			return
		}

		workspace, err := splitCtrl.CreateWorkspace(ctx, session, in)
		if err != nil {
			render.TranslatedUserError(ctx, w, err)
			return
		}

		render.JSON(w, http.StatusCreated, workspace)
	}
}

func HandleGetWorkspace(splitCtrl *split.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, _ := request.AuthSessionFrom(ctx)

		id, err := request.GetUUIDParam(r)
		if err != nil {
			render.BadRequestf(ctx, w, "Invalid request body: %s.", err)
			return
		}

		workspace, err := splitCtrl.GetWorkspace(ctx, session, id)
		if err != nil {
			render.TranslatedUserError(ctx, w, err)
			return
		}

		render.JSON(w, http.StatusOK, workspace)
	}
}
