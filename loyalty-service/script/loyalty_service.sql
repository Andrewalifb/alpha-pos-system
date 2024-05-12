CREATE TABLE pos_loyalty_levels (
    level_id UUID PRIMARY KEY,
    level_name VARCHAR(50) NOT NULL,
    points_required INT NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_customer_loyalty (
    customer_id UUID REFERENCES pos_customers(customer_id),
    level_id UUID REFERENCES pos_loyalty_levels(level_id),
    reward_points INT DEFAULT 0,
    PRIMARY KEY (customer_id, level_id),
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_points_transactions (
    transaction_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES pos_customers(customer_id),
    points INT NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,  -- e.g., 'Earned', 'Redeemed'
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_rewards (
    reward_id UUID PRIMARY KEY,
    reward_name VARCHAR(255) NOT NULL,
    points_required INT NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_reward_redemptions (
    redemption_id UUID PRIMARY KEY,
    reward_id UUID REFERENCES pos_rewards(reward_id),
    customer_id UUID REFERENCES pos_customers(customer_id),
    redemption_date TIMESTAMP NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);


-- Memasukkan data ke dalam pos_loyalty_levels
INSERT INTO pos_loyalty_levels (level_id, level_name, points_required, company_id) VALUES
('aaa8400-e29b-41d4-a716-446655440000', 'Bronze', 1000, '550e8400-e29b-41d4-a716-446655440000'),
('aaa8400-e29b-41d4-a716-446655440001', 'Silver', 2000, '550e8400-e29b-41d4-a716-446655440000'),
('aaa8400-e29b-41d4-a716-446655440002', 'Gold', 3000, '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_customer_loyalty
INSERT INTO pos_customer_loyalty (customer_id, level_id, reward_points, company_id) VALUES
('111e8400-e29b-41d4-a716-446655440000', 'aaa8400-e29b-41d4-a716-446655440000', 1500, '550e8400-e29b-41d4-a716-446655440000'),
('111e8400-e29b-41d4-a716-446655440001', 'aaa8400-e29b-41d4-a716-446655440001', 2500, '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_points_transactions
INSERT INTO pos_points_transactions (transaction_id, customer_id, points, transaction_date, transaction_type, company_id) VALUES
('bbb8400-e29b-41d4-a716-446655440000', '111e8400-e29b-41d4-a716-446655440000', 500, '2024-05-01 09:30:00', 'Earned', '550e8400-e29b-41d4-a716-446655440000'),
('bbb8400-e29b-41d4-a716-446655440001', '111e8400-e29b-41d4-a716-446655440000', 1000, '2024-05-02 09:30:00', 'Redeemed', '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_rewards
INSERT INTO pos_rewards (reward_id, reward_name, points_required, company_id) VALUES
('ccc8400-e29b-41d4-a716-446655440000', 'Diskon 10%', 500, '550e8400-e29b-41d4-a716-446655440000'),
('ccc8400-e29b-41d4-a716-446655440001', 'Diskon 20%', 1000, '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_reward_redemptions
INSERT INTO pos_reward_redemptions (redemption_id, reward_id, customer_id, redemption_date, company_id) VALUES
('ddd8400-e29b-41d4-a716-446655440000', 'ccc8400-e29b-41d4-a716-446655440000', '111e8400-e29b-41d4-a716-446655440000', '2024-05-02 09:30:00', '550e8400-e29b-41d4-a716-446655440000');
