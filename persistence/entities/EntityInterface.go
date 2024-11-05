package entities

type EntityInterface interface {
	DBTableName() string
	EntityName() string
	EntityFields() []string
}
