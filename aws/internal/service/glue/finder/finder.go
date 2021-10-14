package finder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	tfglue "github.com/hashicorp/terraform-provider-aws/aws/internal/service/glue"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/tfresource"
)

func DevEndpointByName(conn *glue.Glue, name string) (*glue.DevEndpoint, error) {
	input := &glue.GetDevEndpointInput{
		EndpointName: aws.String(name),
	}

	output, err := conn.GetDevEndpoint(input)

	if tfawserr.ErrCodeEquals(err, glue.ErrCodeEntityNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.DevEndpoint == nil {
		return nil, &resource.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return output.DevEndpoint, nil
}

// TableByName returns the Table corresponding to the specified name.
func TableByName(conn *glue.Glue, catalogID, dbName, name string) (*glue.GetTableOutput, error) {
	input := &glue.GetTableInput{
		CatalogId:    aws.String(catalogID),
		DatabaseName: aws.String(dbName),
		Name:         aws.String(name),
	}

	output, err := conn.GetTable(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// TriggerByName returns the Trigger corresponding to the specified name.
func TriggerByName(conn *glue.Glue, name string) (*glue.GetTriggerOutput, error) {
	input := &glue.GetTriggerInput{
		Name: aws.String(name),
	}

	output, err := conn.GetTrigger(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// RegistryByID returns the Registry corresponding to the specified ID.
func RegistryByID(conn *glue.Glue, id string) (*glue.GetRegistryOutput, error) {
	input := &glue.GetRegistryInput{
		RegistryId: tfglue.CreateAwsGlueRegistryID(id),
	}

	output, err := conn.GetRegistry(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// SchemaByID returns the Schema corresponding to the specified ID.
func SchemaByID(conn *glue.Glue, id string) (*glue.GetSchemaOutput, error) {
	input := &glue.GetSchemaInput{
		SchemaId: tfglue.CreateAwsGlueSchemaID(id),
	}

	output, err := conn.GetSchema(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// SchemaVersionByID returns the Schema corresponding to the specified ID.
func SchemaVersionByID(conn *glue.Glue, id string) (*glue.GetSchemaVersionOutput, error) {
	input := &glue.GetSchemaVersionInput{
		SchemaId: tfglue.CreateAwsGlueSchemaID(id),
		SchemaVersionNumber: &glue.SchemaVersionNumber{
			LatestVersion: aws.Bool(true),
		},
	}

	output, err := conn.GetSchemaVersion(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// PartitionByValues returns the Partition corresponding to the specified Partition Values.
func PartitionByValues(conn *glue.Glue, id string) (*glue.Partition, error) {

	catalogID, dbName, tableName, values, err := tfglue.ReadAwsGluePartitionID(id)
	if err != nil {
		return nil, err
	}

	input := &glue.GetPartitionInput{
		CatalogId:       aws.String(catalogID),
		DatabaseName:    aws.String(dbName),
		TableName:       aws.String(tableName),
		PartitionValues: aws.StringSlice(values),
	}

	output, err := conn.GetPartition(input)
	if err != nil {
		return nil, err
	}

	if output == nil && output.Partition == nil {
		return nil, nil
	}

	return output.Partition, nil
}

// ConnectionByName returns the Connection corresponding to the specified Name and CatalogId.
func ConnectionByName(conn *glue.Glue, name, catalogID string) (*glue.Connection, error) {
	input := &glue.GetConnectionInput{
		CatalogId: aws.String(catalogID),
		Name:      aws.String(name),
	}

	output, err := conn.GetConnection(input)
	if tfawserr.ErrCodeEquals(err, glue.ErrCodeEntityNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.Connection == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.Connection, nil
}

// PartitionIndexByName returns the Partition Index corresponding to the specified Partition Index Name.
func PartitionIndexByName(conn *glue.Glue, id string) (*glue.PartitionIndexDescriptor, error) {

	catalogID, dbName, tableName, partIndex, err := tfglue.ReadAwsGluePartitionIndexID(id)
	if err != nil {
		return nil, err
	}

	input := &glue.GetPartitionIndexesInput{
		CatalogId:    aws.String(catalogID),
		DatabaseName: aws.String(dbName),
		TableName:    aws.String(tableName),
	}

	var result *glue.PartitionIndexDescriptor

	output, err := conn.GetPartitionIndexes(input)

	if tfawserr.ErrCodeEquals(err, glue.ErrCodeEntityNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	for _, partInd := range output.PartitionIndexDescriptorList {
		if partInd == nil {
			continue
		}

		if aws.StringValue(partInd.IndexName) == partIndex {
			result = partInd
			break
		}
	}

	if result == nil {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	return result, nil
}
