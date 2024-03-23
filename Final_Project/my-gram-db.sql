-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 22 Mar 2024 pada 23.08
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `my-gram-db`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_comments`
--

CREATE TABLE `tb_comments` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `message` text NOT NULL,
  `photo_id` bigint(20) UNSIGNED DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_photo`
--

CREATE TABLE `tb_photo` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` longtext NOT NULL,
  `caption` longtext DEFAULT NULL,
  `photo_url` longtext NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tb_photo`
--

INSERT INTO `tb_photo` (`id`, `created_at`, `updated_at`, `title`, `caption`, `photo_url`, `user_id`) VALUES
(16, '2024-03-22 13:27:45.207', '2024-03-22 13:27:45.207', 'Photo Kucing', 'Ini Foto Kucing Lucu banget', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.halodoc.com%2Fartikel%2Fketahui-7-fakta-menarik-tentang-kucing-bengal&psig=AOvVaw3N3uIfov2XjSwxlGGEEv6I&ust=1711099624230000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCLDBq_2EhYUDFQAAAAAdAAAAABAE', 18),
(17, '2024-03-22 13:29:35.193', '2024-03-22 13:29:35.193', 'Kucing Raja', 'Kucingnya jadi raja', 'https://www.facebook.com/PcintaKucingImut/', 18);

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_social_media`
--

CREATE TABLE `tb_social_media` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `social_media_url` varchar(255) NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_users`
--

CREATE TABLE `tb_users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `username` longtext NOT NULL,
  `email` longtext NOT NULL,
  `password` longtext NOT NULL,
  `age` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tb_users`
--

INSERT INTO `tb_users` (`id`, `created_at`, `updated_at`, `username`, `email`, `password`, `age`) VALUES
(18, '2024-03-22 13:23:34.233', '2024-03-22 13:27:14.843', 'Budi Saja', 'budi@mail.com', '$2a$08$53GgfhEH98lzDjLBj0gZ5.5Mj6nrKTsxLj3G5StVOcqXu7q4/mmh.', 30);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `tb_comments`
--
ALTER TABLE `tb_comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_tb_photos_comments` (`photo_id`),
  ADD KEY `fk_tb_users_comments` (`user_id`);

--
-- Indeks untuk tabel `tb_photo`
--
ALTER TABLE `tb_photo`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_tb_users_photos` (`user_id`);

--
-- Indeks untuk tabel `tb_social_media`
--
ALTER TABLE `tb_social_media`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_tb_users_social_media` (`user_id`);

--
-- Indeks untuk tabel `tb_users`
--
ALTER TABLE `tb_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_tb_users_username` (`username`) USING HASH,
  ADD UNIQUE KEY `idx_tb_users_email` (`email`) USING HASH;

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `tb_comments`
--
ALTER TABLE `tb_comments`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `tb_photo`
--
ALTER TABLE `tb_photo`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT untuk tabel `tb_social_media`
--
ALTER TABLE `tb_social_media`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `tb_users`
--
ALTER TABLE `tb_users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `tb_comments`
--
ALTER TABLE `tb_comments`
  ADD CONSTRAINT `fk_tb_photos_comments` FOREIGN KEY (`photo_id`) REFERENCES `tb_photo` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_tb_users_comments` FOREIGN KEY (`user_id`) REFERENCES `tb_users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `tb_photo`
--
ALTER TABLE `tb_photo`
  ADD CONSTRAINT `fk_tb_users_photos` FOREIGN KEY (`user_id`) REFERENCES `tb_users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `tb_social_media`
--
ALTER TABLE `tb_social_media`
  ADD CONSTRAINT `fk_tb_users_social_media` FOREIGN KEY (`user_id`) REFERENCES `tb_users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
