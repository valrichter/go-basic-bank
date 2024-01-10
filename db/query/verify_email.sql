-- name: CreateVerifyEmail :one
INSERT INTO
    verify_email (username, email, secret_code)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE verify_email
SET is_used = TRUE
WHERE
    id = @id
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expires_at > now() RETURNING *;