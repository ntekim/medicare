CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE consultations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    doctor_id UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    consultation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    vitals JSONB,
    diagnosis TEXT,
    prescription TEXT,
    notes TEXT
);