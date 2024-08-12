package main

import (
	"github.com/alwaysaashutosh/MongoDB-Golang/database"
	"go.mongodb.org/mongo-driver/bson"
)

var c database.Database

func main() {
	cfg := &database.DatabaseConfig{
		Driver:   "postgres",
		DbName:   "starter",
		Host:     "mongodb://localhost/?directConnection=true",
		Username: "",
		Password: "",
	}
	c = database.NewDBClient(cfg)

	switchvalues := struct {
		InsertSingleElement           bool
		InsertMultipleElement         bool
		InsertNestedElements          bool
		ReadAllElement                bool
		ReadWithWhereClause           bool
		ReadWithInOperator            bool
		ReadWithLtOperator            bool
		ReadWithOrOperator            bool
		ReadWithAndOperator           bool
		ReadWithRegexOperator         bool
		QueryNestedElementmatch       bool
		QueryNestedWithAnd            bool
		QueryOnEmbeededNestedDocument bool
		InsertArrayElements           bool
		QueryArrayElements            bool
		InsertEmbeededDocuments       bool
		QueryOnEmbeededDocuments      bool
		ProjectDocument               bool
	}{
		InsertSingleElement:           false,
		InsertMultipleElement:         false,
		InsertNestedElements:          false,
		ReadAllElement:                false,
		ReadWithWhereClause:           false,
		ReadWithInOperator:            false,
		ReadWithLtOperator:            false,
		ReadWithOrOperator:            false,
		ReadWithAndOperator:           false,
		ReadWithRegexOperator:         false,
		QueryNestedElementmatch:       false,
		QueryNestedWithAnd:            false,
		QueryOnEmbeededNestedDocument: false,
		InsertArrayElements:           false,
		QueryArrayElements:            false,
		InsertEmbeededDocuments:       false,
		QueryOnEmbeededDocuments:      false,
		ProjectDocument:               true,
	}

	if switchvalues.InsertSingleElement {
		element := &bson.D{
			{Key: "item", Value: "babu-gopal"},
			{Key: "qty", Value: 100},
			{Key: "tags", Value: bson.A{"cotton", "silk", "nylon", "rayon"}},
			{Key: "size", Value: bson.D{
				{Key: "h", Value: 28},
				{Key: "w", Value: 35.5},
				{Key: "uom", Value: "cm"},
			}},
			{Key: "some-Datatypes", Value: bson.M{"jai shree ram": "rajaram", "raghupati raghav": "sitaram"}},
		}
		c.InsertElement(element)
	}

	if switchvalues.InsertMultipleElement {
		elements := []interface{}{
			bson.D{
				{Key: "item", Value: "multi-canvas"},
				{Key: "qty", Value: 1},
				{Key: "tags", Value: bson.A{"cotton", "silk", "nylon", "rayon"}},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 28},
					{Key: "w", Value: 35.5},
					{Key: "uom", Value: "cm"},
				}},
				{Key: "some-Datatypes", Value: bson.M{"jai shree ram": "rajaram", "raghupati raghav": "sitaram"}},
			},
			bson.D{
				{Key: "item", Value: "multi-canvas"},
				{Key: "qty", Value: 2},
				{Key: "tags", Value: bson.A{"cotton", "silk", "nylon", "rayon"}},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 28},
					{Key: "w", Value: 35.5},
					{Key: "uom", Value: "cm"},
				}},
				{Key: "some-Datatypes", Value: bson.M{"jai shree ram": "rajaram", "raghupati raghav": "sitaram"}},
			},
			bson.D{
				{Key: "item", Value: "multi-canvas"},
				{Key: "qty", Value: 3},
				{Key: "tags", Value: bson.A{"cotton", "silk", "nylon", "rayon"}},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 28},
					{Key: "w", Value: 35.5},
					{Key: "uom", Value: "cm"},
				}},
				{Key: "some-Datatypes", Value: bson.M{"jai shree ram": "rajaram", "raghupati raghav": "sitaram"}},
			},
		}

		c.InsertElements(elements)
	}

	// To fetch all the elements .
	if switchvalues.ReadAllElement {
		condition := &bson.D{}
		c.ReadElements(condition, nil)
	}

	// find elements With some Condition .
	if switchvalues.ReadWithWhereClause {
		condition := &bson.D{{Key: "qty", Value: 2}}
		// condition = &bson.D{{"qty", 2}}
		c.ReadElements(condition, nil)
	}

	// $in operator in find query .
	if switchvalues.ReadWithInOperator {
		condition := &bson.D{
			{Key: "item", Value: bson.D{{Key: "$in", Value: bson.A{"canvas", "dirty-canvas"}}}},
		}
		c.ReadElements(condition, nil)
	}

	// with $lt and $gt operator .
	if switchvalues.ReadWithLtOperator {
		condition := &bson.D{
			{Key: "item", Value: "multi-canvas"},
			{Key: "qty", Value: bson.D{{Key: "$lt", Value: 100}}},
		}
		c.ReadElements(condition, nil)
	}

	if switchvalues.ReadWithOrOperator {
		condition := &bson.D{
			{Key: "item", Value: "multi-canvas"},
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "", Value: ""}},
				bson.D{{Key: "", Value: ""}},
			}},
		}
		c.ReadElements(condition, nil)
	}

	if switchvalues.ReadWithAndOperator {
		condition := &bson.D{
			{Key: "item", Value: "multi-canvas"},
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "", Value: ""}},
				bson.D{{Key: "", Value: ""}},
			}},
		}
		c.ReadElements(condition, nil)
	}

	if switchvalues.ReadWithRegexOperator {
		condition := &bson.D{
			{Key: "item", Value: bson.D{{Key: "$regex", Value: "^ba"}}},
		}
		c.ReadElements(condition, nil)
	}

	if switchvalues.InsertNestedElements {

		elements := []interface{}{
			bson.D{
				{Key: "item", Value: "journal"},
				{Key: "qty", Value: 25},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 14},
					{Key: "w", Value: 21},
					{Key: "uom", Value: "cm"},
				}},
				{Key: "status", Value: "A"},
			},
			bson.D{
				{Key: "item", Value: "notebook"},
				{Key: "qty", Value: 50},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 8.5},
					{Key: "w", Value: 11},
					{Key: "uom", Value: "in"},
				}},
				{Key: "status", Value: "A"},
			},
			bson.D{
				{Key: "item", Value: "paper"},
				{Key: "qty", Value: 100},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 8.5},
					{Key: "w", Value: 11},
					{Key: "uom", Value: "in"},
				}},
				{Key: "status", Value: "D"},
			},
			bson.D{
				{Key: "item", Value: "planner"},
				{Key: "qty", Value: 75},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 22.85},
					{Key: "w", Value: 30},
					{Key: "uom", Value: "cm"},
				}},
				{Key: "status", Value: "D"},
			},
			bson.D{
				{Key: "item", Value: "postcard"},
				{Key: "qty", Value: 45},
				{Key: "size", Value: bson.D{
					{Key: "h", Value: 10},
					{Key: "w", Value: 15.25},
					{Key: "uom", Value: "cm"},
				}},
				{Key: "status", Value: "A"},
			},
		}

		c.InsertElements(elements)

	}

	if switchvalues.QueryNestedElementmatch {
		condition := &bson.D{
			{Key: "size.h", Value: bson.D{
				{Key: "$lt", Value: 15},
			}},
		}
		c.ReadElements(condition, nil)
	}

	if switchvalues.QueryNestedWithAnd {
		condition := &bson.D{
			{Key: "status", Value: "A"},
			{Key: "size.uom", Value: "in"},
			{Key: "size.h", Value: bson.D{{
				Key: "$lt", Value: 10,
			}}},
		}
		c.ReadElements(condition, nil)
	}

	if switchvalues.QueryOnEmbeededNestedDocument {

		condition := &bson.D{
			{Key: "size", Value: bson.D{
				// The below embeeded query does the exact match , If we change the order of elements the results will get changed.
				{Key: "h", Value: 14},
				{Key: "w", Value: 21},
				{Key: "uom", Value: "cm"},
			}},
		}

		// In the below query the change in order won't make any effect .
		/*
			condition := &bson.D{
				{Key: "size.w", Value: 21},
				{Key: "size.h", Value: 14},
				{Key: "size.uom", Value: "cm"},
			}
		*/

		c.ReadElements(condition, nil)
	}

	if switchvalues.InsertArrayElements {

		elements := []interface{}{
			bson.D{
				{"item", "journal"},
				{"qty", 25},
				{"tags", bson.A{"blank", "red"}},
				{"dim_cm", bson.A{14, 27}},
			},
			bson.D{
				{"item", "notebook"},
				{"qty", 50},
				{"tags", bson.A{"red", "blank"}},
				{"dim_cm", bson.A{14, 21}},
			},
			bson.D{
				{"item", "paper"},
				{"qty", 100},
				{"tags", bson.A{"red", "blank", "plain"}},
				{"dim_cm", bson.A{14, 21}},
			},
			bson.D{
				{"item", "planner"},
				{"qty", 75},
				{"tags", bson.A{"blank", "red"}},
				{"dim_cm", bson.A{22.85, 30}},
			},
			bson.D{
				{"item", "postcard"},
				{"qty", 45},
				{"tags", bson.A{"blue"}},
				{"dim_cm", bson.A{10, 15.25}},
			},
		}
		c.InsertElements(elements)
	}
	if switchvalues.QueryArrayElements {
		//Here it will pick only those docs where the only element in the array is [red,blank] and also the order needs to be the same .
		// condition := &bson.D{
		// 	{Key: "tags", Value: bson.A{"red", "blank"}},
		// }

		/* Here it will find all the documents wherever it found ["red","blank"] in tags.*/
		// condition := &bson.D{
		// 	{Key: "tags", Value: bson.D{
		// 		{Key: "$all", Value: bson.A{"red", "blank"}},
		// 	}},
		// }

		// Here we are querying all the documents wherever red in tags.
		// condition := &bson.D{
		// 	{Key: "tags", Value: "red"},
		// }

		// Here we have to find every document wherever the dim in greater than 25.
		// condition := &bson.D{
		// 	{"dim_cm", bson.D{
		// 		{"$gt", 25},
		// 	}},
		// }

		// query with compound filter Condition . It filters out all the documents whichever is satisfysing the condition .
		// e.g., one element can satisfy the greater than 15 condition and another element can satisfy the less than 20 condition, or a single element can satisfy both:
		// condition := &bson.D{
		// 	{Key: "dim_cm", Value: bson.D{
		// 		{"$gt", 15},
		// 		{"$lt", 25},
		// 	}},
		// }

		// This cndition will make sure that any of the element in the array should match both the criteria .
		// condition := &bson.D{
		// 	{Key: "dim_cm", Value: bson.D{
		// 		{Key: "$elemMatch", Value: bson.D{
		// 			{Key: "$gt", Value: 15},
		// 			{Key: "$lt", Value: 25},
		// 		}},
		// 	}},
		// }

		/*------------------------------------------The above query uses not opeartor which i have to study next .-----------------------------------------------*/
		// condition := &bson.D{
		// 	{Key: "dim_cm", Value: bson.D{
		// 		{Key: "$exists", Value: true},
		// 		{Key: "$not", Value: bson.D{
		// 			{Key: "$elemMatch", Value: bson.D{
		// 				{Key: "$not", Value: bson.D{
		// 					{Key: "$gt", Value: 15},
		// 					{Key: "$lt", Value: 31},
		// 				}},
		// 			}},
		// 		}},
		// 	}},
		// }

		// Query on Array Index
		// condition := &bson.D{
		// 	{"dim_cm.1", bson.D{
		// 		{"$gt", 25},
		// 	}},
		// }

		// Query on array size . select document where tag size is 3
		condition := &bson.D{
			{"tags", bson.D{
				{"$size", 3},
			}},
		}

		c.ReadElements(condition, nil)
	}

	if switchvalues.InsertEmbeededDocuments {

		docs := []interface{}{
			bson.D{
				{"item", "journal"},
				{"instock", bson.A{
					bson.D{
						{"warehouse", "A"},
						{"qty", 5},
					},
					bson.D{
						{"warehouse", "C"},
						{"qty", 15},
					},
				}},
			},
			bson.D{
				{"item", "notebook"},
				{"instock", bson.A{
					bson.D{
						{"warehouse", "C"},
						{"qty", 5},
					},
				}},
			},
			bson.D{
				{"item", "paper"},
				{"instock", bson.A{
					bson.D{
						{"warehouse", "A"},
						{"qty", 60},
					},
					bson.D{
						{"warehouse", "B"},
						{"qty", 15},
					},
				}},
			},
			bson.D{
				{"item", "planner"},
				{"instock", bson.A{
					bson.D{
						{"warehouse", "A"},
						{"qty", 40},
					},
					bson.D{
						{"warehouse", "B"},
						{"qty", 5},
					},
				}},
			},
			bson.D{
				{"item", "postcard"},
				{"instock", bson.A{
					bson.D{
						{"warehouse", "B"},
						{"qty", 15},
					},
					bson.D{
						{"warehouse", "C"},
						{"qty", 35},
					},
				}},
			},
		}

		c.InsertElements(docs)

	}
	if switchvalues.QueryOnEmbeededDocuments {

		condition := &bson.D{
			{Key: "instock", Value: bson.D{
				{Key: "warehouse", Value: "B"},
				{Key: "qty", Value: 5},
			}},
		}

		// Specify Query condition on the field of array of documents.
		/*
			condition := &bson.D{
				{Key: "instock.qty", Value: bson.D{
					{Key: "$lt", Value: 10},
				}},
			}
		*/

		// Querying using Array Index.
		/*
			condition := &bson.D{
				{Key: "instock.0.qty", Value: bson.D{
					{Key: "$lt", Value: 10},
				}},
			}
		*/

		// Query using elemMatch Operator.
		/*
			condition = &bson.D{
				{Key: "instock", Value: bson.D{
					{Key: "$elemMatch", Value: bson.D{
						{Key: "warehouse", Value: "C"},
						{Key: "qty", Value: bson.D{
							{Key: "$lt", Value: 20},
						}},
					}},
				}},
			}
		*/

		/*
			condition := &bson.D{
				{"instock", bson.D{
					{"$elemMatch", bson.D{
						{"qty", bson.D{
							{"$gt", 10},
							{"$lte", 20},
						}},
					}},
				}},
			}
		*/

		/*
			condition := &bson.D{
				{"instock.qty", bson.D{
					{"$gt", 10},
					{"$lte", 20},
				}},
			}
		*/

		// Note : Element match will match the particular document whose individual elements satisfies the condition .
		// WIthout element match it will return that doument whose individual elements are satisfying diffrent conditions .

		c.ReadElements(condition, nil)
	}

	// Project Fields to Return from Query - PROJECTION

	if switchvalues.ProjectDocument {

		// docs := []interface{}{
		// 	bson.D{
		// 		{"item", "journal"},
		// 		{"status", "A"},
		// 		{"size", bson.D{
		// 			{"h", 14},
		// 			{"w", 21},
		// 			{"uom", "cm"},
		// 		}},
		// 		{"instock", bson.A{
		// 			bson.D{
		// 				{"warehouse", "A"},
		// 				{"qty", 5},
		// 			},
		// 		}},
		// 	},
		// 	bson.D{
		// 		{"item", "notebook"},
		// 		{"status", "A"},
		// 		{"size", bson.D{
		// 			{"h", 8.5},
		// 			{"w", 11},
		// 			{"uom", "in"},
		// 		}},
		// 		{"instock", bson.A{
		// 			bson.D{
		// 				{"warehouse", "EC"},
		// 				{"qty", 5},
		// 			},
		// 		}},
		// 	},
		// 	bson.D{
		// 		{"item", "paper"},
		// 		{"status", "D"},
		// 		{"size", bson.D{
		// 			{"h", 8.5},
		// 			{"w", 11},
		// 			{"uom", "in"},
		// 		}},
		// 		{"instock", bson.A{
		// 			bson.D{
		// 				{"warehouse", "A"},
		// 				{"qty", 60},
		// 			},
		// 		}},
		// 	},
		// 	bson.D{
		// 		{"item", "planner"},
		// 		{"status", "D"},
		// 		{"size", bson.D{
		// 			{"h", 22.85},
		// 			{"w", 30},
		// 			{"uom", "cm"},
		// 		}},
		// 		{"instock", bson.A{
		// 			bson.D{
		// 				{"warehouse", "A"},
		// 				{"qty", 40},
		// 			},
		// 		}},
		// 	},
		// 	bson.D{
		// 		{"item", "postcard"},
		// 		{"status", "A"},
		// 		{"size", bson.D{
		// 			{"h", 10},
		// 			{"w", 15.25},
		// 			{"uom", "cm"},
		// 		}},
		// 		{"instock", bson.A{
		// 			bson.D{
		// 				{"warehouse", "B"},
		// 				{"qty", 15},
		// 			},
		// 			bson.D{
		// 				{"warehouse", "EC"},
		// 				{"qty", 35},
		// 			},
		// 		}},
		// 	},
		// }

		// c.InsertElements(docs)

		projection := &bson.D{
			/*
				This will only include item and status as well as excluding the _id to the resultant .
			*/

			// {Key: "item", Value: 1},
			// {Key: "status", Value: 1},
			// {"_id", 0},

			/*
				This will exclue status and instock from the resultant
			*/
			// {"status", 0},
			// {"instock", 0},
		}

		condition := &bson.D{
			{Key: "status", Value: "A"},
		}

		c.ReadElements(condition, projection)

		/*
			Some useful operators are
			i)   $ne      - not eqaultity operator
			ii)  $type    - type operator to check the bson type . https://www.mongodb.com/docs/manual/reference/bson-types/#std-label-bson-types
			iii) $exists - used to check the existence of object .
			iv)
		*/

	}

}

// Tesing the db for first thing
// testing for second thimng
