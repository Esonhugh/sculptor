package GoDataExtractor

func (q *DataExtractor[T]) SetTargetStruct(targetStruct any) *DataExtractor[T] {
	q.targetStruct = targetStruct
	return q
}

func (q *DataExtractor[T]) NewTargetStruct() *DataExtractor[T] {
	var targetInstance T
	q.targetStruct = targetInstance
	return q
}

func (q *DataExtractor[T]) CurrentTarget() T {
	return q.targetStruct
}
