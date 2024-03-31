CREATE TABLE "user"
(
    id      SERIAL PRIMARY KEY,
    user_id VARCHAR(50) UNIQUE NOT NULL,
    user_pw VARCHAR(50)        NOT NULL
);

-- image 테이블 생성
CREATE TABLE image
(
    id                SERIAL PRIMARY KEY,
    user_id           VARCHAR(50) REFERENCES "user" (user_id),
    image_server_path VARCHAR(200) NOT NULL
);