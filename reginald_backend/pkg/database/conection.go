package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Fallback for local development if not provided
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db
	log.Println("Database connected successfully")
	return db, nil
}

func InitDB() {
	if DB == nil {
		log.Fatal("Database not connected")
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role TEXT NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL,
		created_by TEXT NOT NULL REFERENCES users(email) ON DELETE CASCADE,
		assigned_to TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Fix constraints and columns
	DO $$ 
	BEGIN 
		-- Add assigned_to column if it doesn't exist
		IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='tasks' AND column_name='assigned_to') THEN
			ALTER TABLE tasks ADD COLUMN assigned_to TEXT;
		END IF;

		-- Data Migration: Update tasks.created_by from ID to Email if it matches a user ID
		UPDATE tasks t SET created_by = u.email FROM users u WHERE t.created_by = u.id AND t.created_by NOT LIKE '%@%';
		-- Same for assigned_to
		UPDATE tasks t SET assigned_to = u.email FROM users u WHERE t.assigned_to = u.id AND t.assigned_to NOT LIKE '%@%';

		-- Drop old constraints if they point to users(id)
		IF EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='tasks_created_by_fkey' AND table_name='tasks') THEN
			ALTER TABLE tasks DROP CONSTRAINT tasks_created_by_fkey;
		END IF;
		
		-- Try to add the new constraint.
		BEGIN
			ALTER TABLE tasks ADD CONSTRAINT tasks_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(email) ON DELETE CASCADE;
		EXCEPTION WHEN OTHERS THEN
			RAISE NOTICE 'Could not add tasks_created_by_fkey: %', SQLERRM;
		END;
	END $$;

	CREATE TABLE IF NOT EXISTS requests (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL,
		created_by TEXT NOT NULL REFERENCES users(email) ON DELETE CASCADE,
		handled_by TEXT REFERENCES users(email) ON DELETE SET NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Fix requests constraints and data
	DO $$ 
	BEGIN 
		-- Data Migration
		UPDATE requests r SET created_by = u.email FROM users u WHERE r.created_by = u.id AND r.created_by NOT LIKE '%@%';
		UPDATE requests r SET handled_by = u.email FROM users u WHERE r.handled_by = u.id AND r.handled_by NOT LIKE '%@%';

		IF EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='requests_created_by_fkey' AND table_name='requests') THEN
			ALTER TABLE requests DROP CONSTRAINT requests_created_by_fkey;
		END IF;
		
		BEGIN
			ALTER TABLE requests ADD CONSTRAINT requests_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(email) ON DELETE CASCADE;
		EXCEPTION WHEN OTHERS THEN
			RAISE NOTICE 'Could not add requests_created_by_fkey: %', SQLERRM;
		END;

		IF EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='requests_handled_by_fkey' AND table_name='requests') THEN
			ALTER TABLE requests DROP CONSTRAINT requests_handled_by_fkey;
		END IF;
		
		BEGIN
			ALTER TABLE requests ADD CONSTRAINT requests_handled_by_fkey FOREIGN KEY (handled_by) REFERENCES users(email) ON DELETE SET NULL;
		EXCEPTION WHEN OTHERS THEN
			RAISE NOTICE 'Could not add requests_handled_by_fkey: %', SQLERRM;
		END;
	END $$;

	CREATE TABLE IF NOT EXISTS audit_logs (
		id TEXT PRIMARY KEY,
		user_email TEXT NOT NULL,
		action TEXT NOT NULL,
		resource TEXT NOT NULL,
		resource_id TEXT,
		status_code INTEGER DEFAULT 200,
		duration_ms INTEGER DEFAULT 0,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Update audit_logs with new columns if they don't exist
	DO $$ 
	BEGIN 
		IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='audit_logs' AND column_name='status_code') THEN
			ALTER TABLE audit_logs ADD COLUMN status_code INTEGER DEFAULT 200;
		END IF;
		IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='audit_logs' AND column_name='duration_ms') THEN
			ALTER TABLE audit_logs ADD COLUMN duration_ms INTEGER DEFAULT 0;
		END IF;
	END $$;`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	log.Println("Database tables initialized")
}
