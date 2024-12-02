CREATE TABLE "public"."clientes" (
    "ativo" boolean NOT NULL,
    "id" uuid NOT NULL,
    "cpf" character varying(255),
    "email" character varying(255),
    "nome" character varying(255),
    CONSTRAINT "clientes_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "clientes" ("ativo", "id", "cpf", "email", "nome") VALUES
('t',	'63a59178-39f8-4a28-a2c7-989a57ca7b54',	'12312312312',	'filipe@email.com',	'FILIPE ANDRADE'),
('f',	'5793fc61-8d22-4183-9b20-079e624074a3',	'78978978978',	'murilo@email.com',	'MURILO MARTINS'),
('t',	'b57b4dcc-c47f-40f0-8331-6185bb9b3568',	'45645645645',	'joao@email.com',	'JOAO MARCOS'),
('t',	'b57b4dcc-c47f-40f0-8331-6185bb343443',	'35645645644',	'caio@email.com',	'CAIO MATOS');