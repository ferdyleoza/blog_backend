package controller

import (
	"Backend/config"
	"Backend/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllArtikels(ctx context.Context)([]model.Artikel, error) {
	collection := config.DB.Collection("artikel")
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllArtikels: %v\n", err)
		return nil, err
	}

	var data []model.Artikel
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Printf("GetAllArtikels: %v\n", err)
		return nil, err
	}
	return data, nil
}

func CreateArtikel(artikel *model.Artikel) error {
	collection := config.DB.Collection("artikel")

	// Buat ID otomatis
	artikel.ID = primitive.NewObjectID().Hex()

	_, err := collection.InsertOne(context.TODO(), artikel)
	if err != nil {
		fmt.Println("CreateArtikel:", err)
		return err
	}

	return nil
}

func GetArtikelByID(ctx context.Context, id string) (model.Artikel, error) {
	var artikel model.Artikel
	collection := config.DB.Collection("artikel")
	
	// Karena _id bertipe string
	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&artikel)
	if err != nil {
		return artikel, fmt.Errorf("data tidak ditemukan: %v", err)
	}

	return artikel, nil
}

// controller/artikel_controller.go
func UpdateArtikelByID(ctx context.Context, id string, data model.Artikel) error {
	collection := config.DB.Collection("artikels")
	
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"judul":       data.Judul,
			"isi":         data.Isi,
			"tanggal":     data.Tanggal,
			"id_penulis":  data.IDPenulis,
			"id_kategori": data.IDKategori,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArtikelByID(ctx context.Context, id string) error {
	collection := config.DB.Collection("artikels")

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}



