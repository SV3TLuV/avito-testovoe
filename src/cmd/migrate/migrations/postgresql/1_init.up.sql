CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
            CREATE TYPE organization_type AS ENUM (
                'IE',
                'LLC',
                'JSC'
                );
        END IF;
    END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_service_type') THEN
            CREATE TYPE tender_service_type AS ENUM (
                'Construction',
                'Delivery',
                'Manufacture'
                );
        END IF;
    END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_status') THEN
            CREATE TYPE tender_status AS ENUM (
                'Created',
                'Published',
                'Closed'
                );
        END IF;
    END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_status') THEN
            CREATE TYPE bid_status AS ENUM (
                'Created',
                'Published',
                'Canceled'
                );
        END IF;
    END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_decision') THEN
            CREATE TYPE bid_decision AS ENUM (
                'Approved',
                'Rejected'
                );
        END IF;
    END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'author_type') THEN
            CREATE TYPE author_type AS ENUM (
                'Organization',
                'User'
                );
        END IF;
    END $$;
