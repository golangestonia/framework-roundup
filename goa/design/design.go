package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("petstore", func() {
	Title("Pet Store Service")
	Description("HTTP service for managing pets")
	Server("pets", func() {
		Host("localhost", func() { URI("http://127.0.0.1:8080") })
	})
})

var Pet = Type("Pet", func() {
	Attribute("id", Int64)
	Attribute("name", String)
	Attribute("category", func() {
		Attribute("id", Int64)
		Attribute("name", String)
	})
	Attribute("status", String)
})

var ResultPet = ResultType("Pet", func() {
	Attribute("id", Int64)
	Attribute("name", String)
	Attribute("category", func() {
		Attribute("id", Int64)
		Attribute("name", String)
	})
	Attribute("status", String)
})

var _ = Service("pets", func() {
	Description("The pet management service")

	Method("AllPets", func() {
		Result(CollectionOf(ResultPet))
		HTTP(func() {
			GET("/pets")
			Response(StatusInternalServerError)
			Response(StatusOK)
		})
	})

	Method("PetByID", func() {
		Payload(func() {
			Attribute("id", Int)
			Required("id")
		})
		Result(ResultPet)
		HTTP(func() {
			GET("/pet/{id}")
			Response(StatusNotFound)
			Response(StatusInternalServerError)
			Response(StatusOK)
		})
	})

	Method("CreatePet", func() {
		Payload(Pet)
		Result(Int)
		HTTP(func() {
			PUT("/pet")
			Response(StatusNotFound)
			Response(StatusInternalServerError)
			Response(StatusOK)
		})
	})
})
