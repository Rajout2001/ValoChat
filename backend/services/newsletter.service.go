package services

import (
	"encoding/json"
	"net/http"
	prisma "pc3r/prisma"
	db "pc3r/prisma/db"
)

func subscribeNewsletter (res http.ResponseWriter, req *http.Request) {
	var body SubscribeNewsletterSchema
	err := json.NewDecoder(req.Body).Decode(&body)
	if (err != nil) || (body.Email == "") {
		res.WriteHeader(http.StatusBadRequest)
		message := "Veuillez vérifier l'émail."
		json.NewEncoder(res).Encode(CustomError(message, INPUT_ERROR))
		return
	}
	prisma, ctx := prisma.GetPrisma()
	// vérifier que l'utilisateur n'est pas déjà inscrit
	subsriber , _ := prisma.NewsLetter.FindFirst(
		db.NewsLetter.Email.Equals(body.Email),
	).Exec(ctx)
	
	if (subsriber != nil) {
		res.WriteHeader(http.StatusConflict)
		message := "Vous êtes déjà inscrit à la newsletter"
		json.NewEncoder(res).Encode(CustomError(message, INPUT_ERROR))
		return
	}

	// inscrire l'utilisateur à la newsletter
	_ , err = prisma.NewsLetter.CreateOne(
		db.NewsLetter.Email.Set(body.Email),
	).Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		message := "Erreur lors de l'inscription à la newsletter"
		json.NewEncoder(res).Encode(CustomError(message, INPUT_ERROR))
		return
	}

	newsletterResponse := NewsletterResponse{
		Code: "201",
		Message: "Inscription à la newsletter réussie",
	}
	
	// l'utilisateur est bien inscrit à la newsletter
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newsletterResponse)
}