package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dobefu/csb/cmd/api"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/server/validation"
)

func GetEntryByUrl(w http.ResponseWriter, r *http.Request) {
	params, err := validation.CheckRequiredQueryParams(
		r,
		"url",
		"locale",
	)

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	url := params["url"].([]string)[0]
	locale := params["locale"].([]string)[0]

	entry, err := api.GetEntryByUrl(url, locale, false)

	if err != nil {
		fmt.Fprintf(w, `{"data": null, "error": "%s"}`, err.Error())
		return
	}

	output := map[string]interface{}{
		"data": map[string]interface{}{
			"entry": entry,
		},
		"error": nil,
	}

	json, err := json.Marshal(output)

	if err != nil {
		logger.Error(err.Error())
		fmt.Fprintf(w, `{"data": null, "error": "%s"}`, err.Error())
		return
	}

	fmt.Fprint(w, string(json))
}
