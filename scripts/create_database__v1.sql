CREATE TYPE GENDER AS ENUM ('Male', 'Female', 'Other');

CREATE TABLE survivors (
    id SERIAL,
    "name" VARCHAR(256),
    age INTEGER,
    gender GENDER,
    location_latitude NUMERIC(12, 10),
    location_longitude NUMERIC(12, 10),
    location_timezone VARCHAR(80),
    infected BOOLEAN,
    deceased BOOLEAN,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    
    PRIMARY KEY (id)
);

CREATE TABLE credentials (
    id SERIAL,
    survivor_id INTEGER,
    username VARCHAR(80),
    "password" VARCHAR(256),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_credential_survivor 
        FOREIGN KEY (survivor_id) 
        REFERENCES survivors (id) 
            ON UPDATE CASCADE
);

CREATE TABLE inventories (
    id SERIAL,
    owner_id INTEGER,
    "disabled" BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_inventory_owner
        FOREIGN KEY (owner_id) 
        REFERENCES survivors (id) 
            ON UPDATE CASCADE
);

CREATE TYPE ITEM_RARITY AS ENUM ('Common', 'Uncommon', 'Rare', 'Epic');

CREATE TABLE items (
    id SERIAL,
    "name" VARCHAR(256),
    icon VARCHAR(256),
    price INTEGER,
    rarity ITEM_RARITY,
    
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE resources (
    id SERIAL,
    inventory_id INTEGER,
    item_id INTEGER,
    quantity INTEGER NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_resource_inventory 
        FOREIGN KEY (inventory_id) 
        REFERENCES inventories (id) 
            ON UPDATE CASCADE,
    CONSTRAINT fk_resource_item
        FOREIGN KEY (item_id) 
        REFERENCES items (id) 
            ON UPDATE CASCADE
);

CREATE TABLE groups (
    id SERIAL,
    "name" VARCHAR(80),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE group_members (
    id SERIAL,
    survivor_id INTEGER,
    group_id INTEGER,
    "role" VARCHAR(80),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_group_member_member
        FOREIGN KEY (survivor_id) 
        REFERENCES survivors (id) 
            ON UPDATE CASCADE,
    CONSTRAINT fk_group_member_group 
        FOREIGN KEY (group_id) 
        REFERENCES groups (id) 
            ON UPDATE CASCADE
);

CREATE TABLE infection_reports (
    id SERIAL,
    reported_id INTEGER,
    reportee_id INTEGER,
    annotation TEXT,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_infection_reported 
        FOREIGN KEY (reported_id) 
        REFERENCES survivors (id) 
            ON UPDATE CASCADE,
    CONSTRAINT fk_infection_reportee 
        FOREIGN KEY (reportee_id) 
        REFERENCES survivors (id) 
            ON UPDATE CASCADE
);

CREATE TABLE trade_inventories (
    id SERIAL,
    survivor_id INTEGER,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_trade_inventory_survivor 
        FOREIGN KEY (survivor_id) 
        REFERENCES survivors (id) 
            ON UPDATE CASCADE
);

CREATE TABLE trade_resources (
    id SERIAL,
    item_id INTEGER,
    inventory_id INTEGER,
    quantity INTEGER NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_trade_resource_item 
        FOREIGN KEY (item_id) 
        REFERENCES items (id) 
            ON UPDATE CASCADE,
    CONSTRAINT fk_trade_resource_inventory 
        FOREIGN KEY (inventory_id) 
        REFERENCES trade_inventories (id) 
            ON UPDATE CASCADE
);

CREATE TYPE TRADE_STATUS AS ENUM ('Open', 'Accepted', 'Rejected'); 

CREATE TABLE trades (
    id SERIAL,
    sender_id INTEGER,
    receiver_id INTEGER,
    "status" TRADE_STATUS,
    annotation TEXT,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id),
    CONSTRAINT fk_trade_sender 
        FOREIGN KEY (sender_id) 
        REFERENCES trade_inventories (id) 
            ON UPDATE CASCADE,
    CONSTRAINT fk_trade_receiver 
        FOREIGN KEY (receiver_id) 
        REFERENCES trade_inventories (id) 
            ON UPDATE CASCADE
);
