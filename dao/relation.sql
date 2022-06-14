DROP TABLE IF EXISTS `relation` ;
CREATE TABLE `relation` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                         `fuid` int(10) NOT NULL COMMENT 'from user',
                         `tuid` int(10) NOT NULL COMMENT 'to user',
                         `relation` int(10) NOT NULL DEFAULT 0 COMMENT 'relation, 0 for none follow, 1 for fuid follow tuid, 2 for tuid follow fuid, 3 for mutual follow',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8