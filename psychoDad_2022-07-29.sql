# ************************************************************
# Sequel Ace SQL dump
# 版本號： 20033
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主機: localhost (MySQL 8.0.29)
# 數據庫: psychoDad
# 生成時間: 2022-07-28 17:22:55 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 轉儲表 counties
# ------------------------------------------------------------

CREATE TABLE `counties` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `counties` WRITE;
/*!40000 ALTER TABLE `counties` DISABLE KEYS */;

INSERT INTO `counties` (`id`, `name`)
VALUES
	(1,'基隆市'),
	(2,'台北市'),
	(3,'新北市'),
	(4,'桃園市'),
	(5,'新竹市'),
	(6,'新竹縣'),
	(7,'苗栗縣'),
	(8,'台中市'),
	(9,'彰化縣'),
	(10,'南投縣'),
	(11,'雲林縣'),
	(12,'嘉義市'),
	(13,'嘉義縣'),
	(14,'台南市'),
	(15,'高雄市'),
	(17,'屏東縣'),
	(18,'花蓮縣'),
	(19,'宜蘭縣'),
	(20,'澎湖縣'),
	(21,'金門縣'),
	(22,'連江縣'),
	(23,'台東縣');

/*!40000 ALTER TABLE `counties` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 districts
# ------------------------------------------------------------

CREATE TABLE `districts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `county_id` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `districts` WRITE;
/*!40000 ALTER TABLE `districts` DISABLE KEYS */;

