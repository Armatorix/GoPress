-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users(
    id SERIAL PRIMARY KEY NOT NULL,
    email text UNIQUE NOT NULL,
    nick text NOT NULL DEFAULT '',
    avatar_url text NOT NULL DEFAULT '',
    password text NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS public.articles(
    id SERIAL PRIMARY KEY NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    body text NOT NULL,
    released boolean NOT NULL DEFAULT false,
    author_id int NOT NULL REFERENCES public.users(id),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    FOREIGN KEY (author_id) REFERENCES public.users(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd