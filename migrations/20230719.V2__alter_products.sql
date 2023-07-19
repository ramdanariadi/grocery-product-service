ALTER TABLE products ADD shop_id VARCHAR(36);
ALTER TABLE products ADD FOREIGN KEY (shop_id) REFERENCES shops(id);