ALTER TABLE "selfservice_login_request_methods" RENAME TO "selfservice_login_flow_methods";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_login_requests" RENAME TO "selfservice_login_flows";COMMIT TRANSACTION;BEGIN TRANSACTION;