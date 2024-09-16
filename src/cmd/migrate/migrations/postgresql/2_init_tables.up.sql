CREATE TABLE IF NOT EXISTS employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS organization (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    service_type tender_service_type NOT NULL,
    status tender_status NOT NULL,
    organization_id UUID NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tender_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tender_id UUID NOT NULL REFERENCES tender(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    service_type tender_service_type NOT NULL,
    status tender_status NOT NULL,
    organization_id UUID NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (tender_id, version)
);

CREATE TABLE IF NOT EXISTS employee_tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID REFERENCES employee(id) ON DELETE CASCADE,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bid (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status bid_status NOT NULL,
    tender_id UUID NOT NULL REFERENCES tender(id) ON DELETE CASCADE,
    author_type author_type NOT NULL,
    author_id UUID NOT NULL REFERENCES employee(id) ON DELETE CASCADE,
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bid_history (
    id UUID NOT NULL,
    bid_id UUID NOT NULL REFERENCES bid(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status bid_status NOT NULL,
    tender_id UUID NOT NULL REFERENCES tender(id) ON DELETE CASCADE,
    author_type author_type NOT NULL,
    author_id UUID NOT NULL REFERENCES employee(id) ON DELETE CASCADE,
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (bid_id, version)
);

CREATE TABLE IF NOT EXISTS bid_review (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bid_id UUID NOT NULL REFERENCES bid(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bid_employee_decision (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bid_id uuid NOT NULL REFERENCES bid(id) ON DELETE CASCADE,
    employee_id uuid NOT NULL REFERENCES bid(id) ON DELETE CASCADE,
    decision bid_decision NOT NULL,
    UNIQUE (bid_id, employee_id)
);