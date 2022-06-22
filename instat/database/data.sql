CREATE TABLE IF NOT EXISTS "newsCategories" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS `news` (
    `news_id` SERIAL PRIMARY KEY,
    `category_id` int(11) NOT NULL,
    `date_posted` date NOT NULL,
    `news_title` varchar(100) NOT NULL,
    `news_content` text NOT NULL,
    `date_updated` date NOT NULL,
    `news_status` int(1) NOT NULL,
    `comment_status` int(1) NOT NULL,
    `author_id` int(11) NOT NULL,
    PRIMARY KEY (`news_id`),
    FOREIGN KEY `category_id` REFERENCES (`category_id`,`author_id`),
    FOREIGN  KEY `author_id` REFERENCES (`author_id`)
);

CREATE TABLE "accounts" (
	"user_id" serial PRIMARY KEY,
	"username" VARCHAR ( 50 ) UNIQUE NOT NULL,
	"password" VARCHAR ( 50 ) NOT NULL,
	"email" VARCHAR ( 255 ) UNIQUE NOT NULL,
	"created_on" TIMESTAMP NOT NULL,
    "last_login" TIMESTAMP 
);

CREATE TABLE "roles"(
   "role_id" serial PRIMARY KEY,
   "role_name" VARCHAR (255) UNIQUE NOT NULL
);

CREATE TABLE "account_roles" (
    "user_id" INT NOT NULL,
    "role_id" INT NOT NULL,
    "grant_date" TIMESTAMP,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (role_id) REFERENCES roles (role_id),
    FOREIGN KEY (user_id) REFERENCES accounts (user_id)
);