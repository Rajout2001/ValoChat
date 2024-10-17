/*
  Warnings:

  - A unique constraint covering the columns `[chat_id]` on the table `matchResult` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `chat_id` to the `matchResult` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "matchResult" ADD COLUMN     "chat_id" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "matchResult_chat_id_key" ON "matchResult"("chat_id");

-- AddForeignKey
ALTER TABLE "matchResult" ADD CONSTRAINT "matchResult_chat_id_fkey" FOREIGN KEY ("chat_id") REFERENCES "Chat"("id") ON DELETE CASCADE ON UPDATE CASCADE;
