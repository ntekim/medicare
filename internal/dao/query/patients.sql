-- name: CreatePatient :one
INSERT INTO patients (
    id, first_name, last_name, date_of_birth, gender, contact_number, address, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
) RETURNING id;

-- name: GetPatientByID :one
SELECT * FROM patients WHERE id = $1;

-- name: ListPatients :many
SELECT * FROM patients ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: UpdatePatient :exec
UPDATE patients
SET first_name = $2,
    last_name = $3,
    date_of_birth = $4,
    gender = $5,
    contact_number = $6,
    address = $7,
    updated_at = NOW()
WHERE id = $1;

-- name: DeletePatient :exec
DELETE FROM patients WHERE id = $1;
