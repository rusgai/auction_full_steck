-- Create enum type
CREATE TYPE status AS ENUM('Live', 'Finished', 'ReserveNotMed');

-- Create tables
CREATE TABLE
	"auction" (
		"id" varchar(32) NOT NULL PRIMARY KEY,
		"reserve_price" int,
		"seller" varchar(32),
		"winner" varchar(32),
		"sold_amount" int,
		"current_high_bid" int,
		"created_at" timestamp,
		"updated_at" timestamp,
		"auction_end" timestamp,
		"status" status
	);

CREATE TABLE
	"item" (
		"id" varchar(32) NOT NULL PRIMARY KEY,
		"make" varchar(50),
		"model" varchar(50),
		"year" int,
		"color" varchar(50),
		"mileage" int,
		"image_url" varchar(200),
		"auction_id" varchar(32) NOT NULL
	);

-- Alters tables
ALTER TABLE "item"
ADD CONSTRAINT "fk_item_auction" FOREIGN KEY ("auction_id") REFERENCES "auction" ("id") ON DELETE CASCADE;
-- insert auction table
INSERT INTO auction VALUES
('155225c1-4448-4066-9886-6786536e05ea',50000,'tom',NULL,'NULL,NULL,'2025-02-26 20:19:30.745501+00','2025-02-26 20:19:30.745501+00','2025-02-16 20:19:30.745501+00',3),
('3659ac24-29dd-407a-81f5-ecfe6f924b9b',20000,'bob',NULL,'NULL,NULL,'2025-02-26 20:19:30.745504+00','2025-02-26 20:19:30.745504+00','2025-04-15 20:19:30.745504+00',1),
('40490065-dac7-46b6-acc4-df507e0d6570',20000,'tom',NULL,'NULL,NULL,'2025-02-26 20:19:30.745504+00','2025-02-26 20:19:30.745504+00','2025-03-18 20:19:30.745504+00',1),
('466e4744-4dc5-4987-aae0-b621acfc5e39',20000,'alice',NULL,'NULL,NULL,'2025-02-26 20:19:30.745501+00','2025-02-26 20:19:30.745501+00','2025-03-28 20:19:30.745501+00',1),
('47111973-d176-4feb-848d-0ea22641c31a',150000,'alice',NULL,'NULL,NULL,'2025-02-26 20:19:30.745503+00','2025-02-26 20:19:30.745503+00','2025-03-11 20:19:30.745503+00',1),
('6a5011a1-fe1f-47df-9a32-b5346b289391',0,'bob',NULL,'NULL,NULL,'2025-02-26 20:19:30.745503+00','2025-02-26 20:19:30.745503+00','2025-03-17 20:19:30.745504+00',1),
('afbee524-5972-4075-8800-7d1f9d7b0a0c',20000,'bob',NULL,'NULL,NULL,'2025-02-26 20:19:30.744816+00','2025-02-26 20:19:30.744816+00','2025-03-08 20:19:30.745032+00',1),
('bbab4d5a-8565-48b1-9450-5ac2a5c4a654',0,'bob',NULL,'NULL,NULL,'2025-02-26 20:19:30.7455+00  ','2025-02-26 20:19:30.7455+00  ','2025-03-02 20:19:30.7455+00  ',1),
('c8c3ec17-01bf-49db-82aa-1ef80b833a9f',90000,'alice',NULL,'NULL,NULL,'2025-02-26 20:19:30.745497+00','2025-02-26 20:19:30.745498+00','2025-04-27 20:19:30.745499+00',1),
('dc1e4071-d19d-459b-b848-b5c3cd3d151f',20000,'bob',NULL,'NULL,NULL,'2025-02-26 20:19:30.745502+00','2025-02-26 20:19:30.745502+00','2025-04-12 20:19:30.745502+00',1);
-- insert item table
INSERT INTO auction VALUES
('2cae3c40-e30d-473d-bbfd-7e3e12e5d407','Audi','TT','2020','Black',25400,'https://cdn.pixabay.com/photo/2016/09/01/15/06/audi-1636320_960_720.jpg','40490065-dac7-46b6-acc4-df507e0d6570'),
('49a1f275-e5a5-4bc7-a37e-4441ca14acec','Ferrari','F-430  ','2022','Red',  5000,'https://cdn.pixabay.com/photo/2017/11/08/14/39/ferrari-f430-2930661_960_720.jpg','47111973-d176-4feb-848d-0ea22641c31a'),
('4dc2812c-34e3-474c-badc-6a1fd546a9a7','Bugatti','Veyron ','2018','Black', 15035,'https://cdn.pixabay.com/photo/2012/05/29/00/43/car-49278_960_720.jpg','c8c3ec17-01bf-49db-82aa-1ef80b833a9f'),
('63f043a9-6b9b-4224-bf7a-bd0cbfef89e8','Ford','Mustang','2023','Black', 65125,'https://cdn.pixabay.com/photo/2012/11/02/13/02/car-63930_960_720.jpg','bbab4d5a-8565-48b1-9450-5ac2a5c4a654'),
('6baf26e8-98da-4521-bb11-e229f3230247','Ford','GT','2020','White', 50000,'https://cdn.pixabay.com/photo/2016/05/06/16/32/car-1376190_960_720.jpg','afbee524-5972-4075-8800-7d1f9d7b0a0c'),
('73b23fb4-9f0c-4e65-9ef3-5eb26880f5d3','Ferrari','Spider','2015','Red', 50000,'https://cdn.pixabay.com/photo/2017/11/09/01/49/ferrari-458-spider-2932191_960_720.jpg','dc1e4071-d19d-459b-b848-b5c3cd3d151f'),
('760e71fc-30b4-4485-85f6-b723ec1f7603','Mercedes','SLK','2020','Silver', 15001,'https://cdn.pixabay.com/photo/2016/04/17/22/10/mercedes-benz-1335674_960_720.png','155225c1-4448-4066-9886-6786536e05ea'),
('a526cdb0-8128-433f-8897-965dbce64dbf','BMW','X1','2017','White', 90000,'https://cdn.pixabay.com/photo/2017/08/31/05/47/bmw-2699538_960_720.jpg','466e4744-4dc5-4987-aae0-b621acfc5e39'),
('a652baf7-61e5-4a66-982e-2bba41fa6e87','Audi','R8','2021','White', 10050,'https://cdn.pixabay.com/photo/2019/12/26/20/50/audi-r8-4721217_960_720.jpg','6a5011a1-fe1f-47df-9a32-b5346b289391'),
('ad032ffd-8967-4b4a-aa17-2277b1686c70','Ford','Model T','1938','Rust',150150,'https://cdn.pixabay.com/photo/2017/08/02/19/47/vintage-2573090_960_720.jpg','3659ac24-29dd-407a-81f5-ecfe6f924b9b');



