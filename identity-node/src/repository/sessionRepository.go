package repository

import (
	"identity-node/src/model"
)

// SaveSession - Save user session on "db"
// @findme : servirebbe un controllo consistenza db: user,token_session è univoco, non si dovrebbe inserire una riga duplicata
// da riattenzionare quando si inserisce il db
func SaveSession(s model.Session) (*model.Session, error) {
	_, err := FindUser(s.User)

	if err != nil {
		return nil, err
	}

	sessionList = append(sessionList, s)
	// @findme : una vera sessione dovrebbe scadere, da risistemare quando si mette il db

	return &s, nil
}
