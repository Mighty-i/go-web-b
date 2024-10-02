-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 02, 2024 at 08:06 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gowebpro`
--

-- --------------------------------------------------------

--
-- Table structure for table `todos`
--

CREATE TABLE `todos` (
  `id` int(11) NOT NULL,
  `title` text NOT NULL,
  `done` tinyint(1) DEFAULT 0,
  `name_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `todos`
--

INSERT INTO `todos` (`id`, `title`, `done`, `name_id`) VALUES
(13, 'test2', 1, 1),
(14, 'todoedit', 0, 1),
(15, 'เพิ่มแก้ไข', 0, 1),
(16, 'OPppp', 0, 9),
(17, 'todo', 0, 9),
(18, 'พรุ่งตื่น\n ทำงานโปรเจค', 0, 1),
(19, 'asdas', 0, 9);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `google_id` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `profile_image` varchar(255) DEFAULT NULL,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `google_id`, `email`, `name`, `profile_image`, `username`, `password`, `created_at`) VALUES
(1, 'ya29.a0AcM612zI1ouzErvELqcyfMzr0wzhU2fMmBTixC5kjZleA43gJ7f8nASEQS5Qh--Py4bxp_HedULHFznrH5-UL1x_8nOTDxo4FxhlnvtvJEY-x16NfYMmxnIK1yHKIQy2qGKbEfA-3SFy-CNEk5qGo5pxkntQ8PKXSue0FGggaCgYKAUMSARASFQHGX2MicqMgb7_awbJ4FS4548cvwQ0175', 'kotniya@gmail.com', 'Mighty', 'https://lh3.googleusercontent.com/a/ACg8ocIqXF_LUEqcvMa4qTNw5LvaMXd4ZiYIxRnbwPKUjJ_k4aSkrcTa=s96-c', NULL, NULL, '2024-09-30 04:07:18'),
(9, NULL, 'OP@gmail.com', 'OP', NULL, 'OP', '$2a$10$WLwWzUeSG9lPgzyzrnEr3.1cBsFFqqx4YYKroNTdAXA7oYVTm/Pw.', '2024-09-30 19:25:21'),
(10, NULL, 'user1@gmail.com', 'UI', NULL, 'user1', '$2a$10$VBFKJC0eWzaNWSvWHZjnLuMwrGOlb7tp04Z80ec.LEJXtKk2Q1No2', '2024-09-30 21:23:48');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `todos`
--
ALTER TABLE `todos`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_name_id` (`name_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `google_id` (`google_id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `todos`
--
ALTER TABLE `todos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `todos`
--
ALTER TABLE `todos`
  ADD CONSTRAINT `fk_name_id` FOREIGN KEY (`name_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
