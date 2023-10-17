package frameworkright

const (
	queryToAddAccessToken        = "INSERT INTO tokens.access_token (access_uuid,access_token,at_expires) VALUES (?,?,?); "
	queryToAddRefreshToken       = "INSERT INTO tokens.refresh_token (refresh_uuid,refresh_token,rf_expires) VALUES (?,?,?); "
	queryToDeleteAccessToken     = "DELETE access_uuid,access_token,at_expires FROM tokens.access_token WHERE access_uuid =?;"
	queryToDeleteRefreshToken    = "DELETE refresh_uuid,refresh_token,rf_expires FROM tokens.refesh_token WHERE refresh_uuid =?;"
	QueryToVerifyAccessTokenUUID = "SELECT access_uuid FROM tokens.access_token where access_uuid=?;"
)
