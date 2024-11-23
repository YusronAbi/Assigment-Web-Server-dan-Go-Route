INSERT INTO Produk (nama, deskripsi, harga, kategori)
VALUES
('Kaos Oversize', 'Bahan adem dan tidak mudah kusut', 105000000.00, 'Baju'),
('Kemeja', 'Bahan adem dan tidak mudah kusut', 150000.00, 'Baju'),
('Polo Tshirt', 'Bahan adem dan tidak mudah kusut', 80000000.00, 'Baju');


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
