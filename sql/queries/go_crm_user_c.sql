-- name: GetUserByEmailSQLC :one
Select usr_email, usr_id
from `go_crm_user`
where usr_email = ? limit 1;

-- name: UpdateUserStatusByUserID :exec
Update `go_crm_user`
set usr_status = $2,
    usr_updated_at = $3
where usr_id = $1;

