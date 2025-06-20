-- Remove seeded users on down migration

DELETE FROM users
WHERE email IN ('john.doe@hospital.com', 'lisa.smith@hospital.com');

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_role;