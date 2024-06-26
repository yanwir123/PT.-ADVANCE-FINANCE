create database AdvanceFinance;

CREATE TABLE Keuangan (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Bulan VARCHAR(255),
    Product VARCHAR(300),
    Masuk FLOAT,
    Keluar FLOAT,
    Stok FLOAT,
    Total DECIMAL(15,2),
    Keterangan VARCHAR(255)
);
