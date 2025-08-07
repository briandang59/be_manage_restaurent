-- Migration script for permissions table
-- Run this script in your PostgreSQL database

-- Drop existing table if needed (be careful with this in production!)
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;

-- Create permissions table
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create role_permissions junction table
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- Add some basic permissions
INSERT INTO permissions (name) VALUES
    ('create_user'),
    ('read_user'),
    ('update_user'),
    ('delete_user'),
    ('create_role'),
    ('read_role'),
    ('update_role'),
    ('delete_role'),
    ('manage_permissions')
ON CONFLICT (name) DO NOTHING;