INSERT INTO `districts` (`id`, `name`, `county_id`)
VALUES
	(372,'中正區','2'),
	(373,'大同區','2'),
	(374,'中山區','2'),
	(375,'松山區','2'),
	(376,'大安區','2'),
	(377,'萬華區','2'),
	(378,'信義區','2'),
	(379,'士林區','2'),
	(380,'北投區','2'),
	(381,'內湖區','2'),
	(382,'南港區','2'),
	(383,'文山區','2'),
	(384,'仁愛區','1'),
	(385,'信義區','1'),
	(386,'中正區','1'),
	(387,'中山區','1'),
	(388,'安樂區','1'),
	(389,'暖暖區','1'),
	(390,'七堵區','1'),
	(391,'萬里區','3'),
	(392,'金山區','3'),
	(393,'板橋區','3'),
	(394,'汐止區','3'),
	(395,'深坑區','3'),
	(396,'石碇區','3'),
	(397,'瑞芳區','3'),
	(398,'平溪區','3'),
	(399,'雙溪區','3'),
	(400,'貢寮區','3'),
	(401,'新店區','3'),
	(402,'坪林區','3'),
	(403,'烏來區','3'),
	(404,'永和區','3'),
	(405,'中和區','3'),
	(406,'土城區','3'),
	(407,'三峽區','3'),
	(408,'樹林區','3'),
	(409,'鶯歌區','3'),
	(410,'三重區','3'),
	(411,'新莊區','3'),
	(412,'泰山區','3'),
	(413,'林口區','3'),
	(414,'蘆洲區','3'),
	(415,'五股區','3'),
	(416,'八里區','3'),
	(417,'淡水區','3'),
	(418,'三芝區','3'),
	(419,'石門區','3'),
	(420,'南竿鄉','22'),
	(421,'北竿鄉','22'),
	(422,'莒光鄉','22'),
	(423,'東引鄉','22'),
	(424,'宜蘭市','19'),
	(425,'壯圍鄉','19'),
	(426,'頭城鎮','19'),
	(427,'礁溪鄉','19'),
	(428,'員山鄉','19'),
	(429,'羅東鎮','19'),
	(430,'三星鄉','19'),
	(431,'大同鄉','19'),
	(432,'五結鄉','19'),
	(433,'冬山鄉','19'),
	(434,'蘇澳鎮','19'),
	(435,'南澳鄉','19'),
	(436,'釣魚臺','19'),
	(437,'東區','5'),
	(438,'北區','5'),
	(439,'香山區','5'),
	(440,'寶山鄉','6'),
	(441,'竹北市','6'),
	(442,'湖口鄉','6'),
	(443,'新豐鄉','6'),
	(444,'新埔鎮','6'),
	(445,'關西鎮','6'),
	(446,'芎林鄉','6'),
	(447,'竹東鎮','6'),
	(448,'五峰鄉','6'),
	(449,'橫山鄉','6'),
	(450,'尖石鄉','6'),
	(451,'北埔鄉','6'),
	(452,'峨眉鄉','6'),
	(453,'中壢區','4'),
	(454,'平鎮區','4'),
	(455,'龍潭區','4'),
	(456,'楊梅區','4'),
	(457,'新屋區','4'),
	(458,'觀音區','4'),
	(459,'桃園區','4'),
	(460,'龜山區','4'),
	(461,'八德區','4'),
	(462,'大溪區','4'),
	(463,'復興區','4'),
	(464,'大園區','4'),
	(465,'蘆竹區','4'),
	(466,'竹南鎮','7'),
	(467,'頭份市','7'),
	(468,'三灣鄉','7'),
	(469,'南庄鄉','7'),
	(470,'獅潭鄉','7'),
	(471,'後龍鎮','7'),
	(472,'通霄鎮','7'),
	(473,'苑裡鎮','7'),
	(474,'苗栗市','7'),
	(475,'造橋鄉','7'),
	(476,'頭屋鄉','7'),
	(477,'公館鄉','7'),
	(478,'大湖鄉','7'),
	(479,'泰安鄉','7'),
	(480,'銅鑼鄉','7'),
	(481,'三義鄉','7'),
	(482,'西湖鄉','7'),
	(483,'卓蘭鎮','7'),
	(484,'中區','8'),
	(485,'東區','8'),
	(486,'南區','8'),
	(487,'西區','8'),
	(488,'北區','8'),
	(489,'北屯區','8'),
	(490,'西屯區','8'),
	(491,'南屯區','8'),
	(492,'太平區','8'),
	(493,'大里區','8'),
	(494,'霧峰區','8'),
	(495,'烏日區','8'),
	(496,'豐原區','8'),
	(497,'后里區','8'),
	(498,'石岡區','8'),
	(499,'東勢區','8'),
	(500,'和平區','8'),
	(501,'新社區','8'),
	(502,'潭子區','8'),
	(503,'大雅區','8'),
	(504,'神岡區','8'),
	(505,'大肚區','8'),
	(506,'沙鹿區','8'),
	(507,'龍井區','8'),
	(508,'梧棲區','8'),
	(509,'清水區','8'),
	(510,'大甲區','8'),
	(511,'外埔區','8'),
	(512,'大安區','8'),
	(513,'彰化市','9'),
	(514,'芬園鄉','9'),
	(515,'花壇鄉','9'),
	(516,'秀水鄉','9'),
	(517,'鹿港鎮','9'),
	(518,'福興鄉','9'),
	(519,'線西鄉','9'),
	(520,'和美鎮','9'),
	(521,'伸港鄉','9'),
	(522,'員林市','9'),
	(523,'社頭鄉','9'),
	(524,'永靖鄉','9'),
	(525,'埔心鄉','9'),
	(526,'溪湖鎮','9'),
	(527,'大村鄉','9'),
	(528,'埔鹽鄉','9'),
	(529,'田中鎮','9'),
	(530,'北斗鎮','9'),
	(531,'田尾鄉','9'),
	(532,'埤頭鄉','9'),
	(533,'溪州鄉','9'),
	(534,'竹塘鄉','9'),
	(535,'二林鎮','9'),
	(536,'大城鄉','9'),
	(537,'芳苑鄉','9'),
	(538,'二水鄉','9'),
	(539,'南投市','10'),
	(540,'中寮鄉','10'),
	(541,'草屯鎮','10'),
	(542,'國姓鄉','10'),
	(543,'埔里鎮','10'),
	(544,'仁愛鄉','10'),
	(545,'名間鄉','10'),
	(546,'集集鎮','10'),
	(547,'水里鄉','10'),
	(548,'魚池鄉','10'),
	(549,'信義鄉','10'),
	(550,'竹山鎮','10'),
	(551,'鹿谷鄉','10'),
	(552,'西區','12'),
	(553,'東區','12'),
	(554,'番路鄉','13'),
	(555,'梅山鄉','13'),
	(556,'竹崎鄉','13'),
	(557,'阿里山鄉','13'),
	(558,'中埔鄉','13'),
	(559,'大埔鄉','13'),
	(560,'水上鄉','13'),
	(561,'鹿草鄉','13'),
	(562,'太保市','13'),
	(563,'朴子市','13'),
	(564,'東石鄉','13'),
	(565,'六腳鄉','13'),
	(566,'新港鄉','13'),
	(567,'民雄鄉','13'),
	(568,'大林鎮','13'),
	(569,'溪口鄉','13'),
	(570,'義竹鄉','13'),
	(571,'布袋鎮','13'),
	(572,'斗南鎮','11'),
	(573,'大埤鄉','11'),
	(574,'虎尾鎮','11'),
	(575,'土庫鎮','11'),
	(576,'褒忠鄉','11'),
	(577,'東勢鄉','11'),
	(578,'臺西鄉','11'),
	(579,'崙背鄉','11'),
	(580,'麥寮鄉','11'),
	(581,'斗六市','11'),
	(582,'林內鄉','11'),
	(583,'古坑鄉','11'),
	(584,'莿桐鄉','11'),
	(585,'西螺鎮','11'),
	(586,'二崙鄉','11'),
	(587,'北港鎮','11'),
	(588,'水林鄉','11'),
	(589,'口湖鄉','11'),
	(590,'四湖鄉','11'),
	(591,'元長鄉','11'),
	(592,'中西區','14'),
	(593,'東區','14'),
	(594,'南區','14'),
	(595,'北區','14'),
	(596,'安平區','14'),
	(597,'安南區','14'),
	(598,'永康區','14'),
	(599,'歸仁區','14'),
	(600,'新化區','14'),
	(601,'左鎮區','14'),
	(602,'玉井區','14'),
	(603,'楠西區','14'),
	(604,'南化區','14'),
	(605,'仁德區','14'),
	(606,'關廟區','14'),
	(607,'龍崎區','14'),
	(608,'官田區','14'),
	(609,'麻豆區','14'),
	(610,'佳里區','14'),
	(611,'西港區','14'),
	(612,'七股區','14'),
	(613,'將軍區','14'),
	(614,'學甲區','14'),
	(615,'北門區','14'),
	(616,'新營區','14'),
	(617,'後壁區','14'),
	(618,'白河區','14'),
	(619,'東山區','14'),
	(620,'六甲區','14'),
	(621,'下營區','14'),
	(622,'柳營區','14'),
	(623,'鹽水區','14'),
	(624,'善化區','14'),
	(625,'新市區','14'),
	(626,'大內區','14'),
	(627,'山上區','14'),
	(628,'安定區','14'),
	(629,'新興區','15'),
	(630,'前金區','15'),
	(631,'苓雅區','15'),
	(632,'鹽埕區','15'),
	(633,'鼓山區','15'),
	(634,'旗津區','15'),
	(635,'前鎮區','15'),
	(636,'三民區','15'),
	(637,'楠梓區','15'),
	(638,'小港區','15'),
	(639,'左營區','15'),
	(640,'仁武區','15'),
	(641,'大社區','15'),
	(642,'東沙群島','15'),
	(643,'南沙群島','15'),
	(644,'岡山區','15'),
	(645,'路竹區','15'),
	(646,'阿蓮區','15'),
	(647,'田寮區','15'),
	(648,'燕巢區','15'),
	(649,'橋頭區','15'),
	(650,'梓官區','15'),
	(651,'彌陀區','15'),
	(652,'永安區','15'),
	(653,'湖內區','15'),
	(654,'鳳山區','15'),
	(655,'大寮區','15'),
	(656,'林園區','15'),
	(657,'鳥松區','15'),
	(658,'大樹區','15'),
	(659,'旗山區','15'),
	(660,'美濃區','15'),
	(661,'六龜區','15'),
	(662,'內門區','15'),
	(663,'杉林區','15'),
	(664,'甲仙區','15'),
	(665,'桃源區','15'),
	(666,'那瑪夏區','15'),
	(667,'茂林區','15'),
	(668,'茄萣區','15'),
	(669,'馬公市','20'),
	(670,'西嶼鄉','20'),
	(671,'望安鄉','20'),
	(672,'七美鄉','20'),
	(673,'白沙鄉','20'),
	(674,'湖西鄉','20'),
	(675,'金沙鎮','21'),
	(676,'金湖鎮','21'),
	(677,'金寧鄉','21'),
	(678,'金城鎮','21'),
	(679,'烈嶼鄉','21'),
	(680,'烏坵鄉','21'),
	(681,'屏東市','17'),
	(682,'三地門鄉','17'),
	(683,'霧臺鄉','17'),
	(684,'瑪家鄉','17'),
	(685,'九如鄉','17'),
	(686,'里港鄉','17'),
	(687,'高樹鄉','17'),
	(688,'鹽埔鄉','17'),
	(689,'長治鄉','17'),
	(690,'麟洛鄉','17'),
	(691,'竹田鄉','17'),
	(692,'內埔鄉','17'),
	(693,'萬丹鄉','17'),
	(694,'潮州鎮','17'),
	(695,'泰武鄉','17'),
	(696,'來義鄉','17'),
	(697,'萬巒鄉','17'),
	(698,'崁頂鄉','17'),
	(699,'新埤鄉','17'),
	(700,'南州鄉','17'),
	(701,'林邊鄉','17'),
	(702,'東港鎮','17'),
	(703,'琉球鄉','17'),
	(704,'佳冬鄉','17'),
	(705,'新園鄉','17'),
	(706,'枋寮鄉','17'),
	(707,'枋山鄉','17'),
	(708,'春日鄉','17'),
	(709,'獅子鄉','17'),
	(710,'車城鄉','17'),
	(711,'牡丹鄉','17'),
	(712,'恆春鎮','17'),
	(713,'滿州鄉','17'),
	(714,'臺東市','23'),
	(715,'綠島鄉','23'),
	(716,'蘭嶼鄉','23'),
	(717,'延平鄉','23'),
	(718,'卑南鄉','23'),
	(719,'鹿野鄉','23'),
	(720,'關山鎮','23'),
	(721,'海端鄉','23'),
	(722,'池上鄉','23'),
	(723,'東河鄉','23'),
	(724,'成功鎮','23'),
	(725,'長濱鄉','23'),
	(726,'太麻里鄉','23'),
	(727,'金峰鄉','23'),
	(728,'大武鄉','23'),
	(729,'達仁鄉','23'),
	(730,'花蓮市','18'),
	(731,'新城鄉','18'),
	(732,'秀林鄉','18'),
	(733,'吉安鄉','18'),
	(734,'壽豐鄉','18'),
	(735,'鳳林鎮','18'),
	(736,'光復鄉','18'),
	(737,'豐濱鄉','18'),
	(738,'瑞穗鄉','18'),
	(739,'萬榮鄉','18'),
	(740,'玉里鎮','18'),
	(741,'卓溪鄉','18'),
	(742,'富里鄉','18');

/*!40000 ALTER TABLE `districts` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 favorites
# ------------------------------------------------------------

CREATE TABLE `favorites` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(45) DEFAULT NULL,
  `store_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `favorites` WRITE;
/*!40000 ALTER TABLE `favorites` DISABLE KEYS */;

