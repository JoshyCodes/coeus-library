package db

import (
   "context"
   "os"

	"github.com/joshyCodes/coeus-library/data"

   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(user, pwd, url, table string) (*mongo.Database) {

	uri := fmt.Sprintf("mongodb://%s:%s@%s", user, pwd, url)
	connection := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), connection)
	if err != nil {
		log.Println("Database failed to connect.")
		panic(err)
		os.Exit(1)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	mgdb := client.Database(table)

	cur, _ := mgdb.ListCollectionNames(context.TODO(), bson.M{})

	var Collections []string
	for _, document := range cur {
		Collections = append(Collections, document)
	}

	if !data.IfArrayContainString(Collections, "users") {
		collection := mgdb.Collection("users")
		_, err = collection.InsertOne(context.TODO(), bson.D{
			{Key: "_id", Value: "admin@admin.com"},
			{Key: "password", Value: "$2a$14$2c36B8/AU/Aji.nyjqpaGe.d8M7EAd04Pm6IExKG.mxRha3Jtbk.y"},
			{Key: "status", Value: "1"},
			{Key: "roles", Value: "super-admin"},
		})
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("Users collection has been created with default admin user.")
	}
	if !data.IfArrayContainString(Collections, "roles") {
		collection := mgdb.Collection("roles")
		_, err := collection.InsertOne(context.TODO(), bson.D{
			{Key: "_id", Value: "super-admin"},
			{Key: "permissions", Value: bson.A{"super-admin"}},
		})
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("Roles collection has been created with default Super-Admin Role.")
	}
	if !data.IfArrayContainString(Collections, "collections") {
		collection := mgdb.Collection("collections")
		_, err := collection.InsertMany(context.TODO(), []interface{}{
			bson.D{{Key: "_id", Value: "users"},{Key: "status", Value: "admin"}},
			bson.D{{Key: "_id", Value: "permissions"},{Key: "status", Value: "admin"}},
			bson.D{{Key: "_id", Value: "collections"},{Key: "status", Value: "admin"}},
			bson.D{{Key: "_id", Value: "Roles"},{Key: "status", Value: "admin"}},
		})
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("Collections collection has been created with default Collections List.")
	}
	if !data.IfArrayContainString(Collections, "permissions") {
		collection := mgdb.Collection("permissions")
		_, err := collection.InsertMany(context.TODO(), []interface{}{
			bson.D{{Key: "_id", Value: "admin-users"},{Key: "rules", Value: bson.A{ []string{"Veiw", "Veiw One", "Create New", "Edit", "Delete"}}}},
			bson.D{{Key: "_id", Value: "admin-roles"},{Key: "rules", Value: bson.A{ []string{"Veiw", "Veiw One", "Create New", "Edit", "Delete"}}}},
			bson.D{{Key: "_id", Value: "admin-api-keys"},{Key: "rules", Value: bson.A{ []string{"Veiw", "Veiw One", "Create New", "Edit", "Delete"}}}},
			bson.D{{Key: "_id", Value: "admin-collections"},{Key: "rules", Value: bson.A{ []string{"Veiw", "Veiw One", "Create New", "Edit", "Delete"}}}},
		})

		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("Roles collection has been created with default Super-Admin Role.")
	}

	log.Println("Mongodb database connected.")

	return mgdb
}