CREATE TABLE IF NOT EXISTS receipts (
	id uuid DEFAULT gen_random_uuid(),
	retailer text NOT NULL,
	purchaseDate date NOT NULL,
	purchaseTime time NOT NULL,
	total text NOT NULL,

	PRIMARY KEY (id)
);