INSERT INTO `favorites` (`id`, `user_id`, `store_id`)
VALUES
	(21,'10209887656465095',428);

/*!40000 ALTER TABLE `favorites` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 files
# ------------------------------------------------------------

CREATE TABLE `files` (
  `id` int NOT NULL AUTO_INCREMENT,
  `default_img` varchar(45) DEFAULT NULL,
  `img1` varchar(45) DEFAULT NULL,
  `img2` varchar(45) DEFAULT NULL,
  `store_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `files` WRITE;
/*!40000 ALTER TABLE `files` DISABLE KEYS */;

INSERT INTO `files` (`id`, `default_img`, `img1`, `img2`, `store_id`)
VALUES
	(61,'/img/1659021041001411000.jpeg','','',658),
	(62,'/img/1659023788038727000.jpeg','','',425);

/*!40000 ALTER TABLE `files` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 stores
# ------------------------------------------------------------

CREATE TABLE `stores` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `introduce` varchar(500) DEFAULT NULL,
  `address` varchar(45) DEFAULT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `mobile` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `user_id` int NOT NULL,
  `type_id` int DEFAULT NULL,
  `county_id` int DEFAULT NULL,
  `district_id` int DEFAULT NULL,
  `web_site` varchar(100) DEFAULT NULL,
  `is_dads` tinyint DEFAULT '0',
  `is_dads_recommend` tinyint DEFAULT '0',
  `lon` float DEFAULT NULL,
  `lat` float DEFAULT NULL,
  `is_close_permanently` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `stores` WRITE;
/*!40000 ALTER TABLE `stores` DISABLE KEYS */;

