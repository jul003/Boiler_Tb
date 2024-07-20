package controller

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gadget struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty" example:"iPhone 13"`
	Merk        string             `bson:"merk,omitempty" json:"merk,omitempty" example:"Apple"`
	Harga       float64            `bson:"harga,omitempty" json:"harga,omitempty" example:"15000000"`
	Spesifikasi Spesifikasi        `bson:"spesifikasi,omitempty" json:"spesifikasi,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty" example:"Smartphone dengan layar Super Retina XDR 6.1 inci dan kamera ganda 12 MP."`
}

type Spesifikasi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"987654321"`
	Prosesor    string             `bson:"prosesor,omitempty" json:"prosesor,omitempty" example:"A15 Bionic"`
	RAM         int                `bson:"ram,omitempty" json:"ram,omitempty" example:"6"`
	Storage     int                `bson:"storage,omitempty" json:"storage,omitempty" example:"128"`
}

type Review struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"1122334455"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty" example:"2233445566"`
	GadgetID primitive.ObjectID `bson:"gadget_id,omitempty" json:"gadget_id,omitempty" example:"3344556677"`
	Rating   int                `bson:"rating,omitempty" json:"rating,omitempty" example:"5"`
	Review   string             `bson:"review,omitempty" json:"review,omitempty" example:"Luar biasa! Kinerja cepat dan kamera sangat bagus."`
	Datetime time.Time          `bson:"datetime,omitempty" json:"datetime,omitempty" example:"2024-07-20T00:00:00Z" format:"date-time"`
}

type ReqGadget struct {
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty" example:"iPhone 13"`
	Merk        string             `bson:"merk,omitempty" json:"merk,omitempty" example:"Apple"`
	Harga       float64            `bson:"harga,omitempty" json:"harga,omitempty" example:"15000000"`
	Spesifikasi ReqSpesifikasi     `bson:"spesifikasi,omitempty" json:"spesifikasi,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty" example:"Smartphone dengan layar Super Retina XDR 6.1 inci dan kamera ganda 12 MP."`
}

type ReqSpesifikasi struct {
	Prosesor    string             `bson:"prosesor,omitempty" json:"prosesor,omitempty" example:"A15 Bionic"`
	RAM         int                `bson:"ram,omitempty" json:"ram,omitempty" example:"6"`
	Storage     int                `bson:"storage,omitempty" json:"storage,omitempty" example:"128"`
}

type ReqReview struct {
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty" example:"2233445566"`
	GadgetID primitive.ObjectID `bson:"gadget_id,omitempty" json:"gadget_id,omitempty" example:"3344556677"`
	Rating   int                `bson:"rating,omitempty" json:"rating,omitempty" example:"5"`
	Review   string             `bson:"review,omitempty" json:"review,omitempty" example:"Luar biasa! Kinerja cepat dan kamera sangat bagus."`
}
