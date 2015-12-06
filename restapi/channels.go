package restapi

import (
	"encoding/json"
	"net/http"
)

func GetChannelsHandler(ctx *AppContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelName := r.URL.Query().Get("name")

	//params := mux.Vars(r)
	//params["channelName"]
	channels := ctx.GCTSDBServer.GetChannels(channelName)

	res := []Channel{}
	for _, ch := range channels {
		rch, err := NewRestChannel(&ch)
		if err != nil {
			err.WriteResponse(w)
			return
		}
		res = append(res, *rch)
	}

	str, jsonErr := json.Marshal(res)
	if jsonErr != nil {
		LogReq(r).Errorf("Could not json marshal object: %s", jsonErr)
		NewHttpError(500, "UNKNOWN_ERROR", "Internal Server Error").WriteResponse(w)
		return
	}

	w.WriteHeader(200)
	w.Write(str)
}
