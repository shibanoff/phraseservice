package service

import (
	"net/http"
	"phraseservice/hashcode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetPhraseHashArgs struct {
	Phrase string `json:"phrase"`
	Hash   int64  `json:"hash"`
}

func (args *GetPhraseHashArgs) Validate() error {
	return validation.ValidateStruct(args,
		validation.Field(&args.Phrase, validation.Required),
	)
}

func PhraseHashHandler(w http.ResponseWriter, r *http.Request) {
	var args GetPhraseHashArgs

	if err := decodeRequest(r, &args); err != nil {
		sendError(w, err)
		return
	}

	hashCode := hashcode.NewSHA265()
	hashCode.LoadString(args.Phrase)

	args.Hash = hashCode.GetCode()

	if err := sendJSONResponse(w, args); err != nil {
		sendError(w, err)
	}
}
