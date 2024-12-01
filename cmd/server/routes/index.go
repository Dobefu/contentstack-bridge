package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, apiPath string) {
	output := map[string]interface{}{
		"data": map[string]interface{}{
			"api_endpoints": []string{apiPath},
		},
		"error": nil,
	}

	o, err := json.Marshal(output)

	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, string(o))
}
