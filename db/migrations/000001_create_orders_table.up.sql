BEGIN;

CREATE TABLE IF NOT EXISTS orders (
#   Generate by self uuid
    order_id char(36) NOT NULL comment '订单ID',

#   Changeable datatype for different system, change together with controller's validator if needed.
#   account_id char(36) not null,
    account_id bigint unsigned NOT NULL comment '账号ID',

#   Changeable datatype for different system, change together with controller's validator if needed.
#   product_id char(36) not null,
    product_id bigint unsigned NOT NULL comment '商品ID',

    trade_no varchar(36) NOT NULL comment '订单号',

    total_fee int unsigned NOT NULL comment '金额(分)',

    status tinyint unsigned NOT NULL comment '订单状态(自定义)',

    payment_type tinyint unsigned NOT NULL comment '支付类型(自定义)',
    payment_status tinyint unsigned NOT NULL comment '支付状态(自定义)',
    payment_at datetime DEFAULT NULL comment '支付时间',

#   optional value
    product_name varchar(255) NOT NULL DEFAULT '' comment '商品名快照',

    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime DEFAULT NULL,

    PRIMARY KEY (`order_id`),
    UNIQUE KEY `orders_trade_no_unique` (`trade_no`),
    KEY `orders_product_id_index` (`product_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;