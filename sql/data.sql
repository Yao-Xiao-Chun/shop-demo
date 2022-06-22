CREATE DATABASE user;
USE user;

CREATE TABLE `user` (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                        `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                        `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
                        `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
                        `question` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码问题',
                        `answer` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码答案',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE DATABASE product;
USE product;

CREATE TABLE `product` (
                           `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id',
                           `cateid` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类别Id',
                           `name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
                           `subtitle` varchar(200) DEFAULT NULL DEFAULT '' COMMENT '商品副标题',
                           `images` text COMMENT '图片地址,json格式,扩展用',
                           `detail` text COMMENT '商品详情',
                           `price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '价格,单位-元保留两位小数',
                           `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存数量',
                           `status` int(6) NOT NULL DEFAULT 1 COMMENT '商品状态.1-在售 2-下架 3-删除',
                           `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`id`),
                           KEY `ix_cateid` (`cateid`),
                           KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';


CREATE TABLE `category` (
                            `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
                            `parentid` smallint(6) NOT NULL DEFAULT 0 COMMENT '父类别id当id=0时说明是根节点,一级类别',
                            `name` varchar(50) NOT NULL DEFAULT '' COMMENT '类别名称',
                            `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '类别状态1-正常,2-已废弃',
                            `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品类别表';


CREATE DATABASE cart;
USE cart;

CREATE TABLE `cart` (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '购物车id',
                        `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                        `proid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
                        `quantity` int(11) NOT NULL DEFAULT 0 COMMENT '数量',
                        `checked` int(11) NOT NULL DEFAULT 0 COMMENT '是否选择,1=已勾选,0=未勾选',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        KEY `ix_userid` (`userid`),
                        KEY `ix_proid` (`proid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';

CREATE DATABASE order;
USE order;

CREATE TABLE `orders` (
                          `id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                          `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                          `shoppingid` bigint(20) NOT NUMBER DEFAULT 0 COMMENT '收货信息表id',
                          `payment` decimal(20,2) DEFAULT NULL DEFAULT 0 COMMENT '实际付款金额,单位是元,保留两位小数',
                          `paymenttype` tinyint(4) NOT NULL DEFAULT 1 COMMENT '支付类型,1-在线支付',
                          `postage` int(10)  NOT NULL DEFAULT 0 COMMENT '运费,单位是元',
                          `status` smallint(6) NOT NULL DEFAULT 10 COMMENT '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭',
                          `payment_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付时间',
                          `send_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发货时间',
                          `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '交易完成时间',
                          `close_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '交易关闭时间',
                          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

CREATE TABLE `orderitem` (
                             `id` bigint(20) UNSIGNED NOT NULL COMMENT '订单子表id',
                             `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                             `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                             `proid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
                             `proname` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
                             `proimage` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片地址',
                             `currentunitprice` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '生成订单时的商品单价，单位是元,保留两位小数',
                             `quantity` int(10) NOT NULL DEFAULT 0 COMMENT '商品数量',
                             `totalprice` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '商品总价,单位是元,保留两位小数',
                             `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`id`),
                             KEY `ix_orderid` (`orderid`),
                             KEY `ix_userid` (`userid`),
                             KEY `ix_proid` (`proid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细表';

CREATE TABLE `shopping` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '收货信息表id',
                            `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                            `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                            `receiver_name` varchar(20) NOT NULL DEFAULT '' COMMENT '收货姓名',
                            `receiver_phone` varchar(20) NOT NULL DEFAULT '' COMMENT '收货固定电话',
                            `receiver_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '收货移动电话',
                            `receiver_province` varchar(20) NOT NULL DEFAULT '' COMMENT '省份',
                            `receiver_city` varchar(20) NOT NULL DEFAULT '' COMMENT '城市',
                            `receiver_district` varchar(20) NOT NULL DEFAULT '' COMMENT '区/县',
                            `receiver_address` varchar(200) NOT NULL DEFAULT '' COMMENT '详细地址',
                            `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`id`),
                            KEY `ix_orderid` (`orderid`),
                            KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收货信息表';


CREATE DATABASE pay;
USE pay;

CREATE TABLE `payinfo` (
                           `id` bigint(20) UNSIGNED NOT NULL COMMENT '支付信息表id',
                           `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                           `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                           `payplatform` tinyint(4) NOT NULL DEFAULT 0 COMMENT '支付平台:1-支付宝,2-微信',
                           `platformnumber` varchar(200) NOT NULL DEFAULT '' COMMENT '支付流水号',
                           `platformstatus` varchar(20) NOT NULL DEFAULT '' COMMENT '支付状态',
                           `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`id`),
                           KEY `ix_orderid` (`orderid`),
                           KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付信息表';