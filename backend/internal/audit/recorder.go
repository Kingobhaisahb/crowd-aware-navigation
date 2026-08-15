package audit

type Recorder struct {
	records []AuditRecord
}

func NewRecorder() *Recorder {
	return &Recorder{
		records: make([]AuditRecord, 0),
	}
}

func (r *Recorder) Record(record AuditRecord) {
	for _, existing := range r.records {
		if existing.ID == record.ID {
			return
		}
	}

	r.records = append(r.records, record)
}

func (r *Recorder) GetAll() []AuditRecord {
	result := make([]AuditRecord, len(r.records))
	copy(result, r.records)

	return result
}

func (r *Recorder) Count() int {
	return len(r.records)
}