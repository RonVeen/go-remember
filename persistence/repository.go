package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ReminderRepository struct {
	DB *mongo.Database
}

type Reminder struct {
	UID       primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Comment   string             `bson:"comment,omitempty"`
	Completed bool               `bson:"completed"`
}

func (r *Reminder) String() string {
	return fmt.Sprintf("UID=%s, Title=%s, Comment=%s, Completed=%t", r.UID.String(), r.Title, r.Comment, r.Completed)
}

func (repo *ReminderRepository) AddReminder(r Reminder) *Reminder {
	_, err := repo.collection().InsertOne(context.TODO(), r)
	if err != nil {
		log.Println("Valid to insert entry: " + r.String() + " Reason: " + err.Error())
	}
	return &Reminder{
		Title:     r.Title,
		Comment:   r.Comment,
		Completed: r.Completed,
	}
	return nil
}

func (repo ReminderRepository) DeleteReminder(UID primitive.ObjectID) bool {
	_, err := repo.collection().DeleteOne(context.TODO(), bson.D{{"_id", UID}})
	if err != nil {
		log.Println("Failed to delete entry with id=" + UID.String() + " Reason: " + err.Error())
		return false
	}
	return true
}

func (repo *ReminderRepository) UpdateReminder(r Reminder) *Reminder {
	update := bson.D{
		{"$set", bson.D{
			{"title", r.Title},
			{"comment", r.Comment},
			{"completed", r.Completed},
		},
		},
	}
	_, err := repo.collection().UpdateOne(context.TODO(),
		bson.D{{"_id", r.UID}},
		update,
		options.Update().SetUpsert(true),
	)

	if err != nil {
		log.Println("Failed to update entry with id=" + r.UID.String() + " Reason: " + err.Error())
		return nil
	}
	return &r
}

func (repo *ReminderRepository) FindReminder(UID primitive.ObjectID) *Reminder {
	cursor, err := repo.collection().Find(context.TODO(), bson.D{{"_id", UID}})
	defer cursor.Close(context.TODO())

	if err != nil {
		log.Println("Failed to retrieve entry with id=" + UID.String() + " Reason: " + err.Error())
		return nil
	}

	if cursor.Next(context.TODO()) {
		var r Reminder
		if err := cursor.Decode(&r); err != nil {
			return nil
		}
		return &r
	}

	return nil
}

func (repo *ReminderRepository) FindAllReminder() []Reminder {
	var reminders []Reminder
	cursor, err := repo.collection().Find(context.TODO(), bson.D{{}})
	defer cursor.Close(context.TODO())

	if err != nil {
		log.Println("Failed to all entries, reason: " + err.Error())
		return nil
	}

	for cursor.Next(context.TODO()) {
		var r Reminder
		if err := cursor.Decode(&r); err != nil {
			return nil
		}
		reminders = append(reminders, r)
	}
	return reminders
}

func (repo *ReminderRepository) collection() *mongo.Collection {
	return repo.DB.Collection("reminder")
}
