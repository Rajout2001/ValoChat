package prisma

import (
	"context"
	"pc3r/prisma/db"
)

// code de la connexion avec Prisma est inspiré d'un étudiant dans la salle TME.
var prisma *db.PrismaClient
var ctx context.Context

func Init() {
	prisma = db.NewClient()
	ctx = context.Background()
	prisma.Prisma.Connect()
}

func GetPrisma() (*db.PrismaClient, context.Context) {
	if prisma == nil {
		Init()
	}
	return prisma, ctx
}