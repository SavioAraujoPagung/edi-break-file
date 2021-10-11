create table procedas (
	id int not null,
	File_Name varchar(255),
	Head_File_ID int,
	Head_File_Two_ID int,
	Carrier_ID int,
	primary key(id)
);

create table HeadFiles (
	id int not null,
	Head_File_Record_Identifier int,
	Sender_Name varchar(255),
	Recipient_Name varchar(255),
	Created_At date,
	primary key(id)
);

ALTER TABLE procedas ADD FOREIGN KEY (Head_File_ID) REFERENCES HeadFiles(id);

create table HeadFileTwos (
	id int not null,
	Head_File_Two_Record_Identifier int,
	File_Identifier varchar(255),
	Filler_Head_File_Two varchar(255),
	primary key(id)
);

ALTER TABLE procedas ADD FOREIGN KEY (Head_File_Two_ID) REFERENCES HeadFileTwos(id);

create table Carriers (
	id int not null,
	Carrier_Record_Identifier int,
	Registered_Number_Carrier varchar(255),
	Carrier_Name varchar(255),
	Filler_Carrier varchar(50),
	Transport_Knowledges_ID int
	primary key(id)
);

ALTER TABLE procedas ADD FOREIGN KEY (Carrier_ID) REFERENCES Carriers(id);

create table TransportKnowledges (
	id int not null,
	Transport_Knowledge_Record_Identifier int,
	Registered_Number_Cte varchar(255),
	Contracting_Carrier varchar(255),
	cte_series int,
	cte_number int,
	primary key(id)
);

ALTER TABLE Carriers ADD FOREIGN KEY (Transport_Knowledges_ID) REFERENCES TransportKnowledges(id);

create table Occurrences (
	id int not null,
	Transport_Knowledges_ID int,
	Occurrence_Code_ID int,
	Invoice_ID int,
	Occurrence_Record_Identifier int,
	Occurrence_Date date,
	Observation_Code int,
	text_Occurrence varchar(255),
	filler_Occurrence varchar(255),
	primary key(id)
);

ALTER TABLE Occurrences ADD FOREIGN KEY (Transport_Knowledges_ID) REFERENCES TransportKnowledges(id);

create table Invoices (
	id int not null,
	Registered_Number_Invoice int,
	nfe_series int,
	nfe_number int,
	primary key(id)
);

ALTER TABLE Occurrences ADD FOREIGN KEY (Invoice_ID) REFERENCES Invoices(id);

create table OccurrenceCodes (
	id int not null,
	Code int,
	Description varchar(255),
	primary key(id)
);

ALTER TABLE Occurrences ADD FOREIGN KEY (Occurrence_Code_ID) REFERENCES OccurrenceCodes(id);