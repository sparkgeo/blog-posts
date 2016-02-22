import (
    "github.com/goamz/goamz/dynamodb"
)

// Get table description
func get_table_description(table *dynamodb.Table) *dynamodb.TableDescriptionT {
table_description, td_err := table.DescribeTable()

return table_description
}

// Get a table object to perform table.Operation(param1, ... param n)
func get_table(table_name string, ddbs *dynamodb.Server) *dynamodb.Table {
    table_descriptor, td_err := ddbs.DescribeTable(table_name)

    primary_key := build_primary_key(table_descriptor)

    table := ddbs.NewTable(table_name, primary_key)

    return table
}

// Create Dynamo DB Attribute list
func attribute_list_creator(attribute string) []dynamodb.Attribute {
    var attribute_list = make([]dynamodb.Attribute, 1)

    // cannot have empty attribute list
    attribute := &dynamodb.Attribute{
        Type:      "S",
        Name:      "attribute_name",
        Value:     attribute,
        SetValues: make([]string, 0),
        Exists:    ""}

    attribute_list[0] = *attribute

    return attribute_list
}

// Create Dynamo DB Key
func key_creator(primary_hash_key string) *dynamodb.Key {
    key := &dynamodb.Key{
        HashKey: msid}

    return key
}

// Write to database.
func database_writer(primary_hash_key string, primary_range_key string, attribute_list []dynamodb.Attribute, table dynamodb.Table) bool {
    bool, _ := table.PutItem(primary_hash_key, primary_range_key, attribute_list) // Overwrites

    return bool
}

// Read from Database (GetItem(key(primary_hash_key,primary_range_key)))
// Can be used for checking if record exists or not by the length of the attribute_list
// Consistent always returns the last updated value
func database_reader(table *dynamodb.Table, key *dynamodb.Key) (map[string]*dynamodb.Attribute, int) {

    attribute_map, _ := table.GetItemConsistent(key, true)

    return attribute_map, len(attribute_map)
}
