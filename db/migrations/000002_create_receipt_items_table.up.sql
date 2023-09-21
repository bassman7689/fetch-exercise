CREATE TABLE IF NOT EXISTS receipt_items (
	id uuid DEFAULT gen_random_uuid(),
	shortDescription text NOT NULL,
	price text NOT NULL,

	PRIMARY KEY (id)
);
