UPDATE selfservice_verification_flows SET nid = (SELECT id FROM networks LIMIT 1);