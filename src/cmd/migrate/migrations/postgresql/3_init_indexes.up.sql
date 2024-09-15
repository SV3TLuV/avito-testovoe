-- organization_responsible
CREATE INDEX idx_organization_responsible_organization_id ON organization_responsible (organization_id);
CREATE INDEX idx_organization_responsible_user_id ON organization_responsible (user_id);


-- tender
CREATE INDEX idx_tender_organization_id ON tender (organization_id);
CREATE INDEX idx_tender_service_type ON tender (service_type);
CREATE INDEX idx_tender_status ON tender (status);


-- tender_history
CREATE INDEX idx_tender_history_organization_id ON tender_history (organization_id);
CREATE INDEX idx_tender_history_version ON tender_history (version);


-- employee_tender
CREATE INDEX idx_employee_tender_employee_id ON employee_tender (employee_id);
CREATE INDEX idx_employee_tender_tender_id ON employee_tender (tender_id);


-- bid
CREATE INDEX idx_bid_tender_id ON bid (tender_id);
CREATE INDEX idx_bid_author_id ON bid (author_id);
CREATE INDEX idx_bid_status ON bid (status);


-- bid_history
CREATE INDEX idx_bid_history_tender_id ON bid_history (tender_id);
CREATE INDEX idx_bid_history_author_id ON bid_history (author_id);
CREATE INDEX idx_bid_history_version ON bid_history (version);


-- bid_review
CREATE INDEX idx_bid_review_bid_id ON bid_review (bid_id);
