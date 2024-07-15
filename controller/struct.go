package controller

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gadget struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Merk        string             `bson:"merk,omitempty" json:"merk,omitempty"`
	Harga       float64            `bson:"harga,omitempty" json:"harga,omitempty"`
	Spesifikasi Spesifikasi        `bson:"spesifikasi,omitempty" json:"spesifikasi,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}

type Spesifikasi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Prosesor     string   `bson:"prosesor,omitempty" json:"prosesor,omitempty"`
	RAM          int      `bson:"ram,omitempty" json:"ram,omitempty"`
	Storage      int      `bson:"storage,omitempty" json:"storage,omitempty"`
	Kamera       string   `bson:"kamera,omitempty" json:"kamera,omitempty"`
	Baterai      string   `bson:"baterai,omitempty" json:"baterai,omitempty"`
	OS           string   `bson:"os,omitempty" json:"os,omitempty"`
	Layar        string   `bson:"layar,omitempty" json:"layar,omitempty"`
	FiturLainnya []string `bson:"fitur_lainnya,omitempty" json:"fitur_lainnya,omitempty"`
}

type Review struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	GadgetID primitive.ObjectID `bson:"gadget_id,omitempty" json:"gadget_id,omitempty"`
	Rating   int                `bson:"rating,omitempty" json:"rating,omitempty"`
	Review   string             `bson:"review,omitempty" json:"review,omitempty"`
	Datetime time.Time          `bson:"datetime,omitempty" json:"datetime,omitempty"` // Use time.Time instead of primitive.DateTime
}

type ReqGadget struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Merk        string             `bson:"merk,omitempty" json:"merk,omitempty"`
	Harga       float64            `bson:"harga,omitempty" json:"harga,omitempty"`
	Spesifikasi Spesifikasi        `bson:"spesifikasi,omitempty" json:"spesifikasi,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}

type ReqReview struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	GadgetID primitive.ObjectID `bson:"gadget_id,omitempty" json:"gadget_id,omitempty"`
	Rating   int                `bson:"rating,omitempty" json:"rating,omitempty"`
	Review   string             `bson:"review,omitempty" json:"review,omitempty"`
}
