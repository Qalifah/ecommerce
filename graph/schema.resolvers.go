package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/Qalifah/ecommerce/graph/generated"
	"github.com/Qalifah/ecommerce/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	userCollection := r.DB.Collection("users")
	if sr := userCollection.FindOne(ctx, bson.M{"username": input.Username}); sr.Err() == nil {
		return nil, errors.New("username already taken, choose a new one")
	}
	result, err := userCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	if err = userCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, username string, changes map[string]interface{}) (*model.User, error) {
	userCollection := r.DB.Collection("users")
	result, err := userCollection.UpdateOne(ctx, bson.M{"username": username}, changes)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	if err := userCollection.FindOne(ctx, bson.M{"_id": result.UpsertedID}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, username string) (string, error) {
	userCollection := r.DB.Collection("users")
	_, err := userCollection.DeleteOne(ctx, bson.M{"username": username})
	if err != nil {
		return "Couldn't delete user!", err
	}
	return "User deleted Successfully!", nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CategoryInput) (*model.Category, error) {
	categoryCollection := r.DB.Collection("categories")
	if sr := categoryCollection.FindOne(ctx, bson.M{"name": input.Name}); sr.Err() == nil {
		return nil, errors.New("category already created")
	}
	_, err := categoryCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	category := &model.Category{}
	if err := categoryCollection.FindOne(ctx, bson.M{"name": input.Name}).Decode(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, name string, changes map[string]interface{}) (*model.Category, error) {
	categoryCollection := r.DB.Collection("categories")
	_, err := categoryCollection.UpdateOne(ctx, bson.M{"name": name}, changes)
	if err != nil {
		return nil, err
	}
	category := &model.Category{}
	if err := categoryCollection.FindOne(ctx, bson.M{"name": name}).Decode(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, name string) (string, error) {
	categoryCollection := r.DB.Collection("categories")
	if _, err := categoryCollection.DeleteOne(ctx, bson.M{"name": name}); err != nil {
		return "Couldn't delete category!", err
	}
	return "Category deleted successfully!", nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*model.Product, error) {
	productCollection := r.DB.Collection("products")
	if rs := productCollection.FindOne(ctx, bson.M{"categoryID": input.CategoryID, "name": input.Name, "price": input.Price, "brand": input.Brand, "sellerUsername": input.SellerUsername}); rs == nil {
		return nil, errors.New("product already exist")
	}
	result, err := productCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	product := &model.Product{}
	if err := productCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, changes map[string]interface{}) (*model.Product, error) {
	productCollection := r.DB.Collection("products")
	result, err := productCollection.UpdateOne(ctx, bson.M{"_id": id}, changes)
	if err != nil {
		return nil, err
	}
	product := &model.Product{}
	if err := productCollection.FindOne(ctx, bson.M{"_id": result.UpsertedID}).Decode(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (string, error) {
	productCollection := r.DB.Collection("products")
	if _, err := productCollection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return "Couldn't delete product!", err
	}
	return "Product deleted successfully!", nil
}

func (r *productResolver) Category(ctx context.Context, obj *model.Product) (*model.Category, error) {
	categoryCollection := r.DB.Collection("categories")
	category := &model.Category{}
	if err := categoryCollection.FindOne(ctx, bson.M{"name": obj.CategoryName}).Decode(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *productResolver) Seller(ctx context.Context, obj *model.Product) (*model.User, error) {
	userCollection := r.DB.Collection("users")
	user := &model.User{}
	if err := userCollection.FindOne(ctx, bson.M{"username": obj.SellerUsername}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Category(ctx context.Context, name string) (*model.Category, error) {
	categoryCollection := r.DB.Collection("categories")
	category := &model.Category{}
	if err := categoryCollection.FindOne(ctx, bson.M{"name": name}).Decode(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categoryCollection := r.DB.Collection("categories")
	categories := []*model.Category{}
	cusor, err := categoryCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cusor.All(ctx, categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	productCollection := r.DB.Collection("products")
	product := &model.Product{}
	if err := productCollection.FindOne(ctx, bson.M{"_id":id}).Decode(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (r *queryResolver) ProductsInCategory(ctx context.Context, categoryName string) ([]*model.Product, error) {
	productCollection := r.DB.Collection("products")
	products := []*model.Product{}
	cusor, err := productCollection.Find(ctx, bson.M{"categoryName" : categoryName})
	if err != nil {
		return nil, err
	}
	defer cusor.Close(ctx)
	for cusor.Next(ctx) {
		product := &model.Product{}
		cusor.Decode(product)
		products = append(products, product)
	}
	if cusor.Err() != nil {
		return nil, cusor.Err()
	}
	return products, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	productCollection := r.DB.Collection("products")
	products := []*model.Product{}
	cusor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cusor.Close(ctx)
	for cusor.Next(ctx) {
		product := &model.Product{}
		cusor.Decode(product)
		products = append(products, product)
	}
	if cusor.Err() != nil {
		return nil, cusor.Err()
	}
	return products, nil
}

func (r *queryResolver) User(ctx context.Context, username string) (*model.User, error) {
	userCollection := r.DB.Collection("users")
	user := &model.User{}
	if err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	usersCollection := r.DB.Collection("users")
	users := []*model.User{}
	cusor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cusor.Close(ctx)
	for cusor.Next(ctx) {
		user := &model.User{}
		cusor.Decode(user)
		users = append(users, user)
	}
	if cusor.Err() != nil {
		return nil, cusor.Err()
	}
	return users, nil
}

func (r *queryResolver) SellerProducts(ctx context.Context, sellerUsername string) ([]*model.Product, error) {
	productCollection := r.DB.Collection("products")
	products := []*model.Product{}
	cusor, err := productCollection.Find(ctx, bson.M{"sellerUsername": sellerUsername})
	if err != nil {
		return nil, err
	}
	defer cusor.Close(ctx)
	for cusor.Next(ctx) {
		product := &model.Product{}
		cusor.Decode(product)
		products = append(products, product)
	}
	if cusor.Err() != nil {
		return nil, cusor.Err()
	}
	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
