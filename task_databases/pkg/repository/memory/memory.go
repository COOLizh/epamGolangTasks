package memory

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/COOLizh/epam/epamGolangTasks/hw8/task_databases/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactsRepositoryMongo struct {
	db *mongo.Database
}

func NewContactsRepositoryMongo(database *mongo.Database) *ContactsRepositoryMongo {
	return &ContactsRepositoryMongo{db: database}
}

func (r *ContactsRepositoryMongo) Save(contact model.Contact) (model.Contact, error) {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var foundResult bson.M
	collection.FindOne(ctx, bson.D{{Key: "Email", Value: contact.Email}}).Decode(&foundResult)
	if foundResult["Email"] == contact.Email {
		return model.Contact{}, fmt.Errorf("contact with email %q already exists", contact.Email)
	}
	collection.FindOne(ctx, bson.D{{Key: "Phone", Value: contact.Phone}}).Decode(&foundResult)
	if foundResult["Phone"] == contact.Phone {
		return model.Contact{}, fmt.Errorf("contact with phone %q already exists", contact.Phone)
	}
	ID, _ := collection.CountDocuments(context.Background(), bson.D{})
	ID++
	_, err := collection.InsertOne(ctx, bson.M{
		"ID":        uint(ID),
		"FirstName": contact.FirstName,
		"LastName":  contact.LastName,
		"Phone":     contact.Phone,
		"Email":     contact.Email,
	})
	if err != nil {
		return model.Contact{}, err
	}
	return contact, nil
}

func (r *ContactsRepositoryMongo) ListAll() ([]model.Contact, error) {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []model.Contact{}, err
	}
	defer cur.Close(ctx)
	contacts := make([]model.Contact, 0)
	for cur.Next(ctx) {
		var result bson.M
		err = cur.Decode(&result)
		if err != nil {
			return []model.Contact{}, err
		}
		contacts = append(contacts, model.Contact{
			ID:        uint(result["ID"].(int64)),
			FirstName: result["FirstName"].(string),
			LastName:  result["LastName"].(string),
			Phone:     result["Phone"].(string),
			Email:     result["Email"].(string),
		})
	}
	return contacts, nil
}

func (r *ContactsRepositoryMongo) GetByID(id uint) (model.Contact, error) {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var foundResult bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "ID", Value: id}}).Decode(&foundResult)
	if err == nil {
		return model.Contact{
			ID:        id,
			FirstName: foundResult["FirstName"].(string),
			LastName:  foundResult["LastName"].(string),
			Phone:     foundResult["Phone"].(string),
			Email:     foundResult["Email"].(string),
		}, nil
	}
	return model.Contact{}, fmt.Errorf("record not found")
}

func (r *ContactsRepositoryMongo) GetByPhone(phone string) (model.Contact, error) {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var foundResult bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "Phone", Value: phone}}).Decode(&foundResult)
	if err == nil {
		return model.Contact{
			ID:        uint(foundResult["ID"].(int64)),
			FirstName: foundResult["FirstName"].(string),
			LastName:  foundResult["LastName"].(string),
			Phone:     phone,
			Email:     foundResult["Email"].(string),
		}, nil
	}
	return model.Contact{}, fmt.Errorf("record not found")
}

func (r *ContactsRepositoryMongo) GetByEmail(email string) (model.Contact, error) {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var foundResult bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "Email", Value: email}}).Decode(&foundResult)
	if err == nil {
		return model.Contact{
			ID:        uint(foundResult["ID"].(int64)),
			FirstName: foundResult["FirstName"].(string),
			LastName:  foundResult["LastName"].(string),
			Phone:     foundResult["Phone"].(string),
			Email:     email,
		}, nil
	}
	return model.Contact{}, fmt.Errorf("record not found")
}

func (r *ContactsRepositoryMongo) SearchByName(n string) ([]model.Contact, error) {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []model.Contact{}, err
	}
	defer cur.Close(ctx)
	contacts := make([]model.Contact, 0)
	for cur.Next(ctx) {
		var result bson.M
		err = cur.Decode(&result)
		if err != nil {
			return []model.Contact{}, err
		}
		if strings.HasPrefix(result["FirstName"].(string), n) || strings.HasPrefix(result["LastName"].(string), n) {
			contacts = append(contacts, model.Contact{
				ID:        uint(result["ID"].(int64)),
				FirstName: result["FirstName"].(string),
				LastName:  result["LastName"].(string),
				Phone:     result["Phone"].(string),
				Email:     result["Email"].(string),
			})
		}
	}
	return contacts, nil
}

func (r *ContactsRepositoryMongo) Delete(id uint) error {
	collection := r.db.Collection("Contacts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.DeleteOne(ctx, bson.D{{Key: "ID", Value: id}})
	return err
}
