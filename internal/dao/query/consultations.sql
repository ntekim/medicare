-- name: CreateConsultation :one
INSERT INTO consultations (
    id, patient_id, doctor_id, vitals, diagnosis, prescription, notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: ListConsultationsByPatient :many
SELECT * FROM consultations
WHERE patient_id = $1
ORDER BY consultation_date DESC;

-- name: GetConsultation :one
SELECT * FROM consultations
WHERE id = $1;

-- name: UpdateConsultation :exec
UPDATE consultations
SET vitals = $2,
    diagnosis = $3,
    prescription = $4,
    notes = $5
WHERE id = $1;

-- name: DeleteConsultation :exec
DELETE FROM consultations WHERE id = $1;
