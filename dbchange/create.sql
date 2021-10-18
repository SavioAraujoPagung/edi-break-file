-- public.headfiles definition

-- Drop table

-- DROP TABLE public.headfiles;

CREATE TABLE public.headfiles (
	id int4 NOT NULL,
	head_file_record_identifier int4 NULL,
	sender_name varchar(255) NULL,
	recipient_name varchar(255) NULL,
	created_at date NULL,
	CONSTRAINT headfiles_pkey PRIMARY KEY (id)
);


-- public.headfiletwos definition

-- Drop table

-- DROP TABLE public.headfiletwos;

CREATE TABLE public.headfiletwos (
	id int4 NOT NULL,
	head_file_two_record_identifier int4 NULL,
	file_identifier varchar(255) NULL,
	filler_head_file_two varchar(255) NULL,
	CONSTRAINT headfiletwos_pkey PRIMARY KEY (id)
);


-- public.invoices definition

-- Drop table

-- DROP TABLE public.invoices;

CREATE TABLE public.invoices (
	id int4 NOT NULL,
	registered_number_invoice int4 NULL,
	nfe_series int4 NULL,
	nfe_number int4 NULL,
	CONSTRAINT invoices_pkey PRIMARY KEY (id)
);


-- public.occurrencecodes definition

-- Drop table

-- DROP TABLE public.occurrencecodes;

CREATE TABLE public.occurrencecodes (
	id int4 NOT NULL,
	code int4 NULL,
	description varchar(255) NULL,
	CONSTRAINT occurrencecodes_pkey PRIMARY KEY (id)
);


-- public.transportknowledges definition

-- Drop table

-- DROP TABLE public.transportknowledges;

CREATE TABLE public.transportknowledges (
	id int4 NOT NULL,
	transport_knowledge_record_identifier int4 NULL,
	registered_number_cte int4 NULL,
	contracting_carrier varchar(255) NULL,
	cte_series int4 NULL,
	cte_number int4 NULL,
	CONSTRAINT transportknowledges_pkey PRIMARY KEY (id)
);


-- public.carriers definition

-- Drop table

-- DROP TABLE public.carriers;

CREATE TABLE public.carriers (
	id int4 NOT NULL,
	carrier_record_identifier int4 NULL,
	registered_number_carrier varchar(255) NULL,
	carrier_name varchar(255) NULL,
	filler_carrier varchar(50) NULL,
	transport_knowledges_id int4 NULL,
	CONSTRAINT carriers_pkey PRIMARY KEY (id),
	CONSTRAINT carriers_transport_knowledges_id_fkey FOREIGN KEY (transport_knowledges_id) REFERENCES public.transportknowledges(id)
);


-- public.occurrences definition

-- Drop table

-- DROP TABLE public.occurrences;

CREATE TABLE public.occurrences (
	id int4 NOT NULL,
	transport_knowledges_id int4 NULL,
	occurrence_code_id int4 NULL,
	invoice_id int4 NULL,
	occurrence_record_identifier int4 NULL,
	occurrence_date date NULL,
	observation_code int4 NULL,
	text_occurrence varchar(255) NULL,
	filler_occurrence varchar(255) NULL,
	CONSTRAINT occurrences_pkey PRIMARY KEY (id),
	CONSTRAINT occurrences_invoice_id_fkey FOREIGN KEY (invoice_id) REFERENCES public.invoices(id),
	CONSTRAINT occurrences_occurrence_code_id_fkey FOREIGN KEY (occurrence_code_id) REFERENCES public.occurrencecodes(id),
	CONSTRAINT occurrences_transport_knowledges_id_fkey FOREIGN KEY (transport_knowledges_id) REFERENCES public.transportknowledges(id)
);


-- public.procedas definition

-- Drop table

-- DROP TABLE public.procedas;

CREATE TABLE public.procedas (
	id int4 NOT NULL,
	file_name varchar(255) NULL,
	head_file_id int4 NULL,
	head_file_two_id int4 NULL,
	carrier_id int4 NULL,
	CONSTRAINT procedas_pkey PRIMARY KEY (id),
	CONSTRAINT "Procedas_carrier_id_fkey" FOREIGN KEY (carrier_id) REFERENCES public.carriers(id),
	CONSTRAINT "Procedas_head_file_id_fkey" FOREIGN KEY (head_file_id) REFERENCES public.headfiles(id),
	CONSTRAINT "Procedas_head_file_two_id_fkey" FOREIGN KEY (head_file_two_id) REFERENCES public.headfiletwos(id)
);