-- Migration script to add missing salary_per_hour column
-- Run this script in your PostgreSQL database

-- Add the missing column
ALTER TABLE employees ADD COLUMN IF NOT EXISTS salary_per_hour BIGINT DEFAULT 0;

-- Update existing records if needed (optional)
-- UPDATE employees SET salary_per_hour = 0 WHERE salary_per_hour IS NULL;

-- Verify the column was added
SELECT column_name, data_type 
FROM information_schema.columns 
WHERE table_name = 'employees' AND column_name = 'salary_per_hour'; 