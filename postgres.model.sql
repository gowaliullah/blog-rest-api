-- Create the ENUM type for role
CREATE TYPE user_role AS ENUM ('admin', 'author', 'reader');
DROP TYPE user_role; 

-- Users table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    user_role user_role DEFAULT 'reader',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


SELECT * FROM users;


-- Categories table
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    slug VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Tags table --> for tagging posts
CREATE TABLE tags (
    tag_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
DROP TABLE posts;

-- Posts table (core blog content)
CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    content TEXT NOT NULL,
    slug VARCHAR(200),
    author_id INTEGER NOT NULL,
    category_id INTEGER,
    published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(user_id),
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);



-- Comments table --> user comments on posts
CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Menu table
CREATE TABLE menu (
    menu_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



-- Social Icons table --> social media links, e.g., site-wide
CREATE TABLE social_icons (
    id SERIAL PRIMARY KEY,
    platform VARCHAR(50) NOT NULL, 
    url VARCHAR(255),
    icon_url VARCHAR(255),
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);


--- still below query and model not implement

-- Indexes for performance
CREATE INDEX idx_posts_slug ON posts(slug);
CREATE INDEX idx_posts_author ON posts(author_id);
CREATE INDEX idx_comments_post ON comments(post_id);




-- Junction table for many-to-many post-tags
CREATE TABLE post_tags (
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (post_id, tag_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);


-- Menu items (child of menu, linking to content)
CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    menu_id INTEGER NOT NULL,
    label VARCHAR(100) NOT NULL,
    url VARCHAR(255),  -- Internal/external URL
    post_id INTEGER,  -- Optional link to post
    category_id INTEGER,  -- Optional link to category
    order_index INTEGER DEFAULT 0,  -- For sorting
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (menu_id) REFERENCES menu(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE SET NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);





INSERT INTO users (username, email, password, user_role) VALUES
('johndoe', 'john.doe@example.com', 'hashedpassword123', 'admin'),
('janedoe', 'jane.doe@example.com', 'hashedpassword456', 'author'),
('peterpan', 'peter.pan@example.com', 'hashedpassword789', 'reader'),
('aliceblue', 'alice.blue@example.com', 'hashedpasswordabc', 'author'),
('bobthebuilder', 'bob.builder@example.com', 'hashedpassworddef', 'reader');