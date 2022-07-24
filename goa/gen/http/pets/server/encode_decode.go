// Code generated by goa v3.7.13, DO NOT EDIT.
//
// pets HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/golangestonia/framework-roundup/goa/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	petsviews "github.com/golangestonia/framework-roundup/goa/gen/pets/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAllPetsResponse returns an encoder for responses returned by the pets
// AllPets endpoint.
func EncodeAllPetsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(petsviews.PetCollection)
		enc := encoder(ctx, w)
		body := NewPetResponseCollection(res.Projected)
		w.WriteHeader(http.StatusInternalServerError)
		return enc.Encode(body)
	}
}

// EncodePetByIDResponse returns an encoder for responses returned by the pets
// PetByID endpoint.
func EncodePetByIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*petsviews.Pet)
		enc := encoder(ctx, w)
		body := NewPetByIDNotFoundResponseBody(res.Projected)
		w.WriteHeader(http.StatusNotFound)
		return enc.Encode(body)
	}
}

// DecodePetByIDRequest returns a decoder for requests sent to the pets PetByID
// endpoint.
func DecodePetByIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewPetByIDPayload(id)

		return payload, nil
	}
}

// EncodeCreatePetResponse returns an encoder for responses returned by the
// pets CreatePet endpoint.
func EncodeCreatePetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusNotFound)
		return enc.Encode(body)
	}
}

// DecodeCreatePetRequest returns a decoder for requests sent to the pets
// CreatePet endpoint.
func DecodeCreatePetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreatePetRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewCreatePetPet(&body)

		return payload, nil
	}
}

// marshalPetsviewsPetViewToPetResponse builds a value of type *PetResponse
// from a value of type *petsviews.PetView.
func marshalPetsviewsPetViewToPetResponse(v *petsviews.PetView) *PetResponse {
	res := &PetResponse{
		ID:     v.ID,
		Name:   v.Name,
		Status: v.Status,
	}
	if v.Category != nil {
		res.Category = &struct {
			ID   *int64
			Name *string
		}{
			ID:   v.Category.ID,
			Name: v.Category.Name,
		}
	}

	return res
}
