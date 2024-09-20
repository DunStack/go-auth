-- CreateEnum
CREATE TYPE "CredentialType" AS ENUM ('OAUTH');

-- CreateTable
CREATE TABLE "identities" (
    "id" SERIAL NOT NULL,

    CONSTRAINT "identities_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "credentials" (
    "id" SERIAL NOT NULL,
    "identity_id" INTEGER NOT NULL,
    "credential_type" "CredentialType" NOT NULL,
    "credential_id" INTEGER NOT NULL,

    CONSTRAINT "credentials_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "credentials_oauth" (
    "id" SERIAL NOT NULL,
    "provider" TEXT NOT NULL,
    "uid" TEXT NOT NULL,

    CONSTRAINT "credentials_oauth_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "credentials_credential_type_credential_id_key" ON "credentials"("credential_type", "credential_id");

-- CreateIndex
CREATE UNIQUE INDEX "credentials_oauth_provider_uid_key" ON "credentials_oauth"("provider", "uid");

-- AddForeignKey
ALTER TABLE "credentials" ADD CONSTRAINT "credentials_identity_id_fkey" FOREIGN KEY ("identity_id") REFERENCES "identities"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
