/*
  Warnings:

  - Added the required column `matchId` to the `Chat` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Chat" ADD COLUMN     "matchId" TEXT NOT NULL;
