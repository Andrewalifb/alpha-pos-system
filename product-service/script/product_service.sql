CREATE TABLE pos_product_categories (
    category_id UUID PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_product_sub_categories (
    sub_category_id UUID PRIMARY KEY,
    sub_category_name VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES pos_product_categories(category_id),
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_products (
    product_id UUID PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    cost_price DECIMAL(10, 2),
    category_id UUID REFERENCES pos_product_categories(category_id),
    sub_category_id UUID REFERENCES pos_product_sub_categories(sub_category_id),
    stock_quantity INT NOT NULL,
    reorder_level INT,
    supplier_id UUID,
    product_description TEXT,
    active BOOLEAN DEFAULT TRUE,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_suppliers (
    supplier_id UUID PRIMARY KEY,
    supplier_name VARCHAR(255) NOT NULL,
    contact_name VARCHAR(255),
    contact_email VARCHAR(255),
    contact_phone VARCHAR(20),
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_inventory_history (
    inventory_id UUID PRIMARY KEY,
    product_id UUID REFERENCES pos_products(product_id),
    store_id UUID,
    date TIMESTAMP NOT NULL,
    quantity INT NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_promotions (
    promotion_id UUID PRIMARY KEY,
    product_id UUID REFERENCES pos_products(product_id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    discount_rate DECIMAL(3, 2) NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);


-- Memasukkan data ke dalam pos_product_categories
INSERT INTO pos_product_categories (category_id, category_name, company_id) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Bahan Makanan', '550e8400-e29b-41d4-a716-446655440000'),
('550e8400-e29b-41d4-a716-446655440001', 'Elektronik', '550e8400-e29b-41d4-a716-446655440000'),
('550e8400-e29b-41d4-a716-446655440002', 'Pakaian', '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_product_sub_categories
INSERT INTO pos_product_sub_categories (sub_category_id, sub_category_name, category_id, company_id) VALUES
('660e8400-e29b-41d4-a716-446655440000', 'Buah & Sayuran', '550e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000'),
('660e8400-e29b-41d4-a716-446655440001', 'Susu', '550e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000'),
('660e8400-e29b-41d4-a716-446655440002', 'Roti', '550e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_products
INSERT INTO pos_products (product_id, product_name, price, cost_price, category_id, sub_category_id, stock_quantity, reorder_level, supplier_id, product_description, active, company_id) VALUES
('770e8400-e29b-41d4-a716-446655440000', 'Apel Merah', 20000.00, 15000.00, '550e8400-e29b-41d4-a716-446655440000', '660e8400-e29b-41d4-a716-446655440000', 100, 20, NULL, 'Apel merah segar', TRUE, '550e8400-e29b-41d4-a716-446655440000'),
('770e8400-e29b-41d4-a716-446655440001', 'TV LED 32"', 3000000.00, 2500000.00, '550e8400-e29b-41d4-a716-446655440001', NULL, 50, 10, NULL, 'TV LED 32" dengan resolusi Full HD', TRUE, '550e8400-e29b-41d4-a716-446655440000'),
('770e8400-e29b-41d4-a716-446655440002', 'Kemeja Pria', 150000.00, 100000.00, '550e8400-e29b-41d4-a716-446655440002', NULL, 200, 50, NULL, 'Kemeja pria berbahan katun', TRUE, '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_suppliers
INSERT INTO pos_suppliers (supplier_id, supplier_name, contact_name, contact_email, contact_phone, company_id) VALUES
('880e8400-e29b-41d4-a716-446655440000', 'Pemasok Buah', 'Budi', 'budi@pemasokbuah.com', '081234567890', '550e8400-e29b-41d4-a716-446655440000'),
('880e8400-e29b-41d4-a716-446655440001', 'Pemasok Elektronik', 'Susi', 'susi@pemasokelektronik.com', '081234567891', '550e8400-e29b-41d4-a716-446655440000'),
('880e8400-e29b-41d4-a716-446655440002', 'Pemasok Pakaian', 'Rudi', 'rudi@pemasokpakaian.com', '081234567892', '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_inventory_history
INSERT INTO pos_inventory_history (inventory_id, product_id, store_id, date, quantity, company_id) VALUES
('990e8400-e29b-41d4-a716-446655440000', '770e8400-e29b-41d4-a716-446655440000', '330e8400-e29b-41d4-a716-446655440000', '2024-05-01 00:00:00', 100, '550e8400-e29b-41d4-a716-446655440000'),
('990e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', '330e8400-e29b-41d4-a716-446655440000', '2024-05-01 00:00:00', 50, '550e8400-e29b-41d4-a716-446655440000'),
('990e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440002', '330e8400-e29b-41d4-a716-446655440000', '2024-05-01 00:00:00', 200, '550e8400-e29b-41d4-a716-446655440000');

-- Memasukkan data ke dalam pos_promotions
INSERT INTO pos_promotions (promotion_id, product_id, start_date, end_date, discount_rate, company_id) VALUES
('aa0e8400-e29b-41d4-a716-446655440000', '770e8400-e29b-41d4-a716-446655440000', '2024-05-01', '2024-05-07', 0.10, '550e8400-e29b-41d4-a716-446655440000'),
('aa0e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', '2024-05-01', '2024-05-07', 0.15, '550e8400-e29b-41d4-a716-446655440000'),
('aa0e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440002', '2024-05-01', '2024-05-07', 0.20, '550e8400-e29b-41d4-a716-446655440000');