INSERT INTO `stores` (`id`, `name`, `introduce`, `address`, `phone`, `mobile`, `email`, `user_id`, `type_id`, `county_id`, `district_id`, `web_site`, `is_dads`, `is_dads_recommend`, `lon`, `lat`, `is_close_permanently`)
VALUES
	(425,'咱的海產店','恆春美女干貝醬/恆春伴手禮/墾丁伴手禮/恆春美食/墾丁海鮮/後壁湖海鮮','946台灣屏東縣恆春鎮大光路118-2號','','088866621','',0,1,17,712,'',1,1,120.74,21.9462,0),
	(426,'麵霸','','236台灣新北市土城區仁愛路42號','0961188790','','',0,0,3,406,'',0,0,121.462,24.9792,1),
	(427,'武藏田食堂','','236台灣新北市土城區清水路139號 1 樓','0908906881','','',0,1,3,406,'',1,0,121.458,24.9816,0),
	(428,'初原麵場(宜蘭宜中店)','拉麵','260台灣宜蘭縣宜蘭市宜中路216號','039332966','','',0,1,19,424,'',1,0,121.744,24.7504,0),
	(429,'屋台燒餃子','一中創始店','455台灣台中市北區雙十路二段11號號旁','0958830840',NULL,NULL,1,1,8,488,NULL,0,0,120.688,24.1522,1),
	(430,'黑市shabu','火鍋','251台灣新北市淡水區新春街146-9號','','0226210086','',0,1,3,417,'',1,0,121.442,25.1814,0),
	(431,'金鮮麻醬麵店','','台灣新北市三重區力行路一段275號1樓','','','',0,1,3,410,'',1,0,121.482,25.0716,0),
	(432,'南機場營養三明治','','100台灣台北市中正區中華路二段313巷20號','0912699203','','',0,1,2,374,'',1,0,121.507,25.0284,0),
	(433,'涓邸居酒屋','','115台灣台北市南港區三重路32巷8號','','0227832022','',0,1,2,382,'',1,0,121.614,25.0561,0),
	(434,'麵猴麵疙瘩','','220台灣新北市板橋區溪崑二街112巷1號附近','0968410076','','',0,1,3,393,'',1,0,121.425,24.9892,0),
	(435,'金華永和豆漿','','324台灣桃園市平鎮區金陵路三段86號','0966831318','','',0,1,4,454,'',1,0,121.227,24.928,0),
	(436,'無白牛排','潭子平價牛排|高CP美食|必吃牛排|推薦牛排|便宜牛排|人氣美食','427台灣台中市潭子區環中東路一段13號',NULL,'0425343050',NULL,1,1,8,502,NULL,1,0,120.704,24.2011,1),
	(437,'象泰陽泰式料理','','709台灣台南市安南區安吉路一段222號','0909805075',NULL,NULL,1,1,14,597,NULL,1,0,120.189,23.0466,1),
	(438,'和堂鍋物He Tang','火烤兩吃/石頭火鍋','900台灣屏東縣屏東市民教路85號','','087212890','',0,1,17,681,'',1,0,120.5,22.6677,0),
	(439,'Local Bistro • 在地餐酒館','','811台灣高雄市楠梓區德惠路72號','','073647514','',0,1,15,637,'',1,0,120.303,22.7272,0),
	(440,'泰泰好辣','','505台灣彰化縣鹿港鎮美市街53號','0903297879','','',0,1,9,517,'',1,0,120.434,24.0533,0),
	(441,'溫州大餛飩之家','原中華商場','108台灣台北市萬華區西寧南路63-3號','','0223822853','',0,1,2,377,'',1,0,121.507,25.0464,0),
	(442,'清心福全-桃園龍安店','珍珠奶茶手搖飲料專賣店','330台灣桃園市桃園區龍安街172之2號',NULL,'033780055',NULL,1,1,4,459,NULL,1,0,121.279,24.9922,0),
	(443,'菓點子-龍安店','','330台灣桃園市桃園區龍安街172號','','033786676','',0,1,4,459,'',1,0,121.279,24.9922,0),
	(444,'四食五號早餐店','小弟在鶯歌，慢慢接手爸媽的早餐店','239台灣新北市鶯歌區永和街45號','','0226796196','',0,1,3,409,'',1,0,121.332,24.9675,0),
	(445,'蚵男','生蠔 海物 燒烤','台灣台南市北區海安路三段926號','','062820808','',0,1,14,595,'',1,0,120.2,23.019,0),
	(446,'松仁串燒','日式烤雞串餐廳','800台灣高雄市新興區民族二路140號','','072238989','',0,1,15,629,'',1,0,120.307,22.6231,0),
	(447,'阿強西瓜','花蓮玉里。 ','台灣花蓮縣玉里鎮98131-15號','0911333793','','',0,1,18,740,'',1,0,121.339,23.4061,0),
	(448,'日輕好食','【桃園總店】| 輕食餐廳','330台灣桃園市桃園區同德五街25號1樓','','033585950','',0,1,4,459,'',1,0,121.301,25.0147,0),
	(449,'新街中西式早餐店','','551台灣南投縣名間鄉彰南路465號','','0492205989','',0,1,10,545,'',1,0,120.693,23.8793,0),
	(450,'醬厚呷官財板','台藝大文創園區對面','220台灣新北市板橋區龍興街99巷2號巷子轉角處','0917816105','','',0,1,3,393,'',1,0,121.444,25.003,0),
	(451,'韓雞Bar-台中興大旗艦店','','402台灣台中市南區國光路216-12號','','0422859661','',0,1,8,486,'https://www.han7bar.com/',1,0,120.679,24.1255,0),
	(452,'穆門什味鍋物','台中美食/台中火鍋/台中火鍋推薦/西區美食/西區火鍋/勤美美食','403台灣台中市西區向上北路182號','','0423011235','',0,1,8,487,'https://mumen-hotpot-tw.weebly.com/menu.html',1,0,120.661,24.148,0),
	(453,'可不可熟成紅茶-鳳山五甲店','','830台灣高雄市鳳山區自強二路115號',NULL,'078319171',NULL,1,1,15,654,NULL,1,0,120.327,22.5959,0),
	(454,'可不可熟成紅茶-高雄文山店','','830台灣高雄市鳳山區濱山街31號',NULL,'077777948',NULL,1,1,15,654,NULL,1,0,120.351,22.6458,0),
	(455,'筠芝和牛本舖','','111台灣台北市士林區中山北路六段35巷42號','','0228353135','',0,1,2,379,'',1,0,121.525,25.1026,0),
	(457,'韓哥哥銅盤烤肉','韓式燒烤餐廳','台灣新北市新莊區復興路一段203號','','0229900292','',0,1,3,411,'',1,0,121.452,25.0438,0),
	(459,'豪霸漢堡','','244台灣新北市林口區竹林路78號','','0226016982','',0,1,3,413,'https://brunch-restaurant-709.business.site/',1,0,121.388,25.0786,0),
	(460,'驀然餐酒館','鼓山小酌酒吧|必吃餐酒館|特色餐酒館|在地推薦酒館|巨蛋附近餐酒館|人氣餐酒館','804台灣高雄市鼓山區篤敬路35號','','075500352','',0,1,15,633,'',1,0,120.302,22.6643,0),
	(461,'等咧粉圓-恆春店','','946台灣屏東縣恆春鎮光明路50號','0986076958','','',0,1,17,712,'',1,0,120.744,22.0018,0),
	(462,'TUTTI57敦南會館','','105台灣台北市松山區敦化南路1段57號','','0225770507','',0,1,2,375,'',1,0,121.549,25.0457,0),
	(463,'善化牛肉x手工煎餃','','407台灣台中市西屯區河南路二段465號','','0427087979','',0,1,8,490,'',1,0,120.643,24.1693,0),
	(464,'台北蔥油餅','在歸仁黃昏市場','711台灣台南市歸仁區民權北路','','063387108','',0,1,14,599,'',1,0,120.288,22.9698,0),
	(465,'什一堂烘焙坊','','540台灣南投縣南投市民族路319號',NULL,'0492222953',NULL,1,1,10,539,'http://www.bakery-11.com.tw/',1,0,120.685,23.908,0),
	(466,'盛煎潮牌鍋貼專賣店','中友創始店 | 中友百貨公司','台灣台中市北區三民路三段161號B棟中友百貨公司',NULL,'0422210500',NULL,1,NULL,8,488,NULL,1,0,120.685,24.1522,0),
	(467,'爆Q美式炸雞','嘉義民族店','600台灣嘉義市西區民族路659號','0979579366',NULL,NULL,1,1,12,552,NULL,1,0,120.444,23.4754,0),
	(468,'謝家庄佛跳牆','','500台灣彰化縣彰化市彰南路二段151號',NULL,'047377077',NULL,1,1,9,513,NULL,1,0,120.58,24.0854,0),
	(469,'原野屋串燒','','744台灣台南市新市區中正路66巷17號',NULL,'065896226',NULL,1,1,14,625,NULL,1,0,120.292,23.0719,0),
	(470,'底厝DiCho','僅預訂自取 營業時間以IG公告為主','600台灣嘉義市東區忠孝路342號',NULL,NULL,NULL,1,1,12,553,'https://www.instagram.com/dicho2019/?igshid=155o6m6x2kr5j',1,0,120.454,23.489,0),
	(471,'鱻將割烹','正宗日式料理餐廳','251台灣新北市淡水區重建街71號',NULL,'0226291979',NULL,1,1,3,417,NULL,1,0,121.44,25.1726,0),
	(472,'杜芳子古味茶鋪-新莊思源店','新莊思源店','242台灣新北市新莊區思源路351號',NULL,'0229967958',NULL,1,1,3,411,'https://ice-cream-and-drink-shop-6278.business.site/?utm_source=gmb&utm_medium=referral',1,0,121.46,25.0505,0),
	(474,'漁網居酒屋','新北耶誕城|人氣生魚片刺身|海鮮燒烤|串燒|小酌聚餐推薦|','220台灣新北市板橋區民生路二段234巷24弄 4號一樓',NULL,'0222545101',NULL,1,1,3,393,NULL,1,0,121.47,25.0224,0),
	(475,'日光找食','','235台灣新北市中和區福美路240號',NULL,'0222439321',NULL,1,1,3,405,NULL,1,0,121.499,25.0053,1),
	(476,'台灣芒果冰專賣店','','241台灣新北市三重區溪尾街279號',NULL,'0229882116',NULL,1,1,3,410,NULL,1,0,121.487,25.0824,0),
	(477,'黑堂甜品專賣店','','台灣新北市三重區正義南路39號',NULL,'0229747070',NULL,1,1,3,410,NULL,1,0,121.498,25.0609,0),
	(478,'三番日式食堂','','300台灣新竹市北區湳雅街42-1號',NULL,'035326809',NULL,1,1,5,438,NULL,1,0,120.971,24.8177,0),
	(479,'南魂','串燒·酒場','106台灣台北市大安區通安街37號',NULL,'0227002717',NULL,1,1,2,376,NULL,1,0,121.553,25.0316,0),
	(480,'鉑聚炭火料理工坊','','號 No. 140, 美術南二路鼓山區高雄市台灣 804',NULL,'075530123',NULL,1,1,15,633,NULL,1,0,120.285,22.6523,1),
	(481,'阿育哥大腸麵線','','235台灣新北市中和區員山路568號','0952886913',NULL,NULL,1,1,3,405,NULL,1,0,121.479,25.0107,0),
	(482,'順敏水果行','','830台灣高雄市鳳山區五甲一路451-16號',NULL,NULL,NULL,1,1,15,654,NULL,1,0,120.353,22.6136,0),
	(484,'黑炭燒烤本鋪','','202台灣基隆市中正區信四路18號',NULL,'0224272477',NULL,1,1,1,386,NULL,1,0,121.746,25.1339,0),
	(485,'香園牛肉麵','金山店','807台灣高雄市三民區金山路208號',NULL,'073483969',NULL,1,1,15,636,NULL,1,0,120.322,22.6626,0),
	(487,'飪室咖哩 RenshiCurry','勤美店','403台灣台中市西區美村路一段149巷19號',NULL,'0423220857',NULL,1,NULL,NULL,NULL,'https://www.renshi20161101.com/',1,0,120.663,24.1515,0),
	(488,'基隆活海產','（鱉大王）','807台灣高雄市三民區鼎力路114號',NULL,'073473229',NULL,1,1,15,636,NULL,1,0,120.322,22.6652,0),
	(489,'拱心意義大利麵','','115台灣台北市南港區昆陽街150-4號',NULL,'0226517217',NULL,1,1,2,382,NULL,1,0,121.594,25.0491,0),
	(490,'探扒霸','平價午晚餐/南屯美食/南屯美食推薦/南屯小吃/南屯熱炒/南屯炒飯/南屯必吃','408台灣台中市南屯區環中路四段1368之28號',NULL,'0424793099',NULL,1,NULL,8,491,NULL,1,0,120.63,24.1269,0),
	(491,'迥畫廊咖啡','自家烘焙咖啡','510台灣彰化縣員林市台鳳路2號',NULL,'048358800',NULL,1,NULL,9,522,'http://www.changhualancafe.com/',1,0,120.568,23.9635,0),
	(492,'Rich Daddy cafe & kids 富爸爸餐飲會所','','100台灣台北市中正區金山北路1號2樓',NULL,'0223962518',NULL,1,1,2,372,NULL,1,0,121.53,25.0442,0),
	(494,'干單茶飲','','885台灣澎湖縣湖西鄉161-4號','0926662100',NULL,NULL,1,1,20,674,NULL,1,0,119.661,23.5848,0),
	(495,'咻蹦花火冰菓室','','303台灣新竹縣湖口鄉中山路二段262號','0934366224',NULL,NULL,1,1,6,442,NULL,1,0,121.046,24.9049,0),
	(496,'青畑九號豆製所','台中美村店','403台灣台中市西區美村路一段87號',NULL,'0423265416',NULL,1,1,8,487,NULL,1,0,120.662,24.153,1),
	(497,'庄稼餐館','','900台灣屏東縣屏東市信義路319號',NULL,'087663598',NULL,1,1,17,681,NULL,1,0,120.491,22.6848,0),
	(498,'南寮風車有機農場休息站','','885台灣澎湖縣湖西鄉南寮村200號',NULL,'069920561',NULL,1,1,20,674,NULL,1,0,119.669,23.5865,0),
	(499,'隼丼 | 海鮮丼','日式餐廳','32056台灣桃園市中壢區領航南路三段279號',NULL,'032873871',NULL,1,1,4,453,'https://www.facebook.com/hayabusaDonburi/menu',1,0,121.221,25.0122,0),
	(500,'Q Burger-新莊新豐店','新莊新豐店','242台灣新北市新莊區新豐街25號',NULL,'0222010220',NULL,1,1,3,411,NULL,1,0,121.448,25.034,0),
	(501,'Q Burger-板橋文聖店','早餐','220台灣新北市板橋區文聖街142號',NULL,'0222500123',NULL,1,1,3,393,NULL,1,0,121.479,25.028,0),
	(502,'又一村','龍城市場 刺身 握壽司 丼飯','105台灣台北市松山區光復北路190巷39號','0922515453',NULL,NULL,1,1,2,375,NULL,1,0,121.554,25.0547,0),
	(503,'牛佬牛肉麵專門店','','406台灣台中市北屯區興安路一段20號',NULL,'0422385300',NULL,1,1,8,489,NULL,1,0,120.691,24.1658,0),
	(504,'牛佬牛肉麵專門店-十期店','十期店','406台灣台中市北屯區建和路二段20號',NULL,'0424351535',NULL,1,1,8,489,NULL,1,0,120.725,24.1672,0),
	(505,'牛佬牛肉麵專門店-西屯店','西屯店','407台灣台中市西屯區黎明路三段541號',NULL,'0424522512',NULL,1,1,8,490,NULL,1,0,120.64,24.1867,0),
	(506,'牛佬牛肉麵專門店-太平店','太平店','411台灣台中市太平區樹德路370號',NULL,'0423912093',NULL,1,1,8,492,NULL,1,0,120.713,24.1595,0),
	(507,'日晨咖啡烘焙','','540台灣南投縣南投市彰南路二段102號',NULL,'0492244037',NULL,1,1,10,539,'https://cafe-sundawn.business.site/',1,0,120.683,23.9104,0),
	(508,'Range鉄板居酒屋','','320台灣桃園市中壢區廣安街32號',NULL,'034222588',NULL,1,1,4,453,NULL,1,0,121.223,24.9585,1),
	(509,'焼肉ショジョ Yakiniku SHOJO-台南旗艦店','台南全球旗艦店','701台灣台南市東區育樂街250號',NULL,'062088400',NULL,1,NULL,NULL,NULL,NULL,1,0,120.214,22.992,0),
	(510,'焼肉ショジョ Yakiniku SHOJO-台南安平店','台南安平店','708台灣台南市安平區怡平路1號',NULL,'062989006',NULL,1,1,14,596,NULL,1,0,120.183,22.9926,0),
	(511,'焼肉ショジョ Yakiniku SHOJO-高雄左營店','高雄左營店','80452台灣高雄市鼓山區大順一路441號',NULL,'075527002',NULL,1,1,15,633,NULL,1,0,120.303,22.6558,0),
	(512,'焼肉ショジョ Yakiniku SHOJO-高雄形象概念店','高雄形象概念店','802台灣高雄市苓雅區同慶路31號',NULL,'077228814',NULL,1,NULL,NULL,NULL,NULL,1,0,120.323,22.628,0),
	(513,'焼肉ショジョ Yakiniku SHOJO-台中公益店','台中公益店','408012台灣台中市南屯區公益路二段919號',NULL,'0423899888',NULL,1,1,8,491,NULL,1,0,120.627,24.1513,0),
	(514,'焼肉ショジョ Yakiniku SHOJO-雲林斗六店','雲林斗六店','640台灣雲林縣斗六市中堅西路585號',NULL,'055361761',NULL,1,1,11,581,'https://yakinikushojoyunlin.business.site/',1,0,120.534,23.701,0),
	(515,'根記碗粿大業店','','408台灣台中市南屯區大業路133號',NULL,'0423275198',NULL,1,1,8,491,NULL,0,0,120.652,24.1535,1),
	(516,'林內下厝臭豆腐','','643台灣雲林縣林內鄉進興路60號','0937781871',NULL,NULL,1,1,11,582,NULL,1,0,120.583,23.7729,0),
	(517,'TopFire頂焰精肉小酒館','','105台灣台北市松山區市民大道四段219號',NULL,'0225796677',NULL,1,1,2,375,NULL,1,0,121.553,25.0446,0),
	(518,'Golden-蔗家店','（延吉總店）','106台灣台北市大安區延吉街135號',NULL,'0287722005',NULL,1,1,2,376,NULL,1,0,121.554,25.0424,0),
	(519,'朱佳水餃','2022/7/9搬家新北市板橋區新月一街80號','220台灣新北市板橋區四維路279巷38號','0936969422',NULL,NULL,1,1,3,393,NULL,1,0,121.461,25.028,0),
	(520,'燚茶咖啡','','64341台灣雲林縣林內鄉中正路187號',NULL,'055890600',NULL,1,1,11,582,NULL,1,0,120.616,23.7588,0),
	(521,'瑪賀牛肉麵-板橋店','板橋店','220台灣新北市板橋區吳鳳路64號',NULL,'0282522522',NULL,1,1,3,393,'https://www.mahobeefnoodles.com/',1,0,121.468,25.0307,0),
	(522,'阿妙意麵','','737台灣台南市鹽水區康樂路67號',NULL,'066531031',NULL,1,1,14,623,NULL,1,0,120.268,23.3196,0),
	(523,'麥食達韓式料理','','100台灣台北市中正區懷寧街86號',NULL,'0223899048',NULL,1,1,2,372,NULL,1,0,121.514,25.0429,0),
	(524,'山男咖啡','','946台灣屏東縣恆春鎮紅柴路2-9號',NULL,NULL,NULL,1,1,17,712,NULL,1,0,120.718,21.981,0),
	(525,'思茶MissingTea手作飲品-新竹新豐店','新竹新豐店','304台灣新竹縣新豐鄉建興路92號',NULL,'035571414',NULL,1,1,6,443,NULL,1,0,120.995,24.8719,0),
	(526,'柒玖冰室','甘蔗冰 飲料 豆花 冰 手搖飲 埔里','545台灣南投縣埔里鎮南盛街79號',NULL,'0492992397',NULL,1,1,10,543,NULL,1,0,120.967,23.9624,0),
	(527,'台北港生猛活海鮮','','116台灣台北市文山區木新路三段302號',NULL,'0229372155',NULL,1,1,2,383,NULL,1,0,121.558,24.981,0),
	(529,'小晨事(果汁茶飲店)','營建署旁八德路上','105台灣台北市松山區八德路二段346巷1弄7號',NULL,'0287737398',NULL,1,1,2,375,NULL,1,0,121.545,25.0477,0),
	(532,'洋爸爸健康餐-中華店','中華店','701台灣台南市東區中華東路二段185巷6號',NULL,'062688898',NULL,1,1,14,593,NULL,1,0,120.233,22.9845,0),
	(533,'洋爸爸健康餐-永康創始店','','71052台灣台南市永康區忠義街20號',NULL,'063125765',NULL,1,1,14,598,NULL,1,0,120.237,22.9992,0),
	(534,'棧直火廚房C-kitchen by FUK','','10686台灣台北市大安區仁愛路四段345巷4弄4號',NULL,'0287739938',NULL,1,1,2,376,NULL,1,0,121.554,25.0388,0),
	(535,'甜久坊','','108台灣萬華區西寧南路50巷7號','0926284312',NULL,NULL,1,1,2,377,NULL,1,0,121.506,25.0444,0),
	(536,'上味食堂','台東美食|台東人氣餐廳|台東小吃|台東銅板美食|台東滷肉飯節嚴選商家|台東滷肉飯|IG打卡餐廳','950台灣台東縣台東市新生路447號',NULL,NULL,NULL,1,NULL,23,714,NULL,1,0,121.14,22.7598,0),
	(539,'幸福鍋貼-府前店','府前店','700台灣台南市中西區府前路一段28號',NULL,'062207789',NULL,1,1,14,592,NULL,1,0,120.21,22.989,0),
	(540,'仙鼎茶商行','','262台灣宜蘭縣礁溪鄉礁溪路二段61號',NULL,'039280187',NULL,1,1,19,427,NULL,1,0,121.764,24.7971,0),
	(541,'杰迪烘焙坊','','973台灣花蓮縣吉安鄉吉祥四街53號',NULL,'038534562',NULL,1,1,18,733,NULL,1,0,121.584,23.9754,0),
	(542,'喫上飲','','106台灣台北市大安區忠孝東路三段217巷8弄2-1號',NULL,'0227317478',NULL,1,1,2,376,NULL,1,0,121.54,25.0437,0),
	(543,'虎笑麵屋','牛肉湯/老虎麵專賣','105台灣台北市松山區三民路113巷18號1樓',NULL,'0227666978',NULL,1,1,2,375,NULL,1,0,121.562,25.0581,0),
	(544,'吉品香小火鍋','','812台灣高雄市小港區宏平路372號',NULL,'078053392',NULL,1,1,15,638,NULL,1,0,120.359,22.5701,0),
	(545,'幸福鍋貼-永福店','永福店','700台灣台南市中西區永福路二段70號',NULL,'062202926',NULL,1,NULL,14,592,NULL,1,0,120.201,22.9921,0),
	(546,'紅茶洋行-國聯店','（國聯店）','970台灣花蓮縣花蓮市國民九街55號','0920320686',NULL,NULL,1,1,18,730,'https://ice-cream-and-drink-shop-3592.business.site',1,0,121.606,23.9947,0),
	(547,'冬瓜幫','花蓮公園店','970台灣花蓮縣花蓮市公園路2號',NULL,'038330661',NULL,1,1,18,730,NULL,1,0,121.61,23.9763,0),
	(548,'宏冠排骨麵','','40666台灣台中市北屯區天津路三段121號',NULL,'0422327588',NULL,1,1,8,489,NULL,1,0,120.683,24.1696,0),
	(550,'飛翔的魚flyingfish_brunch','美式料理餐廳 早午餐 漢堡 三明治 墨西哥烤餅 義大利麵 燉飯 人氣甜點 特色飲品推薦 異國料理','300台灣新竹市北區興南街31號',NULL,'035221107',NULL,1,1,5,438,NULL,1,0,120.963,24.8012,0),
	(551,'Schumann\'s Bistro No. 6 舒曼六號餐館','南京店','105台灣台北市松山區南京東路三段303巷8弄11號',NULL,'0227188108',NULL,1,1,2,375,'https://www.schumannsbistrono6.com/',1,0,121.547,25.053,0),
	(552,'坤伯傳家餃','','221台灣新北市汐止區南陽街250號',NULL,'0226933663',NULL,1,1,3,394,NULL,1,0,121.627,25.0567,0),
	(553,'久違石頭火鍋','(推薦人氣必吃) 海鮮 個人鍋物 宵夜聚餐','407台灣台中市西屯區河南路二段245號',NULL,'0427060680',NULL,1,1,8,490,NULL,1,0,120.648,24.1745,0),
	(554,'日光造咖','咖哩製造所 逢甲河南店','407台灣台中市西屯區河南路二段241之101號',NULL,'0424528122',NULL,1,1,8,490,'https://restaurant-70962.business.site/',1,0,120.649,24.1749,0),
	(555,'C.Chill Days(混日子)運動飛鏢吧','','500台灣彰化縣彰化市自強路110-2號',NULL,'047299977',NULL,1,4,9,513,NULL,1,0,120.541,24.0898,0),
	(556,'五億雞排','新竹白飯魚店','30074台灣新竹市東區介壽路13號','0903239259',NULL,NULL,1,1,5,437,NULL,1,0,121.021,24.7813,0),
	(557,'誠壽司','','220台灣新北市板橋區文化路二段182巷7弄5號',NULL,'0222503611',NULL,1,1,3,393,NULL,1,0,121.472,25.0272,0),
	(558,'蘭波 Lan Po','','260台灣宜蘭縣宜蘭市女中路二段185號',NULL,'039358516',NULL,1,1,19,424,'https://shop.ichefpos.com/store/80u2PXf9/ordering',1,0,121.757,24.7467,0),
	(559,'歐膩ohni麻糬','','326台灣桃園市楊梅區新興街57號','0908686897',NULL,NULL,1,1,4,456,NULL,1,0,121.184,24.918,0),
	(560,'熊貓廚房','小弟在內湖737美食街 賣炒飯炒泡麵 Fb可以搜尋 也有uber外送。','114台灣台北市內湖區內湖路一段737巷',NULL,NULL,NULL,1,1,2,381,NULL,1,0,121.579,25.0797,0),
	(561,'張張手作手工包子饅頭','Chang Chang Homemade「手工包子饅頭」','420台灣台中市豐原區永康路262號','0956562006',NULL,NULL,1,1,8,496,NULL,1,0,120.724,24.2407,0),
	(562,'邑品涼麵','','326台灣桃園市楊梅區新農街465號',NULL,'034882727',NULL,1,1,4,456,NULL,1,0,121.159,24.9127,0),
	(563,'Louisa Coffee 路易．莎咖啡(高雄華夏門市)','如果從高鐵高雄站出發或是來到高雄<br>我們路易莎咖啡華夏店歡迎您<br>Louisa Coffee 路易．莎咖啡(高雄華夏門市)<br>07 348 4613<br>https://goo.gl/maps/ME42G9o2NvkpHSKH7<br>高雄大叔<br>也可以來路易莎咖啡華夏店坐坐喔','813高雄市左營區華夏路1303號',NULL,'073484613',NULL,1,1,15,639,NULL,1,0,120.311,22.6858,0),
	(564,'漢寶哥（上客）炸雞','','357台灣苗栗縣通霄鎮自強路a59號號','0910455898',NULL,NULL,1,2,7,472,NULL,1,0,120.683,24.4887,0),
	(566,'超吉大盛-昌平店','','406台灣台中市北屯區昌平路一段26-1號',NULL,'0422337258',NULL,1,1,8,489,NULL,1,0,120.694,24.1663,0),
	(567,'烏弄Unocha-大社中山店','大社中山店','815台灣高雄市大社區中山路203號',NULL,'073538855',NULL,1,1,15,641,NULL,1,0,120.347,22.7321,0),
	(568,'香祺上海烤饅頭','','800台灣高雄市新興區中正三路186號',NULL,'072360787',NULL,1,1,15,629,NULL,1,0,120.303,22.6316,0),
	(569,'金城檸檬冰','','236台灣新北市土城區金城路二段68號','0966679274',NULL,NULL,1,1,3,406,NULL,1,0,121.453,24.9799,0),
	(570,'連大牛排小火鍋','','643台灣雲林縣林內鄉中正路172號',NULL,'055899899',NULL,1,1,11,582,NULL,1,0,120.616,23.7593,0),
	(571,'艾利波波竹圍店','','337台灣桃園市大園區建國路141號',NULL,'033935088',NULL,1,1,4,464,NULL,1,0,121.25,25.1094,0),
	(572,'咕茶','超推薦，我一星期至少喝10杯','237台灣新北市三峽區光明路100之1號',NULL,'0226735758',NULL,1,1,3,407,NULL,1,0,121.376,24.9281,0),
	(573,'杏糖屋','','114台灣台北市內湖區民權東路六段90巷16弄1號',NULL,'0227918623',NULL,1,1,2,381,NULL,1,0,121.585,25.0679,0),
	(574,'涼心涼麵 cool your heart️','','802台灣高雄市苓雅區仁智街19號','0936833930',NULL,NULL,1,1,15,631,'https://cool-your-heart.business.site',1,0,120.305,22.6149,0),
	(575,'阿香澎湖碳烤','','832台灣高雄市林園區半廍路16巷1之1號','0931796486',NULL,NULL,1,1,15,656,NULL,1,0,120.386,22.49,0),
	(576,'頂富現做赤肉羹','','236台灣新北市土城區中央路四段24號','0989339809',NULL,NULL,1,1,3,406,NULL,1,0,121.421,24.9611,0),
	(578,'南法小步舞曲','','704台灣台南市北區賢北街15巷9號','0989864825',NULL,NULL,1,5,14,595,'https://twstay.com/RWD1/index.aspx?BNB=mufu#templatemo',1,0,120.184,23.0063,0),
	(579,'AQUA Lounge','喝喝小酒！抽抽水煙！談談心！忠孝東路四段216巷32弄14號','106台灣台北市大安區忠孝東路四段216巷32弄14號',NULL,'0227219440',NULL,1,1,2,376,NULL,1,0,121.552,25.0398,0),
	(580,'KoDō和牛燒肉','','403台灣台中市西區公益路260號',NULL,'0423220312',NULL,1,1,8,487,NULL,1,0,120.655,24.1511,0),
	(581,'賈董的麵','台中中山醫學院美食街','402台灣台中市南區建國北路一段110號',NULL,'0424739595',NULL,1,1,8,486,NULL,1,0,120.65,24.1215,0),
	(583,'呷品香滷味','','220台灣新北市板橋區重慶路290巷8號',NULL,'0289534179',NULL,1,1,3,393,NULL,1,0,121.463,24.999,0),
	(584,'景興古早味','','406台灣台中市北屯區河北路三段123號','0920159332',NULL,NULL,1,1,8,489,NULL,1,0,120.684,24.1801,0),
	(585,'大宅門乾鍋鴨頭','','81360台灣高雄市左營區自由二路338號',NULL,'075564078',NULL,1,1,15,639,NULL,1,0,120.309,22.6657,0),
	(586,'御私藏鮮奶茶專賣店-瑞隆店','瑞隆店','806台灣高雄市前鎮區瑞隆路206號',NULL,'077213855',NULL,1,1,15,635,'https://www.possession-tea.com/',1,0,120.334,22.6058,0),
	(587,'喫堡早午餐','','265台灣宜蘭縣羅東鎮中山路四段218號',NULL,'039511289',NULL,1,1,19,429,NULL,1,0,121.759,24.6735,0),
	(588,'一也黃金豆乳雞','台南公園店','700台灣台南市中西區公園路121號','0938632917',NULL,NULL,1,1,14,592,'https://chicken-wings-restaurant-2.business.site/',1,0,120.206,22.9968,0),
	(589,'老賴茶棧-桃園桃鶯店','桃園桃鶯店','33067台灣桃園市桃園區桃鶯路213號',NULL,'033672673',NULL,1,1,4,459,'https://www.liketeashop.com/',1,0,121.318,24.9846,0),
	(590,'老賴茶棧-新莊四維店','新莊四維店','242台灣新北市新莊區四維路159號',NULL,'0222046660',NULL,1,1,3,411,NULL,1,0,121.423,25.0218,0),
	(591,'SUNDaWN COFFEE ROASTER','營業時間：09:00-17:00 週一公休 自家烘焙咖啡廳 全店使用單品豆作為出發 可客選濾杯做沖煮','南投縣草屯鎮草溪路906號',NULL,'049-2313460',NULL,1,1,10,541,NULL,1,0,120.68,23.9746,0),
	(592,'Hungry Dogs 二犬廚房','餐盒','115台灣台北市南港區興中路44巷28號1樓','0937037702',NULL,NULL,1,1,2,382,'https://hungrydogs.business.site/',1,0,121.606,25.0558,0),
	(593,'波妞芋圓','','350台灣苗栗縣竹南鎮龍山路一段222號',NULL,'037551102',NULL,1,1,7,467,NULL,1,0,120.867,24.6935,0),
	(594,'Jocha吼呷-熟成廚房','精品熟成食材-酒食私廚','262台灣宜蘭縣礁溪鄉柴圍路9鄰77-5號',NULL,'039288992',NULL,1,1,19,427,NULL,1,0,121.756,24.8053,0),
	(595,'Summer Box 夏天的盒子','','946003台灣屏東縣恆春鎮恆公路967號','0936155472',NULL,NULL,1,1,17,712,NULL,1,0,120.741,22.0091,0),
	(596,'心心漢堡丨xinxinBurger','手作和牛漢堡吃到飽','220台灣新北市板橋區新海路120號',NULL,'0222511099',NULL,1,1,3,393,NULL,1,0,121.461,25.0229,0),
	(598,'有煎餃子館-蘆洲正和店','(蘆洲正和店)','247台灣新北市蘆洲區正和街27號',NULL,'0222812095',NULL,1,1,3,414,NULL,1,0,121.47,25.082,0),
	(599,'心心麻辣鍋-西門店','台北和牛吃到飽 ｜ 特色七龍珠火鍋 ｜ 西門町人氣美食','108台灣台北市萬華區成都路141號',NULL,'0223889299',NULL,1,1,2,377,NULL,1,0,121.503,25.0435,0),
	(600,'吳家媽媽地瓜酥-安慶門市','安慶門市','950台灣台東縣台東市安慶街57號','089333732',NULL,NULL,1,1,23,714,'https://ogamama.1shop.tw/tn3cis',1,0,121.151,22.7528,0),
	(601,'好石跡精緻石頭火鍋-蘆洲集賢店','(蘆洲集賢店)','247台灣新北市蘆洲區集賢路222巷9號',NULL,'0222887576',NULL,1,1,3,414,NULL,1,0,121.482,25.0872,0),
	(602,'沛本丸-手作飯丸','','326台灣桃園市楊梅區新農街443號1樓',NULL,NULL,NULL,1,1,4,456,NULL,1,0,121.159,24.9126,0),
	(603,'樂咖 Lego Brick Cafe','','800台灣高雄市新興區復興一路70-10號','0919880552',NULL,NULL,1,1,15,629,NULL,1,0,120.308,22.6321,0),
	(604,'鋒憶麵坊','','338台灣桃園市蘆竹區南上路14號',NULL,'032221537',NULL,1,1,4,465,NULL,1,0,121.295,25.0475,0),
	(605,'這是車輪餅-和美店','（和美店）','台灣彰化縣和美鎮彰美路六段1號','0979153098',NULL,NULL,1,NULL,NULL,NULL,NULL,1,0,120.498,24.1124,0),
	(606,'木子樑室-岡山店','岡山店','82042台灣高雄市岡山區維仁路151之1號',NULL,'076222219',NULL,1,1,15,644,'https://ice-cream-and-drink-shop-4728.business.site/?utm_source=gmb&utm_medium=referral',1,0,120.293,22.7949,0),
	(607,'大師姐足體養生館','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.223,24.9634,0),
	(608,'就水 Jo Shui','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.748,25.1293,0),
	(609,'足工場精緻足體養生會館','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.561,25.0831,0),
	(610,'部長電池','台南東區',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.246,22.9832,0),
	(611,'語璽新能源科技有限公司','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.357,25.0698,0),
	(612,'繼達汽車材料','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.277,22.9964,0),
	(613,'阿寶機車行','道路救援 機車維修 AMA行車紀錄器',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.64,24.1385,0),
	(614,'阿寶機車行','西屯店 道路救援 機車維修 AMA行車紀錄器',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.647,24.1701,0),
	(615,'阿寶機車行','一乙久店',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.662,24.1365,0),
	(616,'PIT崴特國際 Alfa Romeo Fiat Abarth','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.595,25.0625,0),
	(617,'聖帕斯強化考爾','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.311,24.9875,0),
	(618,'聖帕斯強化考爾總公司倉庫','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.311,24.9873,0),
	(619,'中美車坊','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.323,24.9739,0),
	(620,'車際聯盟汽車鍍膜','林口施工店',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.372,25.0525,0),
	(621,'洗刷刷自助洗車場','草屯店',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.675,23.9549,0),
	(622,'HOT大聯盟 廷豐車業','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.309,24.9755,0),
	(623,'北大藥局','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.379,24.9438,0),
	(624,'展榮商號','鹿港店／堅果／五穀粉／糙米麩／麵茶／五穀雜糧／寶寶米果／米香／黑豆茶／養生／米麩／推薦／彰化',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.439,24.0497,0),
	(625,'展榮商號','堅果／五穀粉／糙米麩／麵茶／五穀雜糧／寶寶米果／米香／黑豆茶／養生／伴手禮／禮盒/米麩／推薦／台中烏日',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.626,24.1113,0),
	(626,'自產自榨自銷一條龍小農第一道冷壓初榨橄欖油','，真的可以喝的橄欖原汁<br>台灣可以百貨公司網路購買，主要是做外銷 <br>https://taiwan.auganic.co/<br>台北公司 ; 新北市汐止區康寧街 169巷 25號 2F-3<br>澳洲地址 : P.O BOX 112 GINGIN 6503 WESTERN AUSTRALIA',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.623,25.0684,0),
	(627,'柒目刺青屋 Seven Eyes Tattoo House','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.299,22.726,0),
	(628,'丸群水產行','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.19,22.8745,0),
	(629,'森涵淨水有限公司','新北居家飲用水|三重淨水設備推薦|全屋淨水器|Ro飲用水|生飲設備施作|商用生飲設備',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.494,25.0721,0),
	(630,'【芽米寶貝】寶寶副食品','總公司',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.485,25.0008,0),
	(631,'40 INN 背包倉庫肆零居','Backpackers Warehouse ( 台東國際背包客棧)',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.15,22.7542,0),
	(632,'墾丁船帆石風尚會館','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,120.824,21.9334,0),
	(633,'盛湖草莓園','','364台灣苗栗縣大湖鄉364','0976057425',NULL,NULL,1,4,7,478,NULL,1,0,120.864,24.4254,0),
	(634,'白宮行舘沙灘溫泉','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.69,25.1809,0),
	(635,'黑炭燒烤本鋪','',NULL,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,0,0,121.746,25.1339,0),
	(638,'不二煮藝','','106台灣台北市大安區仁愛路四段112巷9弄5號1樓','','0227009880','',0,1,2,376,'',1,0,121.55,25.037,0),
	(639,'大師姐足體養生館','','320桃園市中壢區新生路269號','','034255072','',0,4,4,453,'https://massage-center-99.business.site/',1,0,121.223,24.9635,0),
	(640,'就水 Jo Shui','','201基隆市信義區義四路7-1號2','','0224251281','',0,7,1,385,'',1,0,121.747,25.1295,0),
	(641,'足工場精緻足體養生會館','','10462台北市中山區堤頂大道二段500號','','0285025757','',0,4,2,374,'',1,0,121.561,25.0833,0),
	(642,'大叔理髮俱樂部','','404台中市北區崇德路一段499號','','','',0,7,8,488,'',1,0,120.685,24.1646,0),
	(643,'部長電池-台南東區','','701台南市東區裕文一街58號','0912762822','','',0,2,14,593,'',1,0,120.246,22.9834,0),
	(644,'墾丁船帆石風尚會館','','946屏東縣恆春鎮船帆路676巷5-1號','','088851196','',0,5,17,712,'http://www.funsonhotel.com.tw/',1,0,120.824,21.9335,0),
	(645,'40 INN 背包倉庫肆零居','Backpackers Warehouse ( 台東國際背包客棧)','950台東縣台東市文化街40號','0906404040','','',0,5,23,714,'',1,0,121.15,22.7543,0),
	(646,'蘭陽溪口親子餐廳','','268宜蘭縣五結鄉溪濱路一段133號','0929338803','','',0,1,19,432,'',1,0,121.828,24.7048,0),
	(647,'均韋自動車','基本維修保養(安檢)\n各大牌品輪胎銷售\n底盤調教 四輪定位\n動力提升 日系 歐系\n各車種晶片調教\nHS power flash ; E tuners\n精品改裝\n卡鉗 避震器 鋁圈 排氣管 電裝品\n渦輪 水噴','324桃園市平鎮區民族路雙連三段89-1號','0915213970','','',0,2,4,454,'',1,0,121.178,24.9588,0),
	(648,'品藴家 PIN JUN JIA','神經病大叔特調','500彰化縣彰化市華山路121號','','047260999','',0,1,9,513,'https://shop.ichefpos.com/store/dDZf1oWP/ordering',1,0,120.54,24.075,0),
	(649,'ONE BEAR 一隻熊甜點森林','','436台中市清水區五權路90號','','0426265866','',0,1,8,509,'',1,0,120.569,24.2771,0),
	(650,'紅心地瓜球-龍井店','龍德文具前','台中市龍井區中央路二段104號','0903673995','','',0,1,8,507,'',1,0,120.523,24.2055,0),
	(658,'暖南旅宿','','946屏東縣恆春鎮恆公路238號','','088892080','',0,5,17,712,'',1,0,0,0,0);

/*!40000 ALTER TABLE `stores` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 thumbs
# ------------------------------------------------------------

CREATE TABLE `thumbs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `store_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `thumbs` WRITE;
/*!40000 ALTER TABLE `thumbs` DISABLE KEYS */;

INSERT INTO `thumbs` (`id`, `store_id`, `user_id`)
VALUES
	(6,1,4),
	(7,1,5),
	(8,1,8);

/*!40000 ALTER TABLE `thumbs` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 types
# ------------------------------------------------------------

CREATE TABLE `types` (
  `id` int NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `types` WRITE;
/*!40000 ALTER TABLE `types` DISABLE KEYS */;

INSERT INTO `types` (`id`, `name`)
VALUES
	(1,'餐飲'),
	(2,'交通工具'),
	(3,'建築修繕'),
	(4,'休閒/運動'),
	(5,'旅遊/住宿'),
	(6,'電腦/3C'),
	(7,'流行/服飾');

/*!40000 ALTER TABLE `types` ENABLE KEYS */;
UNLOCK TABLES;


# 轉儲表 users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL,
  `fb_id` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `avatar` varchar(45) DEFAULT NULL,
  `is_admin` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `username`, `fb_id`, `email`, `password`, `avatar`, `is_admin`)
VALUES
	(1,'Davis',NULL,'davis@gmail.com','0659c7992e268962384eb17fafe88364',NULL,0),
	(7,'DOlala','',NULL,NULL,NULL,0),
	(8,'Wulala','',NULL,NULL,NULL,0),
	(10,'Chi Tai','10209887656465095','taichiayin@livemail.tw',NULL,NULL,0);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
