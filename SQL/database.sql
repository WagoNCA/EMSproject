CREATE DATABASE EMSproject;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE site_type AS ENUM (
    'office',
    'factory',
    'warehouse',
    'other'
);

CREATE TYPE meter_type AS ENUM (
    'electricity',
    'gas',
    'water'
);


CREATE TYPE meter_unit AS ENUM (
    'kWh',
    'm3'
);


CREATE TABLE site (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    type site_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE meter (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    site_id UUID REFERENCES site(id),
    type meter_type,
    unit meter_unit,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
