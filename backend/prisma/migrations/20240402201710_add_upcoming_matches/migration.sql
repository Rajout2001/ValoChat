-- CreateTable
CREATE TABLE "upcomingMatches" (
    "id" TEXT NOT NULL,
    "team1" TEXT NOT NULL,
    "team2" TEXT NOT NULL,
    "score1" INTEGER NOT NULL,
    "score2" INTEGER NOT NULL,
    "flag1" TEXT NOT NULL,
    "flag2" TEXT NOT NULL,
    "time_until_match" TEXT NOT NULL,
    "round_info" TEXT NOT NULL,
    "tournament_name" TEXT NOT NULL,
    "match_page" TEXT NOT NULL,
    "tournament_icon" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "upcomingMatches_pkey" PRIMARY KEY ("id")
);
