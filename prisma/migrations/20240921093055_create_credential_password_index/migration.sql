-- CreateIndex
CREATE UNIQUE INDEX "credentials_credential_passoword_key" ON "credentials"("credential_type", "identity_id") WHERE ("credential_type" = 'PASSWORD');
