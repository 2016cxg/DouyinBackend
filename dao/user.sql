DROP TABLE IF EXISTS `users` ;
CREATE TABLE `users` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `username` varchar(30) NOT NULL ,
                         `password` varchar(100) NOT NULL ,
                         `follow_count` int(10) NOT NULL DEFAULT 0 ,
                         `follower_count` int(10) NOT NULL DEFAULT 0,
                         `is_follow` boolean NOT NULL DEFAULT FALSE,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8