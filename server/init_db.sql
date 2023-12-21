-- Create Users Table
CREATE TABLE Users (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    Username VARCHAR(255),
    Password VARCHAR(255),
    Name VARCHAR(255),
    NIP VARCHAR(255),
    Golongan VARCHAR(255),
    Role VARCHAR(255)
);

-- Create TahunAjar Table
CREATE TABLE TahunAjar (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    Tahun VARCHAR(255)
);

-- Create Mapel Table
CREATE TABLE Mapel (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    Mapel VARCHAR(255)
);

-- Create Kelas Table
CREATE TABLE Kelas (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    Kelas VARCHAR(255),
    Jurusan VARCHAR(255)
);

-- Create Capaian Table
CREATE TABLE Capaian (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    JudulCapaian VARCHAR(255),
    UserID UUID,
    MapelID UUID,
    KelasID UUID,
    TahunAjarID UUID,
    JudulElemen VARCHAR(255),
    KetElemen VARCHAR(255),
    KetProsesMengamati VARCHAR(255),
    KetProsesMempertanyakan VARCHAR(255),
    KetProsesMerencanakan VARCHAR(255),
    KetProsesMemproses VARCHAR(255),
    KetProsesMengevaluasi VARCHAR(255),
    KetProsesMengkomunikasikan VARCHAR(255),
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    FOREIGN KEY (MapelID) REFERENCES Mapel(ID),
    FOREIGN KEY (KelasID) REFERENCES Kelas(ID),
    FOREIGN KEY (TahunAjarID) REFERENCES TahunAjar(ID)
);

-- Create AlurTP Table
CREATE TABLE AlurTP (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    Elemen VARCHAR(255),
    LingkupMateri VARCHAR(255),
    TujuanPembelajaran VARCHAR(255),
    KodeTP VARCHAR(255),
    AlokasiWaktu VARCHAR(255),
    SumberBelajar VARCHAR(255),
    ProjekPPancasila VARCHAR(255)
);

-- Create ModulAjar Table
CREATE TABLE ModulAjar (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UserID UUID,
    TahunAjarID UUID,
    KelasID UUID,
    Sekolah VARCHAR(255),
    AlokasiWaktu VARCHAR(255),
    KompetensiAwal VARCHAR(255),
    ProjekPPancasila VARCHAR(255),
    SaranaPrasarana VARCHAR(255),
    TargetPesertaDidik VARCHAR(255),
    ModelPembelajaran VARCHAR(255),
    TujuanPembelajaran VARCHAR(255),
    PemahamanBermakna VARCHAR(255),
    PertanyaanPemantik VARCHAR(255),
    KegiatanPembelajaran VARCHAR(255),
    Asesnmen VARCHAR(255),
    PengayaanRemedial VARCHAR(255),
    Refleksi VARCHAR(255),
    Glosarium VARCHAR(255),
    DaftarPustaka VARCHAR(255),
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    FOREIGN KEY (TahunAjarID) REFERENCES TahunAjar(ID),
    FOREIGN KEY (KelasID) REFERENCES Kelas(ID)
);

-- Create Prp Table
CREATE TABLE Prp (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    AsesemenNonKognitif VARCHAR(255),
    AsesemenKognitif VARCHAR(255),
    AsesemenFormatif VARCHAR(255),
    AsesemenSumatif VARCHAR(255),
    Pengayaan VARCHAR(255),
    Remedial VARCHAR(255)
);

-- Create Soal Table
CREATE TABLE Soal (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    KelasID UUID,
    Hari VARCHAR(255),
    Tanggal VARCHAR(255),
    Waktu VARCHAR(255),
    FOREIGN KEY (KelasID) REFERENCES Kelas(ID)
);

-- Create BankSoal Table
CREATE TABLE BankSoal (
    ID UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    Soal VARCHAR(255),
    OptionA VARCHAR(255),
    OptionB VARCHAR(255),
    OptionC VARCHAR(255),
    OptionD VARCHAR(255),
    OptionE VARCHAR(255),
    KunciJawaban VARCHAR(255),
    Materi VARCHAR(255),
    Indikator VARCHAR(255),
    TingkatKesukaran VARCHAR(255),
    FaktorKognitif VARCHAR(255)
);
