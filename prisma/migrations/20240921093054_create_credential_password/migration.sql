/*
  Warnings:

  - A unique constraint covering the columns `[username]` on the table `identities` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[email]` on the table `identities` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[phone]` on the table `identities` will be added. If there are existing duplicate values, this will fail.

*/
-- AlterEnum
ALTER TYPE "CredentialType" ADD VALUE 'PASSWORD';

-- AlterTable
ALTER TABLE "identities" ADD COLUMN     "email" TEXT,
ADD COLUMN     "phone" TEXT,
ADD COLUMN     "username" TEXT;

-- CreateTable
CREATE TABLE "credentials_password" (
    "id" SERIAL NOT NULL,
    "password" VARCHAR(255) NOT NULL,

    CONSTRAINT "credentials_password_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "identities_username_key" ON "identities"("username");

-- CreateIndex
CREATE UNIQUE INDEX "identities_email_key" ON "identities"("email");

-- CreateIndex
CREATE UNIQUE INDEX "identities_phone_key" ON "identities"("phone");
