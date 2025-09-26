CREATE TABLE contract_documents (
    document_hash CHAR(64) DEFAULT NULL,
    contract_id UUID DEFAULT NULL,
    CONSTRAINT fk_contract_document_documents FOREIGN KEY (document_hash) REFERENCES documents(hash) ON DELETE CASCADE,
    CONSTRAINT fk_contract_document_contracts FOREIGN KEY (contract_id) REFERENCES contracts(id) ON DELETE CASCADE,

    url TEXT NOT NULL,
    category SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);