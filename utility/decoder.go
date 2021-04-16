package utility

import (
	"avoxi-interview/models"
	"encoding/json"
	"net/http"
)

// Normally I would only go to the trouble of a wrapper function if I were to have several api calls that
// need to decode the same payload, but it does clean up the code ever so slightly, and this is after all
// an audition.  Please excuse the vanity.
func PayloadDecoder(w http.ResponseWriter, r *http.Request, p *models.Payload, operatedFunction func() ()) {
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	operatedFunction()
}
