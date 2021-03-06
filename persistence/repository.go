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
	result, err := repo.collection().InsertOne(context.TODO(), r)
	if err != nil {
		log.Println("Valid to insert entry: " + r.String() + " Reason: " + err.Error())
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	r.UID = oid
	return &r
}

func (repo ReminderRepository) DeleteReminder(UID primitive.ObjectID) (bool, int64) {
	dr, err := repo.collection().DeleteOne(context.TODO(), bson.D{{"_id", UID}})
	if err != nil {
		log.Println("Failed to delete entry with id=" + UID.String() + " Reason: " + err.Error())
		return false, 0
	}
	return true, dr.DeletedCount
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
	return repo.findAllInternal(bson.D{{}})
}

func (repo *ReminderRepository) FindCompleted() []Reminder {
	filter := bson.D{
		{"completed", true},
	}
	return repo.findAllInternal(filter)
}

func (repo *ReminderRepository) FindUncompleted() []Reminder {
	filter := bson.D{
		{"completed", false},
	}
	return repo.findAllInternal(filter)
}

func (repo *ReminderRepository) collection() *mongo.Collection {
	return repo.DB.Collection("reminder")
}

func (repo *ReminderRepository) findAllInternal(filter interface{}) []Reminder {
	reminders := make([]Reminder, 0)
	cursor, err := repo.collection().Find(context.TODO(), filter)
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
