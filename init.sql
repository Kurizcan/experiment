-- MySQL dump 10.13  Distrib 8.0.13, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: db_experiment
-- ------------------------------------------------------
-- Server version	8.0.13

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `answer`
--

DROP TABLE IF EXISTS `answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `answer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupId` int(11) NOT NULL,
  `studentId` int(64) NOT NULL,
  `problemId` int(11) NOT NULL,
  `status` int(10) DEFAULT '0',
  `score` int(10) DEFAULT '0',
  `submit` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `error` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `correct` tinyint(1) DEFAULT '0',
  `updateTime` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `groupId` (`groupId`),
  KEY `problemId` (`problemId`),
  KEY `union_index` (`studentId`,`groupId`,`problemId`),
  CONSTRAINT `answer_ibfk_1` FOREIGN KEY (`groupId`) REFERENCES `experiment` (`groupid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `answer_ibfk_2` FOREIGN KEY (`studentId`) REFERENCES `user` (`userid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `answer_ibfk_3` FOREIGN KEY (`problemId`) REFERENCES `problem` (`problemid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `class`
--

DROP TABLE IF EXISTS `class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `class` (
  `classId` varchar(64) NOT NULL,
  `grade` int(64) NOT NULL,
  `class` int(64) NOT NULL,
  `college` varchar(10) NOT NULL,
  `major` varchar(20) DEFAULT NULL,
  `number` int(11) DEFAULT NULL,
  PRIMARY KEY (`classId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `class`
--

LOCK TABLES `class` WRITE;
/*!40000 ALTER TABLE `class` DISABLE KEYS */;
INSERT INTO `class` VALUES ('gsnnZTjZg',2016,4,'信息科学技术学院','软件工程',34),('nd77WTCWR',2016,2,'信息科学技术学院','软件工程',32),('nZukWoCWR',2016,3,'信息科学技术学院','软件工程',31),('Ox7nZoCWg',2016,1,'信息科学技术学院','软件工程',30);
/*!40000 ALTER TABLE `class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `class_experiment`
--

DROP TABLE IF EXISTS `class_experiment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `class_experiment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupId` int(11) NOT NULL,
  `classId` varchar(64) NOT NULL,
  `groupName` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `groupId` (`groupId`),
  KEY `union_index` (`classId`,`groupId`),
  CONSTRAINT `class_experiment_ibfk_1` FOREIGN KEY (`groupId`) REFERENCES `experiment` (`groupid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `class_experiment_ibfk_2` FOREIGN KEY (`classId`) REFERENCES `class` (`classid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `experiment`
--

DROP TABLE IF EXISTS `experiment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `experiment` (
  `groupId` int(64) NOT NULL AUTO_INCREMENT,
  `groupName` varchar(20) NOT NULL,
  `poster` varchar(20) NOT NULL,
  `problems` json NOT NULL,
  PRIMARY KEY (`groupId`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `problem`
--

DROP TABLE IF EXISTS `problem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `problem` (
  `problemId` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(255) NOT NULL,
  `example` json NOT NULL,
  `data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `solution` varchar(255) NOT NULL,
  `output` json NOT NULL,
  `poster` varchar(20) NOT NULL,
  PRIMARY KEY (`problemId`)
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `student_class`
--

DROP TABLE IF EXISTS `student_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `student_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `studentId` int(64) NOT NULL,
  `classId` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `studentId` (`studentId`),
  KEY `union_index` (`classId`,`studentId`),
  CONSTRAINT `student_class_ibfk_1` FOREIGN KEY (`studentId`) REFERENCES `user` (`userid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `student_class_ibfk_2` FOREIGN KEY (`classId`) REFERENCES `class` (`classid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student_class`
--

LOCK TABLES `student_class` WRITE;
/*!40000 ALTER TABLE `student_class` DISABLE KEYS */;
INSERT INTO `student_class` VALUES (105,113,'nd77WTCWR'),(106,114,'nd77WTCWR'),(100,107,'nZukWoCWR'),(101,109,'nZukWoCWR'),(102,110,'nZukWoCWR'),(103,111,'nZukWoCWR'),(104,112,'nZukWoCWR');
/*!40000 ALTER TABLE `student_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student_experiment`
--

DROP TABLE IF EXISTS `student_experiment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `student_experiment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupId` int(11) NOT NULL,
  `studentId` int(64) NOT NULL,
  `score` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `groupId` (`groupId`),
  KEY `union_index` (`studentId`,`groupId`),
  CONSTRAINT `student_experiment_ibfk_1` FOREIGN KEY (`groupId`) REFERENCES `experiment` (`groupid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `student_experiment_ibfk_2` FOREIGN KEY (`studentId`) REFERENCES `user` (`userid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=138 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `teacher_class`
--

DROP TABLE IF EXISTS `teacher_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `teacher_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `teacherId` int(64) NOT NULL,
  `classId` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `classId` (`classId`),
  KEY `union_index` (`teacherId`,`classId`),
  CONSTRAINT `teacher_class_ibfk_1` FOREIGN KEY (`teacherId`) REFERENCES `user` (`userid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `teacher_class_ibfk_2` FOREIGN KEY (`classId`) REFERENCES `class` (`classid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teacher_class`
--

LOCK TABLES `teacher_class` WRITE;
/*!40000 ALTER TABLE `teacher_class` DISABLE KEYS */;
INSERT INTO `teacher_class` VALUES (100,108,'gsnnZTjZg'),(101,108,'nd77WTCWR'),(103,108,'nZukWoCWR'),(102,108,'Ox7nZoCWg');
/*!40000 ALTER TABLE `teacher_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `userId` int(64) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `number` varchar(255) NOT NULL,
  `type` int(3) NOT NULL,
  `grade` int(64) DEFAULT NULL,
  `class` int(64) DEFAULT NULL,
  `major` varchar(20) DEFAULT NULL,
  `college` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`userId`),
  UNIQUE KEY `number` (`number`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (107,'乃万','$2a$10$KXesP/282TD.GqcnAhzmXunf1cstkLwSP4wr1PW9r1FZ0sKOowO06','2220162022',0,2016,3,'软件工程','信息科学技术学院'),(108,'chiyo','$2a$10$s.IsTeo9tLJJRjGvznMm6eWm9EhamvFMBC0h9qRCQ2JnnowZgZgF6','1110162020',1,0,0,'','信息科学技术学院'),(109,'高健','$2a$10$yonmyjYkhYfqcI0WJ2OpROONq4qOk56v/rQCWXjCn.dLUMsYS.kf2','2220162025',1,0,3,'软件工程','信息科学技术学院'),(110,'武磊','$2a$10$.ZNYurdbamPgF2l4bhO1IexDYSB0C5cC4psoXMT8V8OexeyILEDjW','2220162026',1,0,3,'软件工程','信息科学技术学院'),(111,'杨幂','$2a$10$mk2yw6pXJGgilkpBulQYr.Nly0lSPmEIuXunYvvwsCxCKp8pBTDWq','2220162027',1,0,3,'软件工程','信息科学技术学院'),(112,'江疏影','$2a$10$UwoUhp9u5U5Tz79Eyg5IAuCZx4yHRH9VJpyAKkr.ISjB75gxhPWG2','2220162028',1,0,3,'软件工程','信息科学技术学院'),(113,'苞米','$2a$10$3mN0tAfmdJ1Vn7BdQi3llOcsemtp8Qv8fbKCahzjijjj8fhKQeDVC','2220162029',1,0,2,'软件工程','信息科学技术学院'),(114,'张景森','$2a$10$QoOhnRJST4bGehS61iXZveUL4..6P6/OjJHMx/UMKs97S/nu7OVZu','2220162030',1,0,2,'软件工程','信息科学技术学院');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-04 22:08:35