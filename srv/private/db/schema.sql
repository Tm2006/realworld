-- Table for users
CREATE TABLE users (
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    email    TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    bio      TEXT,
    image    TEXT
);

-- Table for articles
CREATE TABLE articles (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    slug            TEXT      NOT NULL UNIQUE,
    title           TEXT      NOT NULL,
    description     TEXT      NOT NULL,
    body            TEXT      NOT NULL,
    created_at      DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    favorites_count INTEGER   NOT NULL DEFAULT 0,
    author_id       INTEGER   REFERENCES users (id) ON DELETE SET NULL
);

-- Table for tags
CREATE TABLE tags (
    id  INTEGER PRIMARY KEY AUTOINCREMENT,
    tag TEXT NOT NULL UNIQUE
);

-- Table for article_tags (many-to-many relationship)
CREATE TABLE article_tags (
    article_id INTEGER REFERENCES articles (id) ON DELETE CASCADE,
    tag_id     INTEGER REFERENCES tags (id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, tag_id)
);

-- Table for favorites
CREATE TABLE favorites (
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    article_id INTEGER REFERENCES articles (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, article_id)
);

-- Table for comments
CREATE TABLE comments (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    article_id INTEGER REFERENCES articles (id) ON DELETE CASCADE,
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    body       TEXT      NOT NULL,
    created_at DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE follows (
    follower_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    followee_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (follower_id, followee_id)
);

-- Indexes for performance
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_articles_author_id ON articles (author_id);
CREATE INDEX idx_comments_article_id ON comments (article_id);
CREATE INDEX idx_comments_user_id ON comments (user_id);