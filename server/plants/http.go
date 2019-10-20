package plants

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dkalytovskyi/architecture-lab-2/server/tools"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(greenHouse *GreenHouse) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListPlants(greenHouse, rw)
		} else if r.Method == "POST" {
			handleAddMoistureLevel(r, rw, greenHouse)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleAddMoistureLevel(r *http.Request, rw http.ResponseWriter, greenHouse *GreenHouse) {
	var c Plant
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := greenHouse.AddMoistureLevel(c.Id, c.SoilMoistureLevel, c.SoilDataTimestamp)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListPlants(greenHouse *GreenHouse, rw http.ResponseWriter) {
	res, err := greenHouse.ListCriticalPlants()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
