CREATE TABLE "public"."clientes" (
    "ativo" boolean NOT NULL,
    "id" uuid NOT NULL,
    "cpf" character varying(255),
    "email" character varying(255),
    "nome" character varying(255),
    CONSTRAINT "clientes_pkey" PRIMARY KEY ("id")
) WITH (oids = false);