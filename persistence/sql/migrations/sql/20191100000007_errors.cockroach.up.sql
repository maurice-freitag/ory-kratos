ALTER TABLE "selfservice_errors" ADD COLUMN "csrf_token" VARCHAR (255) NOT NULL DEFAULT '';COMMIT TRANSACTION;BEGIN TRANSACTION;