package GoDataExtractor

func (q *DataExtractor) SetTargetStruct(targetStruct any) *DataExtractor {
	q.targetStruct = targetStruct
	return q
}

func (q *DataExtractor) CurrentTarget() any {
	return q.targetStruct
}
