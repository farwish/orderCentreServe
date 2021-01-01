BEGIN;

CREATE TABLE IF NOT EXISTS orders (
    order_id char(32) NOT NULL comment '订单ID',

#   optional datatype for different system
#   account_id char(32) not null,
    account_id bigint unsigned NOT NULL comment '账号ID',

#   optional datatype for different system
#   product_id char(32) not null,
    product_id bigint unsigned NOT NULL comment '商品ID',

    product_name varchar(255) NOT NULL DEFAULT '' comment '商品名快照',

    trade_no varchar(32) NOT NULL DEFAULT '' comment '订单号',

    total_fee int unsigned NOT NULL DEFAULT 0 comment '金额(分)',

    payment_type tinyint unsigned NOT NULL DEFAULT 0 comment '支付类型(自定义)',
    payment_status tinyint unsigned NOT NULL DEFAULT 0 comment '支付状态(自定义)',
    payment_at datetime DEFAULT NULL comment '支付时间',

    status tinyint unsigned NOT NULL DEFAULT 0 comment '订单状态(自定义)',

    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL,

    PRIMARY KEY (`order_id`),
    UNIQUE KEY `orders_trade_no_unique` (`trade_no`),
    KEY `orders_product_id_index` (`product_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;