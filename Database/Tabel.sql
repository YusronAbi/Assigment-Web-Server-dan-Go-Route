INSERT INTO Produk (nama, deskripsi, harga, kategori)
VALUES
('Laptop', 'Laptop gaming dengan spesifikasi tinggi', 15000000.00, 'Elektronik'),
('Kursi', 'Kursi ergonomis untuk kantor', 1200000.00, 'Furniture'),
('Meja', 'Meja kayu solid untuk ruang kerja', 3000000.00, 'Furniture');


INSERT INTO Inventaris (id_produk, jumlah, lokasi)
VALUES
(1, 10, 'Gudang A'),
(2, 50, 'Gudang B'),
(3, 20, 'Gudang A');

INSERT INTO Pesanan (id_produk, jumlah)
VALUES
(1, 2),
(2, 5),
(3, 1);
