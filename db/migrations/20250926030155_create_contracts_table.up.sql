CREATE TABLE contracts (
    id UUID PRIMARY KEY,
    human_id CHAR(15) NOT NULL UNIQUE,
    email VARCHAR(60) NOT NULL,
    title VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    notes TEXT,
    risk_detection TEXT,
    summarize TEXT,
    risk_level SMALLINT NOT NULL DEFAULT 0,
    status SMALLINT NOT NULL DEFAULT 0,
    application_status SMALLINT NOT NULL DEFAULT 0,
    category SMALLINT NOT NULL,
    start_date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contracts_title ON contracts(title);
CREATE INDEX idx_contracts_company ON contracts(company);