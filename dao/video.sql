DROP TABLE IF EXISTS `videos` ;
CREATE TABLE `videos` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `userid` int(10) NOT NULL ,
                         `title` varchar(100) NOT NULL ,
                         `favorite_count` int(10) NOT NULL DEFAULT 0 ,
                         `comment_count` int(10) NOT NULL DEFAULT 0,
                         `is_favorite` boolean NOT NULL DEFAULT FALSE,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8