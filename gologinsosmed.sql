/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 5.7.24 : Database - gologinsosmed
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`gologinsosmed` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `gologinsosmed`;

/*Table structure for table `oauths` */

DROP TABLE IF EXISTS `oauths`;

CREATE TABLE `oauths` (
  `userid` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `provider` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `oauths` */

insert  into `oauths`(`userid`,`email`,`firstname`,`lastname`,`provider`) values 
('102115792585906930911','danilsyah94@gmail.com','danil','syah','google'),
('106350312954435142124','danilsyaharihardjo@gmail.com','Danil','Syah','google'),
('112653557453500171968','dnhdanil@gmail.com','danil','dnh','google'),
('5019655098047281','danilsyaharihardjo@gmail.com','Danil','Syah','facebook');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
