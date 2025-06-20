CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE user_role AS ENUM ('receptionist', 'doctor');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    role user_role NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


-- Insert default doctor and receptionist accounts

INSERT INTO users (first_name, last_name, password_hash, email, role)
VALUES
('John', 'Doe',   '$2a$10$yci0ivo0LCrNWagrZ.G.c.uTMs8oMRxArXUcCSjpwQ.UdGjfl06pu', 'john.doe@hospital.com', 'doctor'),
('Lisa', 'Smith', '$2a$10$KQhlWMQbBwAcWVsWdILndub6MpljyBCRqq/IhIrpOrnKGl7OSc192', 'lisa.smith@hospital.com', 'receptionist');
-- Password for John: password123

-- Password for Lisa: reception123