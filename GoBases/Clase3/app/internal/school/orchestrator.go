package school

func Orchestrator(operator QualificationOperator) (qop QualificationOperation, err string) {
	switch operator {
	case MinimumConst:
		qop = Minimum
	case MaximumConst:
		qop = Maximum
	case AverageConst:
		qop = Average
	default:
		err = "Invalid operation"
	}
	return
}
